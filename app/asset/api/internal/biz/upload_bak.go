package biz

import (
	"context"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-dandan/app/asset/api/internal/svc"
	"go-zero-dandan/app/asset/api/internal/types"
	"go-zero-dandan/app/asset/model"
	"go-zero-dandan/common/constd"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild"
	"mime/multipart"
	"net/http"
	"path/filepath"
)

type Uploader struct {
	logx.Logger
	ctx        context.Context
	svcCtx     *svc.ServiceContext
	lang       *i18n.Localizer
	platId     int64
	platClasEm int64
	fileType   string
}

const maxImgSize = 5 << 20

func NewUpload(fileType string, ctx context.Context, svcCtx *svc.ServiceContext, platId int64, platClasEm int64, lang *i18n.Localizer) *Upload {

	config := FileExplorerConfig{
		Provider: "tencent",
		Endpoint: "your-endpoint",
		Key:      "your-access-key",
		Secret:   "your-secret-key",
		Bucket:   "your-bucket-name",
	}
	factory, err := NewFileExplorerFactory(config)
	if err != nil {
		// 处理错误
	}
	fileExplorer, err := factory.CreateFileExplorer()
	if err != nil {
		// 处理错误
	}
	fileExplorer.UploadImg()
	biz := &Upload{
		ctx:        ctx,
		svcCtx:     svcCtx,
		lang:       lang,
		platId:     platId,
		platClasEm: platClasEm,
		fileType:   fileType,
	}
	return biz
}
func (t *Uploader) Upload(r *http.Request) (res *UploadRes, err error) {
	//获取文件，前端传递字段要对应
	file, handler, err := t.getFormFile(r)
	if err != nil {
		logc.Error(t.ctx, err)
		return nil, err
	}
	defer file.Close()
	//检查图片格式、大小等
	err = t.setFileData(file, handler, r)
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
		return nil, resd.FailCode(l.lang, resd.Err)
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
		return nil, resd.FailCode(l.lang, resd.Err)
	}
	_, err = assetMainModel.Update(updateData)
	if err != nil {
		logc.Error(l.ctx, err)
		return nil, resd.FailCode(l.lang, resd.MysqlUpdateErr)
	}
	resp.Url = l.fileData.Url
	resp.FileName = l.fileData.Name
	return resp, nil
}

func (t *Uploader) setFileData(file multipart.File, handler *multipart.FileHeader, r *http.Request) (err error) {
	if handler.Size > maxImgSize {
		// 文件大小超过限制
		logc.Error(l.ctx, "image size over 5m")
		return resd.FailCode(l.lang, resd.ImageSizeLimited1, []string{"5M"})
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
		return resd.FailCode(l.lang, resd.Err)
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
		return resd.FailCode(l.lang, resd.Err)
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
func (t *Uploader) getFormFile(r *http.Request) (multipart.File, *multipart.FileHeader, error) {
	return nil, nil, resd.FailCode(t.lang, resd.NotSupportFileType)
}
