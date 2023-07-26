// Code generated by goctl. DO NOT EDIT.
package types

type UploadResp struct {
	Url      string `json:"url"`
	FileName string `json:"fileName"`
}

type UploadImgReq struct {
	WatermarkFlag int64 `form:"watermarkFlag"`
}

type DownloadReq struct {
	Id int64 `form:"id,string"`
}

type DownloadResp struct {
	Content  []byte `json:"content"`
	FileName string `json:"fileName"`
}

type SuccessResp struct {
	Msg string `json:"msg"`
}

type MultipartUploadInitReq struct {
	FileSha1 string `json:"fileSha1"`
	FileSize int64  `json:"fileSize"`
}

type MultipartUploadInitRes struct {
	UserId     int64  `json:"userId,string"`
	FileSha1   string `json:"fileSha1"`
	FileSize   int64  `json:"fileSize"`
	UploadId   int64  `json:"uploadId,string"`
	ChunkSize  int64  `json:"chunkSize"`
	ChunkCount int64  `json:"chunkCount"`
}

type MultipartUploadSendReq struct {
	UploadID   int64 `form:"uploadId,string"`
	ChunkIndex int64 `form:"chunkIndex"`
}

type MultipartUploadCompleteReq struct {
	FileSha1 string `json:"fileSha1"`
	UploadId int64  `json:"uploadId,string"`
}

type MultipartUploadCompleteRes struct {
	AssetId int64 `form:"assetId,string"`
}
