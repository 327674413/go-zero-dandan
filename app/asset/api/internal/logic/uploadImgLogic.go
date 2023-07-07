package logic

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/asset/api/internal/svc"
	"go-zero-dandan/app/asset/api/internal/types"
	"go-zero-dandan/common/constd"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
)

type UploadImgLogic struct {
	logx.Logger
	ctx        context.Context
	svcCtx     *svc.ServiceContext
	lang       *i18n.Localizer
	platId     int64
	platClasEm int64
	fileName   string
}

func NewUploadImgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadImgLogic {
	localizer := ctx.Value("lang").(*i18n.Localizer)
	return &UploadImgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		lang:   localizer,
	}
}

const maxMemorySize = 20 << 20 // 20 MB
const maxImgSize = 5 << 20

func (l *UploadImgLogic) UploadImg(r *http.Request, req *types.UploadImgReq) (resp *types.UploadResp, err error) {
	if err = l.initPlat(); err != nil {
		return nil, err
	}
	_ = r.ParseMultipartForm(maxMemorySize) //控制表单数据在内存中的存储大小，超过该值，则会自动将表单数据写入磁盘临时文件
	file, handler, err := r.FormFile("img")
	if err != nil {
		logc.Error(l.ctx, err)
		return nil, resd.FailCode(l.lang, resd.ReqFieldRequired, []string{"img"})
	}
	defer file.Close()
	err = l.checkImg(file, handler, r)
	if err != nil {
		return nil, err
	}
	uploadImgStrage := map[int64]func(multipart.File, *multipart.FileHeader, *http.Request, *types.UploadImgReq) (*types.UploadResp, error){
		constd.AssetModeLocal:  l.UploadImgLocal,
		constd.AssetModeMinio:  l.UploadImgMinio,
		constd.AssetModeOssAli: l.UploadImgOssAli,
		constd.AssetModeOssTx:  l.UploadImgOssTx,
	}
	if strateFunc, ok := uploadImgStrage[l.svcCtx.Config.AssetMode]; ok {
		return strateFunc(file, handler, r, req)
	} else {
		return l.UploadImgLocal(file, handler, r, req)
	}

}
func (l *UploadImgLogic) checkImg(file multipart.File, handler *multipart.FileHeader, r *http.Request) error {
	if handler.Size > maxImgSize {
		// 文件大小超过限制
		logc.Error(l.ctx, "image size over 5m")
		return resd.FailCode(l.lang, resd.ImageSizeLimited1, []string{"5M"})
	}
	validImageTypes := map[string]bool{
		"image/jpeg":    true,
		"image/jpg":     true,
		"image/png":     true,
		"image/gif":     true,
		"image/bmp":     false,
		"image/svg+xml": false,
	}
	// 读取文件前 512 字节
	buffer := make([]byte, 512)
	if _, err := file.Read(buffer); err != nil {
		logc.Error(l.ctx, "unsupport image type")
		return resd.FailCode(l.lang, resd.ImageSizeLimited1, []string{"5M"})
	}
	// 判断文件 MIME 类型是否为图片类型
	mime := http.DetectContentType(buffer)
	if !validImageTypes[mime] {
		logc.Error(l.ctx, "invalid img type")
		// 文件类型不是图片类型，返回错误信息
		return resd.FailCode(l.lang, resd.NotSupportImageType)
	}
	//重新指向文件头，避免上传minio时长度不对
	if _, err := file.Seek(0, 0); err != nil {
		logc.Error(l.ctx, err)
		return resd.FailCode(l.lang, resd.Err)
	}
	l.fileName = handler.Filename
	return nil
}
func (l *UploadImgLogic) UploadImgMinio(file multipart.File, handler *multipart.FileHeader, r *http.Request, req *types.UploadImgReq) (resp *types.UploadResp, err error) {
	resp = &types.UploadResp{}
	url, err := l.UploadMinio(file, handler, r)
	if err != nil {
		return nil, err
	} else {
		resp.Url = url
		resp.FileName = l.fileName
		return resp, nil
	}
}
func (l *UploadImgLogic) UploadImgLocal(file multipart.File, handler *multipart.FileHeader, r *http.Request, req *types.UploadImgReq) (resp *types.UploadResp, err error) {

	//判断与生成目录
	dirName := getDirName()
	dirPath := filepath.Join(l.svcCtx.Config.AssetPath.Img, dirName)
	if err = os.MkdirAll(dirPath, 0755); err != nil {
		return nil, resd.Fail(err.Error())
	}
	//拼接返回的url地址
	url := ""
	if r.TLS == nil {
		url = "http://"
	} else {
		url = "https://"
	}
	url = url + r.Host + r.URL.Path
	//获取文件后缀
	ext := filepath.Ext(handler.Filename)
	//根据雪花id生成新的文件名
	newFileName := fmt.Sprintf("%d%s", utild.MakeId(), ext)
	//获取完整的存储路径
	savePath := path.Join(dirPath, newFileName)
	//存储文件
	tempFile, err := os.Create(savePath)
	if err != nil {
		fmt.Println(err)
		return nil, resd.Fail(err.Error())
	}
	defer tempFile.Close()
	io.Copy(tempFile, file)

	return &types.UploadResp{
		Url:      url,
		FileName: l.fileName,
	}, nil
}
func (l *UploadImgLogic) UploadMinio(file multipart.File, handler *multipart.FileHeader, r *http.Request) (string, error) {
	minioClient, err := minio.New(l.svcCtx.Config.Minio.Address, &minio.Options{
		Creds:  credentials.NewStaticV4(l.svcCtx.Config.Minio.AccessKey, l.svcCtx.Config.Minio.SecretKey, ""),
		Secure: false,
	})
	if err != nil {
		logc.Error(l.ctx, err)
		return "", resd.FailCode(l.lang, resd.Err)
	}
	bucketName := "public"
	objectName := fmt.Sprintf("img/%d/%s/%d%s", l.platId, getDirName(), utild.MakeId(), path.Ext(handler.Filename))
	_, err = minioClient.PutObject(context.Background(), bucketName, objectName, file, handler.Size,
		minio.PutObjectOptions{ContentType: "binary/octet-stream"})
	if err != nil {
		logc.Error(l.ctx, err)
		return "", resd.FailCode(l.lang, resd.Err)
	}
	return "http://" + l.svcCtx.Config.Minio.Address + "/" + bucketName + "/" + objectName, nil
}

