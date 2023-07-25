package storaged

import (
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/tencentyun/cos-go-sdk-v5"
	"go-zero-dandan/common/imgd"
	"net/http"
	"time"
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
	Init() error
	CreateUploader(*UploaderConfig) (InterfaceUploader, error)
}

// InterfaceUploader 上传器接口
type InterfaceUploader interface {
	UploadImg(r *http.Request, config *UploadImgConfig) (*UploadResult, error)
	Download(w http.ResponseWriter, path string) error
	GetHash(r *http.Request, formKey string) (string, error)
}

// UploadImgResizeType 图片缩放类型
type UploadImgResizeType string

const (
	UploadImgResizeTypeCover     UploadImgResizeType = "cover"
	UploadImgResizeTypeContain   UploadImgResizeType = "contain"
	UploadImgResizeTypeFill      UploadImgResizeType = "fill"
	UploadImgResizeTypeWidthFix  UploadImgResizeType = "widthFix"
	UploadImgResizeTypeHeightFix UploadImgResizeType = "heightFix"
)

// UploadImgResizeConfig 图片缩放配置
type UploadImgResizeConfig struct {
	Height int
	Width  int
	Type   UploadImgResizeType
}

// UploadImgConfig 图片上传配置
type UploadImgConfig struct {
	Quality   int                    //图片质量，默认100不压缩
	Resize    *UploadImgResizeConfig //图片缩放配置
	Watermark *imgd.WatermarkConfig  //水印配置
}

// StorageConfig 配置
type StorageConfig struct {
	Provider  string //服务对象
	Endpoint  string // 端点地址
	Key       string // 访问Key
	Secret    string // 访问Secret
	Bucket    string // 存储桶名称
	LocalPath string //本地存储时的路径
}

// StorageSvc 服务依赖
type StorageSvc struct {
	MinioClient *minio.Client
	TxCosClient *cos.Client
}

// UploaderConfig 上传文件配置
type UploaderConfig struct {
	FileType       FileType
	MaxFileSize    int64 //限制文件大小
	MaxMemorySize  int64 //限制内存大小
	FileMimeAccept []string
}

type UploadRes struct {
	FileName string
	Ext      string
	Mime     string
	Url      string
	Hash     string
	Size     string
	SizeText string
}

func getDirName() string {
	now := time.Now()
	year, month, day := now.Date()
	return fmt.Sprintf("%04d-%02d-%02d", year, int(month), day)
}
