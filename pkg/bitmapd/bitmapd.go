package bitmapd

type Bitmap struct {
	bits []byte
	size int64
}

func NewBitmap(size ...int) *Bitmap {
	s := 100
	if len(size) > 0 {
		s = size[0]
	}
	return &Bitmap{
		bits: make([]byte, s),
		size: int64(s * 8),
	}
}
func Load(bits []byte) *Bitmap {
	if len(bits) == 0 {
		return NewBitmap(0)
	}
	return &Bitmap{
		bits: bits,
		size: int64(len(bits) * 8),
	}
}
func (t *Bitmap) SetId(id string) {
	// id在哪个bit
	idx := hash(id) % t.size
	// 计算在哪个byte
	byteIdx := idx / 8
	// 在这个byte中的哪个bit位置
	bitIdx := idx % 8
	// 设置值
	t.bits[byteIdx] |= 1 << bitIdx
}
func (t *Bitmap) IsSetId(id string) bool {
	// id在哪个bit
	idx := hash(id) % t.size
	// 计算在哪个byte
	byteIdx := idx / 8
	// 在这个byte中的哪个bit位置
	bitIdx := idx % 8
	return (t.bits[byteIdx] & (1 << bitIdx)) != 0
}

func (t *Bitmap) Export() []byte {
	return t.bits
}

func hash(str string) int64 {
	// 使用kor哈希算法
	seed := 131313
	val := 0
	for _, c := range str {
		val = val*seed + int(c)
	}
	return int64(val & 0x7FFFFFF)
}
