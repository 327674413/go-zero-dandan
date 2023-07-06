package logic

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go-zero-dandan/app/asset/api/internal/svc"
	"go-zero-dandan/app/asset/api/internal/types"
	"net/http"
	"path"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild"
)

type UploadImgLogic struct {
	logx.Logger
	ctx        context.Context
	svcCtx     *svc.ServiceContext
	lang       *i18n.Localizer
	platId     int64
	platClasEm int64
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
	_ = r.ParseMultipartForm(maxMemorySize) //控制表单数据在内存中的存储大小，超过该值，则会自动将表单数据写入磁盘临时文件
	//logc.Info(l.ctx, "测试ctx日志")
	file, handler, err := r.FormFile("img")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer file.Close()
	if handler.Size > maxImgSize {
		// 文件大小超过限制
		return nil, resd.Fail("file size must less 5m")
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
	if _, err = file.Read(buffer); err != nil {
		return nil, resd.Fail("unknow file type")
	}
	// 判断文件 MIME 类型是否为图片类型
	mime := http.DetectContentType(buffer)
	if !validImageTypes[mime] {
		// 文件类型不是图片类型，返回错误信息
		return nil, resd.Fail("invalid file type")
	}
	resp = &types.UploadResp{}
	minioRes, err := l.UploadMinio(r)
	if err != nil {
		return nil, resd.Fail(err.Error())
	} else {
		resp.Url = minioRes
		return resp, nil
	}
	/*
		//判断与生成目录
		now := time.Now()
		year, month, day := now.Date()
		dirName := fmt.Sprintf("%04d-%02d-%02d", year, int(month), day)
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
			Url: url,
		}, nil*/
}
func (l *UploadImgLogic) UploadMinio(r *http.Request) (string, error) {
	minioClient, err := minio.New("127.0.0.1:9000", &minio.Options{
		Creds:  credentials.NewStaticV4(l.svcCtx.Config.Minio.AccessKey, l.svcCtx.Config.Minio.SecretKey, ""),
		Secure: false,
	})
	if err != nil {
		fmt.Println("是这里错误")
		return "", err
	}
	// 获取文件信息
	file, fileHeader, err := r.FormFile("img")
	bucketName := "pic"
	objectName := fmt.Sprintf("%d%s", utild.MakeId(), path.Ext(fileHeader.Filename))
	_, err = minioClient.PutObject(context.Background(), bucketName, objectName, file, fileHeader.Size,
		minio.PutObjectOptions{ContentType: "binary/octet-stream"})
	if err != nil {
		return "", err
	}
	return "http://localhost/" + bucketName + "/" + objectName, nil
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
func (l *UploadImgLogic) initPlat() (err error) {
	platClasEm := utild.AnyToInt64(l.ctx.Value("platClasEm"))
	if platClasEm == 0 {
		return resd.FailCode(l.lang, resd.PlatClasErr)
	}
	platClasId := utild.AnyToInt64(l.ctx.Value("platId"))
	if platClasId == 0 {
		return resd.FailCode(l.lang, resd.PlatIdErr)
	}
	l.platId = platClasId
	l.platClasEm = platClasEm
	return nil
}
