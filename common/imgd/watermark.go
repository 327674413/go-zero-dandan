package imgd

type watermarkType int

const (
	WatermarkTypeText watermarkType = 1
	WatermarkTypeImg  watermarkType = 2
)

type watermarkPosition string

const (
	WatermarkPositionLeftTop      watermarkPosition = "lt"
	WatermarkPositionLeftCenter   watermarkPosition = "lc"
	WatermarkPositionLeftBottom   watermarkPosition = "lb"
	WatermarkPositionCenterCenter watermarkPosition = "cc"
	WatermarkPositionCenterTop    watermarkPosition = "ct"
	WatermarkPositionCenterBottom watermarkPosition = "cb"
	WatermarkPositionRightTop     watermarkPosition = "rt"
	WatermarkPositionRightCent    watermarkPosition = "rc"
	WatermarkPositionRightBottom  watermarkPosition = "rb"
	WatermarkPositionContain      watermarkPosition = "contain"
)

type WatermarkConfig struct {
	Type     watermarkType
	Text     string
	Path     string
	Position watermarkPosition
	OffsetX  int
	OffsetY  int
}

func Watermark() {

}