/*
// 下载图片

	func (l *UploadImgLogic) DownloadImg(r *http.Request, req *types.DownloadImgReq) (resp *types.DownloadImgResp, err error) {
		// 获取图片完整路径
		imgPath := path.Join(l.svcCtx.Config.AssetPath.Img, req.Dir, req.Filename)
		file, err := os.Open(imgPath)
		if err != nil {
			return nil, resd.Fail("file not found")
		}
		defer file.Close()
		fileInfo, err := file.Stat()
		if err != nil {
			return nil, resd.Fail("file not found")
		}
		// 设置响应头，让浏览器下载文件
		w := r.Context().Value("response").(http.ResponseWriter)
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", req.Filename))
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Length", strconv.FormatInt(fileInfo.Size(), 10))
		// 将文件内容写入响应
		io.Copy(w, file)
		return &types.DownloadImgResp{}, nil
	}
*/

func (l *UploadImgLogic) UploadImgOssAli(file multipart.File, handler *multipart.FileHeader, r *http.Request, req *types.UploadImgReq) (resp *types.UploadResp, err error) {
	resp = &types.UploadResp{}
	return resp, nil
}
func (l *UploadImgLogic) UploadImgOssTx(file multipart.File, handler *multipart.FileHeader, r *http.Request, req *types.UploadImgReq) (resp *types.UploadResp, err error) {
	resp = &types.UploadResp{}
	return resp, nil
}
func (l *UploadImgLogic) initPlat() (err error) {
	platClasEm := utild.AnyToInt64(l.ctx.Value("platClasEm"))
	if platClasEm == 0 {
		logc.Error(l.ctx, "token not get platClasEm")
		return resd.FailCode(l.lang, resd.PlatClasErr)
	}
	platClasId := utild.AnyToInt64(l.ctx.Value("platId"))
	if platClasId == 0 {
		logc.Error(l.ctx, "token not get platId")
		return resd.FailCode(l.lang, resd.PlatIdErr)
	}
	l.platId = platClasId
	l.platClasEm = platClasEm
	return nil
}
func getDirName() string {
	now := time.Now()
	year, month, day := now.Date()
	return fmt.Sprintf("%04d-%02d-%02d", year, int(month), day)
}
