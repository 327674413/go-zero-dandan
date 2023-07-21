package storaged

import (
	"net/http"
)

type FileType string

const (
	FileTypeImage FileType = "img"
	FileTypeVideo FileType = "video"
	FileTypeFile  FileType = "file"
)
const (
	ProviderLocal  string = "local"
	ProviderMinio  string = "minio"
	ProviderAliOss string = "aliyun"
	ProviderTxCos  string = "tencent"
)

// InterfaceStorage 存储接口
type InterfaceStorage interface {
	UploadImg(r *http.Request) error
	GetHash() (string, error)
	UploadFile()
}

// StorageConfig 配置
type StorageConfig struct {
	Provider string //服务对象
	Endpoint string // 端点地址
	Key      string // 访问Key
	Secret   string // 访问Secret
	Bucket   string // 存储桶名称
	FileType FileType
}
