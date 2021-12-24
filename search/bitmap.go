package search

import (
	"fmt"
	"strings"
)

const (
	bitSize = 8
)

var bitmask = []byte{1, 1 << 1, 1 << 2, 1 << 3, 1 << 4, 1 << 5, 1 << 6, 1 << 7}

// 首字母小写 只能调用 工厂函数 创建
type bitmap struct {
	bits     []byte
	bitCount uint64 // 已填入数字的数量
	capacity uint64 // 容量
}

// 创建工厂函数
func NewBitmap(maxNum uint64) *bitmap {
	return &bitmap{bits: make([]byte, maxNum/bitSize+1), bitCount: 0, capacity: maxNum}
}

// 填入数字
func (b *bitmap) Add(num uint64) {
	byteIndex, bitPos := b.offset(num)
	// 1 左移 bitPos 位 进行 按位或 (置为 1)
	b.bits[byteIndex] |= bitmask[bitPos]
	b.bitCount++
}

// 清除填入的数字
func (b *bitmap) Del(num uint64) {
	byteIndex, bitPos := b.offset(num)
	// 重置为空位 (重置为 0)
	b.bits[byteIndex] &= ^bitmask[bitPos]
	b.bitCount--
}

// 数字是否在位图中
func (b *bitmap) Contains(num uint64) bool {
	byteIndex := num / bitSize
	if byteIndex >= uint64(len(b.bits)) {
		return false
	}
	bitPos := num % bitSize
	//  1左移 bitPos 位 和 1 进行 按位与
	return !(b.bits[byteIndex]&bitmask[bitPos] == 0)
}

func (b *bitmap) offset(num uint64) (byteIndex uint64, bitPos byte) {
	byteIndex = num / bitSize // 字节索引
	if byteIndex >= uint64(len(b.bits)) {
		panic(fmt.Sprintf(" runtime error: index value %d out of range", byteIndex))
		return
	}
	bitPos = byte(num % bitSize) // bit位置
	return byteIndex, bitPos
}

// 位图的容量
func (b *bitmap) Size() uint64 {
	return uint64(len(b.bits) * bitSize)
}

// 是否空位图
func (b *bitmap) IsEmpty() bool {
	return b.bitCount == 0
}

// 是否已填满
func (b *bitmap) IsFully() bool {
	return b.bitCount == b.capacity
}

// 已填入的数字个数
func (b *bitmap) Count() uint64 {
	return b.bitCount
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

func (b *bitmap) String() string {
	var sb strings.Builder
	for index := len(b.bits) - 1; index >= 0; index-- {
		sb.WriteString(byteToBinaryString(b.bits[index]))
		sb.WriteString(" ")
	}
	return sb.String()
}

func byteToBinaryString(data byte) string {
	var sb strings.Builder
	for index := 0; index < bitSize; index++ {
		if (bitmask[7-index] & data) == 0 {
			sb.WriteString("0")
		} else {
			sb.WriteString("1")
		}
	}
	return sb.String()
}
