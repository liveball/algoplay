package main

import (
	"fmt"
)

//BitmapSize for default
const BitmapSize = 0x01 << 32

//BitMap for bitmap obj
type BitMap struct {
	bits    []byte
	bitSize uint64
	// 该 Bitmap 被设置为 1 的最大位置（方便遍历）
	maxpos uint64
}

//NewBitmapSize for create bitmap
func NewBitmapSize(size uint64) *BitMap {
	if size == 0 || size > BitmapSize {
		size = BitmapSize
	} else if remainder := size % 8; remainder != 0 {
		size += 8 - remainder
	}
	return &BitMap{
		bits:    make([]byte, size>>3),
		bitSize: uint64(size - 1),
	}
}

//SetBit 将 offset 位置的 bit 置为 value (0/1)
func (b *BitMap) SetBit(num uint64, value uint8) bool {
	// num/8得到byte[]的index
	index := num >> 3
	// num%8得到在byte[index]的位置
	pos := num & 0x07
	println(num, index, pos, num&0x07 == num%8)
	if b.bitSize < num {
		return false
	}

	if value == 0 {
		// &^ 清位, 将1左移pos后，那个位置自然就是1，然后对取反，再与当前值做&，即可清除当前的位置了.
		b.bits[index] &^= 0x01 << pos
	} else {
		//将1左移pos后，那个位置自然就是1，然后和以前的数据做|，这样，那个位置就替换成1了。
		b.bits[index] |= 0x01 << pos
		// 记录曾经设置为 1 的最大位置
		if b.maxpos < num {
			b.maxpos = num
		}
	}
	return true
}

//Contain 查看某个数是否存在
func (b *BitMap) Contain(num uint64) bool {
	// num/8得到byte[]的index
	index := num >> 3
	// num%8得到在byte[index]的位置
	pos := num & 0x07
	if b.bitSize < num {
		return false
	}

	return (b.bits[index] & (0x01 << pos)) != 0
}

// GetBit 获得 offset 位置处的 value
func (b *BitMap) GetBit(offset uint64) uint8 {
	index, pos := offset/8, offset%8

	if b.bitSize < offset {
		return 0
	}

	return (b.bits[index] >> pos) & 0x01
}

// Maxpos 获的置为 1 的最大位置
func (b *BitMap) Maxpos() uint64 {
	return b.maxpos
}

// String 实现 Stringer 接口（只输出开始的100个元素）
func (b *BitMap) String() string {
	var maxTotal, bitTotal uint64 = 100, b.maxpos + 1

	if b.maxpos > maxTotal {
		bitTotal = maxTotal
	}

	numSlice := make([]uint64, 0, bitTotal)

	var offset uint64
	for offset = 0; offset < bitTotal; offset++ {
		if b.GetBit(offset) == 1 {
			numSlice = append(numSlice, offset)
		}
	}

	return fmt.Sprintf("%v", numSlice)
}
