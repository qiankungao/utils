package main

const (
	bitSize = 8
)

var bitMask = []byte{1, 1 << 1, 1 << 2, 1 << 3, 1 << 4, 1 << 5, 1 << 6, 1 << 7}

type bitmap struct {
	bits []byte
	//已经填入数字的个数
	bitCount uint64
	//容量
	bitCamp uint64
}

//创建
func NewBitMap(maxNum uint64) *bitmap {
	return &bitmap{
		bits:     make([]byte, maxNum/bitSize+1),
		bitCount: 0,
		bitCamp:  maxNum,
	}
}

//add
func (b *bitmap) Add(num uint64) {
	//获取这个数的index和pos
	index, pos := b.Offset(num)
	//加入其实就是把 1左移pos位然后和b.bits[index] 或
	b.bits[index] |= bitMask[pos]
	b.bitCount++
}

//clear 清除一个数
func (b *bitmap) Clear(num uint64) {
	//获取这个数的index和pos
	index, pos := b.Offset(num)
	//清除就是把1左移pos,取反，然后在 &
	b.bits[index] &= ^bitMask[pos]
	b.bitCount--
}

//是否包含
func (b *bitmap) Contains(num uint64) bool {
	//获取这个数的index和pos
	index, pos := b.Offset(num)
	//是否包含
	return !(b.bits[index]&bitMask[pos] == 0)
}

//获取加入的数在bitmap的索引和位置
func (b *bitmap) Offset(num uint64) (uint64, uint64) {
	index := num / bitSize
	pos := num % bitSize
	return index, pos
}

//容量 有多少位bit
func (b *bitmap) Size() uint64 {
	return uint64(len(b.bits) * bitSize)
}

// 获取填入的数字切片
func (b *bitmap) GetData() []uint64 {
	var data []uint64
	count := b.Size()
	for index := uint64(0); index < count; index++ {
		if b.Contains(index) {
			data = append(data, index)
		}
	}
	return data
}

func (b *bitmap) string() string {
	return ""
}

func main() {


}
