package storaged

import (
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/tencentyun/cos-go-sdk-v5"
	"go-zero-dandan/common/imgd"
	"net/http"
	"time"
)

// FileType 定义文件类型，不同文件可能会有不同的特殊方法，如图片类型的裁剪、水印操作
type FileType string

const (
	FileTypeImage     FileType = "img"       //图片
	FileTypeVideo     FileType = "video"     //视频
	FileTypeFile      FileType = "file"      //普通文件
	FileTypeMultipart FileType = "multipart" //分片文件
)

// 定义支持的渠道
const (
	ProviderLocal  string = "local"   //本地
	ProviderMinio  string = "minio"   //minio
	ProviderAliOss string = "aliyun"  //阿里云
	ProviderTxCos  string = "tencent" //腾讯云
)

// InterfaceFactory 文件管理工厂入口
type InterfaceFactory interface {
	// Init 初始化操作
	Init() error
	// CreateUploader 创建上传器
	CreateUploader(config *UploaderConfig) (InterfaceStorage, error)
	// CreateDownloader 创建下载器
	CreateDownloader(config *DownloaderConfig) (InterfaceStorage, error)
}

// InterfaceStorage 文件管理器接口
type InterfaceStorage interface {
	// Upload 简单文件上传
	Upload(r *http.Request, config *UploadConfig) (*UploadResult, error)
	// MultipartUpload 分片上传
	MultipartUpload(r *http.Request, config *UploadConfig) (*UploadResult, error)
	// MultipartMerge 分片合并
	MultipartMerge(fileSha1 string, saveName string, chunkCount int) error
	// UploadImg 图片上传专用，提供一些图片处理方法
	UploadImg(r *http.Request, config *UploadImgConfig) (*UploadResult, error)
	// Download 简单文件下载
	Download(w http.ResponseWriter, objectName string, fileName ...string) error
	// MultipartDownload 分片下载
	MultipartDownload(w http.ResponseWriter, path string) error
	// GetHash 预先获取上传文件的hash值（sha1）
	GetHash(r *http.Request, formKey string) (string, error)
}

// UploadImgResizeType 图片缩放类型
type UploadImgResizeType string

/*
cover 能保证撑满目标尺寸，若比例不同，则在长边上会发生截取
contain 能保证图片完整在目标尺寸中显示，若比例不同，则生成的图片在短边上会有空白
fill 不保持长宽比例直接缩放，若比例不同，则可能会变形
widthFix 宽度变成目标尺寸宽度，高度则等比例缩放
heightFix 高度变成目标尺寸高度，宽度则等比例缩放
*/
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

// UploadConfig 上传配置，预留
type UploadConfig struct {
	IsMultipart bool   //是否分片上传
	FileSha1    string //文件sha1
	ChunkIndex  int64  //分片标识，分片上传使用
}

// UploadImgConfig 图片上传配置
type UploadImgConfig struct {
	//Quality 图片质量，默认100不压缩
	Quality int
	//Resize 图片缩放配置
	Resize *UploadImgResizeConfig
	//Watermark 水印配置
	WatermarkConfig *imgd.WatermarkConfig
}

// ProviderConfig 渠道配置
type ProviderConfig struct {
	Provider  string //服务对象
	Endpoint  string // 端点地址
	Key       string // 访问Key
	Secret    string // 访问Secret
	Bucket    string // 存储桶名称
	LocalPath string //本地存储时的路径
}

// todo::是不是这里没用，为啥阿没有阿里云的
// StorageSvc 服务依赖
type StorageSvc struct {
	MinioClient *minio.Client
	TxCosClient *cos.Client
}

// UploaderConfig 上传文件配置
type UploaderConfig struct {
	FileType      FileType //todo::讲道理不需要在初始化阶段就先定义，应该可以去掉
	MaxFileSize   int64    //限制文件大小
	MaxMemorySize int64    //限制内存大小
	AcceptMimes   []string //支持接收的文件类型
	RejectMimes   []string //拒绝接收的文件类型
	DirName       string   //子目录名称
	Bucket        string   //桶
	FormKey       string   //上传form字段
}

// DownloaderConfig 下载文件配置，暂无用，预留
type DownloaderConfig struct {
}

// UploadRes 上传结果集对象
type UploadRes struct {
	FileName string //文件名称
	Ext      string //文件后缀，带有点号
	Mime     string //文件类型
	Url      string //访问地址
	Hash     string //sha1哈希值
	Size     string //文件大小，byte
	SizeText string //文件大小，格式化处理
}

// GetDateDir 根据年-月-日获取文件夹名称
func GetDateDir() string {
	now := time.Now()
	year, month, day := now.Date()
	return fmt.Sprintf("%04d-%02d-%02d", year, int(month), day)
}
