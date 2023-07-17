package logic

import (
	"context"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-dandan/app/asset/api/internal/svc"
	"go-zero-dandan/app/asset/api/internal/types"
	"go-zero-dandan/app/asset/model"
	"go-zero-dandan/common/constd"
	"go-zero-dandan/common/dao"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/storaged"
	"go-zero-dandan/common/utild"
	"mime/multipart"
	"net/http"
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
	fileData   *model.AssetMain
}

func NewUploadImgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadImgLogic {
	localizer := ctx.Value("lang").(*i18n.Localizer)
	l := &UploadImgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		lang:   localizer,
	}
	l.initPlat()
	return l
}

const maxMemorySize = 20 << 20 // 20 MB
const maxImgSize = 5 << 20

func (l *UploadImgLogic) UploadImg(r *http.Request, req *types.UploadImgReq) (resp *types.UploadResp, err error) {
	uploader, err := l.svcCtx.Storage.CreateUploader(&storaged.UploaderConfig{FileType: storaged.FileTypeImage})
	if err != nil {
		return nil, resd.ApiFail(l.lang, err)
	}
	hash, err := uploader.GetHash(r, "img")
	if err != nil {
		return l.apiFail(err)
	}
	//检查是否已经存在
	assetMainModel := model.NewAssetMainModel(l.svcCtx.SqlConn)
	whereStr := fmt.Sprintf("hash='%s' AND state_em > 1", hash)
	find, err := assetMainModel.WhereStr(whereStr).Find()
	//存在，则秒传
	if err == nil {
		return &types.UploadResp{
			Url:      find.Url,
			FileName: find.Name,
		}, nil
	}
	//查询报错
	if err != sqlx.ErrNotFound {
		return l.apiFail(resd.Error(err))
	}
	//不存在，则上传
	res, err := uploader.UploadImg(r)
	if err != nil {
		return l.apiFail(err)
	}
	assetMainData := &model.AssetMain{
		Id:       utild.MakeId(),
		StateEm:  constd.AssetStateEmFinish,
		Hash:     res.Hash,
		Name:     res.Name,
		ModeEm:   l.svcCtx.Config.AssetMode,
		Mime:     res.Mime,
		SizeNum:  res.SizeByte,
		SizeText: res.SizeText,
		Ext:      res.Ext,
		Url:      res.Url,
	}

	data, err := dao.PrepareData(assetMainData)
	_, err = assetMainModel.Insert(data)
	if err != nil {
		return l.apiFail(err)
	}
	return &types.UploadResp{
		Url:      res.Url,
		FileName: res.Name,
	}, nil
	/*if err = l.initPlat(); err != nil {
		return nil, err
	}
	_ = r.ParseMultipartForm(maxMemorySize) //控制表单数据在内存中的存储大小，超过该值，则会自动将表单数据写入磁盘临时文件
	//获取文件，前端传递字段要对应
	file, handler, err := r.FormFile("img")
	if err != nil {
		logc.Error(l.ctx, err)
		return nil, resd.FailCode(l.lang, resd.ReqFieldRequired, []string{"img"})
	}
	defer file.Close()
	//检查图片格式、大小等
	err = l.setFileData(file, handler, r)
	if err != nil {
		return nil, err
	}
	//检查是否已经存在
	assetMainModel := model.NewAssetMainModel(l.svcCtx.SqlConn, l.platId)
	whereStr := fmt.Sprintf("hash='%s' AND state_em > 1", l.fileData.Hash)
	find, err := assetMainModel.WhereStr(whereStr).Find()
	if err != nil {
		logc.Error(l.ctx, err)
		return nil, resd.FailCode(l.lang, resd.MysqlSelectErr)
	}
	resp = &types.UploadResp{}
	if find.Id != 0 {
		//return nil, resd.FailCode(l.lang, resd.DataExist1, []string{"Image"})
		resp.Url = find.Url
		resp.FileName = find.Name
		return resp, nil
	}
	l.fileData.Id = utild.MakeId()
	l.fileData.ModeEm = l.svcCtx.Config.AssetMode
	insertData, err := utild.MakeModelData(l.fileData, "Id,Hash,Name,Mime,SizeNum,SizeText,ModeEm,Ext")
	if err != nil {
		logc.Error(l.ctx, err)
		return nil, resd.FailCode(l.lang, resd.SysErr)
	}
	_, err = assetMainModel.Insert(insertData)
	if err != nil {
		logc.Error(l.ctx, err)
		return nil, resd.FailCode(l.lang, resd.MysqlInsertErr)
	}
	uploadImgStrage := map[int64]func(multipart.File, *multipart.FileHeader, *http.Request, *types.UploadImgReq) error{
		constd.AssetModeLocal:  l.UploadImgLocal,
		constd.AssetModeMinio:  l.UploadImgMinio,
		constd.AssetModeAliOss: l.UploadImgAliOss,
		constd.AssetModeTxCos:  l.UploadImgTxCos,
	}
	if strateFunc, ok := uploadImgStrage[l.svcCtx.Config.AssetMode]; ok {
		err = strateFunc(file, handler, r, req)
	} else {
		err = l.UploadImgLocal(file, handler, r, req)
	}
	if err != nil {
		return nil, err
	}
	l.fileData.StateEm = constd.AssetStateEmFinish
	updateData, err := utild.MakeModelData(l.fileData, "Id,Path,StateEm,Url")
	if err != nil {
		logc.Error(l.ctx, err)
		return nil, resd.FailCode(l.lang, resd.SysErr)
	}
	_, err = assetMainModel.Update(updateData)
	if err != nil {
		logc.Error(l.ctx, err)
		return nil, resd.FailCode(l.lang, resd.MysqlUpdateErr)
	}
	resp.Url = l.fileData.Url
	resp.FileName = l.fileData.Name*/
	return resp, nil

}
func (l *UploadImgLogic) setFileData(file multipart.File, handler *multipart.FileHeader, r *http.Request) (err error) {
	if handler.Size > maxImgSize {
		// 文件大小超过限制
		logc.Error(l.ctx, "image size over 5m")
		return resd.FailCode(l.lang, resd.UploadImageSizeLimited1, []string{"5M"})
	}
	//获取文件hash
	hash, err := utild.GetFileHashHex(file)
	if err != nil {
		logc.Error(l.ctx, err)
		return resd.FailCode(l.lang, resd.UploadFileFail)
	}
	//重新指向文件头，避免上传minio时长度不对
	if _, err = file.Seek(0, 0); err != nil {
		logc.Error(l.ctx, err)
		return resd.FailCode(l.lang, resd.SysErr)
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
		logc.Error(l.ctx, "unsupport image type")
		return resd.FailCode(l.lang, resd.NotSupportImageType)
	}
	// 判断文件 MIME 类型是否为图片类型
	mime := http.DetectContentType(buffer)
	if !validImageTypes[mime] {
		logc.Error(l.ctx, "invalid img type")
		// 文件类型不是图片类型，返回错误信息
		return resd.FailCode(l.lang, resd.NotSupportImageType)
	}
	//重新指向文件头，避免上传minio时长度不对
	if _, err = file.Seek(0, 0); err != nil {
		logc.Error(l.ctx, err)
		return resd.FailCode(l.lang, resd.SysErr)
	}
	l.fileData = &model.AssetMain{
		Name:     handler.Filename,
		Ext:      filepath.Ext(handler.Filename),
		Hash:     hash,
		Mime:     mime,
		SizeNum:  handler.Size,
		SizeText: utild.FormatFileSize(handler.Size),
	}
	return err
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

func (l *UploadImgLogic) UploadImgAliOss(file multipart.File, handler *multipart.FileHeader, r *http.Request, req *types.UploadImgReq) (err error) {
	return nil
}
func (l *UploadImgLogic) UploadImgTxCos(file multipart.File, handler *multipart.FileHeader, r *http.Request, req *types.UploadImgReq) (err error) {

	objectName := fmt.Sprintf("img/%d/%s/%s%s", l.platId, getDirName(), l.fileData.Hash, l.fileData.Ext)

	/*
		//通过本地文件路径上传
		_, _, err = client.Object.Upload(
			context.Background(), objectName, "localfilePaht", nil,
		)
	*/
	_, err = l.svcCtx.TxCos.Object.Put(context.Background(), objectName, file, nil)
	if err != nil {
		logc.Error(l.ctx, err)
		return resd.FailCode(l.lang, resd.UploadFileFail)
	}
	l.fileData.Url = l.svcCtx.Config.TxCos.PublicBucketAddr + "/" + objectName
	return nil
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

func (l *UploadImgLogic) apiFail(err error) (*types.UploadResp, error) {
	return nil, resd.ApiFail(l.lang, err)
}
