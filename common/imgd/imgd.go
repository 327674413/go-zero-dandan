package imgd

import (
	"fmt"
	"github.com/disintegration/imaging"
	"go-zero-dandan/common/resd"
	"image"
	"path/filepath"
	"strings"
)

type ImgExt string

const (
	ImgExtJpg  ImgExt = ".jpg"
	ImgExtJpeg ImgExt = ".jpeg"
	ImgExtPng  ImgExt = ".png"
	ImgExtGif  ImgExt = ".gif"
)

type Imager struct {
	Img       image.Image
	watermark image.Image
	Path      string
	Ext       ImgExt
}

func NewImg(path string) (*Imager, error) {
	ext := ImgExt(strings.ToLower(filepath.Ext(path)))
	srcImage, err := imaging.Open(path)
	if err != nil {
		return nil, resd.Error(err)
	}
	return &Imager{Img: srcImage, Ext: ext, Path: path}, nil
}

/*
//实际测试imaging反而加大了图片，无效
// Quality 对图片进行质量压缩,原始100
func (t *Imager) Quality(num int) error {
	var buf bytes.Buffer
	err := jpeg.Encode(&buf, t.Img, &jpeg.Options{Quality: num})
	if err != nil {
		return resd.Error(err)
	}
	t.Img, err = imaging.Decode(&buf)
	if err != nil {
		return resd.Error(err)
	}
	return nil
}
*/

// ResizeCover 等比例缩放，保证短边能完全看到，长边可能被裁切
func (t *Imager) ResizeCover(width int, height int) {
	t.Img = imaging.Fill(t.Img, width, height, imaging.Center, imaging.Lanczos)
}

// ResizeContain 等比例缩放，保证长边能完全看到，可能会出现空白区域
func (t *Imager) ResizeContain(width int, height int) {
	t.Img = imaging.Fit(t.Img, width, height, imaging.Lanczos)
}

// ResizeFill 按目标尺寸进行拉伸缩放
func (t *Imager) ResizeFill(width int, height int) {
	t.Img = imaging.Fill(t.Img, width, height, imaging.Center, imaging.Box)
}

// ResizeWidthFix 按目标宽度，对高度等比例缩放
func (t *Imager) ResizeWidthFix(width int) {
	t.Img = imaging.Resize(t.Img, width, 0, imaging.Lanczos)
}

// ResizeHeightFix 按目标高度，对宽度等比例缩放
func (t *Imager) ResizeHeightFix(height int) {
	t.Img = imaging.Resize(t.Img, 0, height, imaging.Lanczos)
}

// Width 获取图片宽度
func (t *Imager) Width() int {
	return t.Img.Bounds().Dx()
}

// Height 获取图片高度
func (t *Imager) Height() int {
	return t.Img.Bounds().Dy()
}

// Open 加载新图片
func (t *Imager) Open(filePath string) (image.Image, error) {
	target, err := imaging.Open(filePath)
	if err != nil {
		return nil, resd.Error(err)
	}
	return target, nil
}

// Output 保存到文件
func (t *Imager) Output(pathAndFileName string) error {
	err := imaging.Save(t.Img, pathAndFileName)
	if err != nil {
		return resd.Error(err)
	} else {
		return nil
	}
}

// WatermarkImg 添加水印
func (t *Imager) WatermarkImg(config *WatermarkConfig) error {
	var err error
	t.watermark, err = imaging.Open(config.Path)
	if err != nil {
		return err
	}

	var offset image.Point
	switch config.Position {
	case WatermarkPositionRightBottom:
		offset, err = t.getWatermarkPositionRightBottom(t.watermark, config.OffsetX, config.OffsetY)
	case WatermarkPositionContain:
		offset, err = t.getWatermarkPositionContain(t.watermark)
	}
	t.Img = imaging.Overlay(t.Img, t.watermark, offset, 1.0)
	return nil
}
func (t *Imager) getWatermarkPositionRightBottom(waterImager image.Image, offsetX int, offsetY int) (image.Point, error) {
	// 计算水印图片的缩放比例
	scale := float64(t.Img.Bounds().Dx()) * 0.2 / float64(t.watermark.Bounds().Dx())
	// 缩放水印图片
	t.watermark = imaging.Resize(t.watermark, int(float64(t.watermark.Bounds().Dx())*scale), 0, imaging.Lanczos)

	return image.Pt(t.Img.Bounds().Dx()-waterImager.Bounds().Dx()-offsetX, t.Img.Bounds().Dy()-waterImager.Bounds().Dy()-offsetY), nil
}
func (t *Imager) getWatermarkPositionContain(waterImager image.Image) (image.Point, error) {
	//offset := image.Pt((t.Img.Bounds().Dx()-waterImager.Bounds().Dx())/2, (t.Img.Bounds().Dy()-waterImager.Bounds().Dy())/2)
	var scale float64
	var offset image.Point
	if t.Img.Bounds().Dx()*waterImager.Bounds().Dy() > t.Img.Bounds().Dy()*waterImager.Bounds().Dx() {
		// 水印图片的长边与目标图片的长边相同或更长
		scale = float64(t.Img.Bounds().Dx()) / float64(waterImager.Bounds().Dx())
		offset.Y = (t.Img.Bounds().Dy() - int(float64(waterImager.Bounds().Dy())*scale)) / 2
	} else {
		// 水印图片的短边与目标图片的长边相同或更长
		scale = float64(t.Img.Bounds().Dy()) / float64(waterImager.Bounds().Dy())
		offset.X = (t.Img.Bounds().Dx() - int(float64(waterImager.Bounds().Dx())*scale)) / 2
	}

	// 缩放水印图片
	t.watermark = imaging.Resize(t.watermark, int(float64(t.watermark.Bounds().Dx())*scale), 0, imaging.Lanczos)
	fmt.Println(offset)
	return offset, nil
}
