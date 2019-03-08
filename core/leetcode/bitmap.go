package main

import (
	"fmt"
)

// 再看代码之前，我们先搞清楚一个问题，一个数怎么快速定位它的索引号，也就是说搞清楚byte[index]的index是多少，position是哪一位。
// 举个例子吧，例如add(14)。14已经超出byte[0]的映射范围，在byte[1]范围之类。
// 那么怎么快速定位它的索引呢。
// 如果找到它的索引号，又怎么定位它的位置呢。Index(N)代表N的索引号，Position(N)代表N的所在的位置号。
// Index(N) = N/8 = N >> 3;
// Position(N) = N%8 = N & 0x07;

//BitMap for bitmap obj
type BitMap struct {
	bits    []byte
	bitSize uint64
	// 该 Bitmap 被设置为 1 的最大位置（方便遍历）
	maxpos uint64
}

//NewBitmapSize for create bitmap
func NewBitmapSize(size uint64) *BitMap {
	// remainder := size % 8
	// if remainder != 0 {
	// 	remainder = 1
	// }
	// println(size/8 + remainder)

	// println((size >> 3) + 1)
	return &BitMap{
		bits:    make([]byte, (size>>3)+1),
		bitSize: uint64(size),
	}
}

//SetBit 将 offset 位置的 bit 置为 value (0/1)
func (b *BitMap) SetBit(num uint64, value uint8) bool {
	// num/8得到byte[]的index
	index := num >> 3
	// num%8得到在byte[index]的位置
	pos := num & 0x07
	// println(num, index, pos, num&0x07 == num%8)
	if b.bitSize < num {
		return false
	}

	// println(b.bitSize, index, pos)

	if value == 0 {
		// &^ 清位, 将1左移pos后，那个位置自然就是1，然后对取反，再与当前值做&，即可清除当前的位置了.
		b.bits[index] &= ^(0x01 << pos)
	} else {
		//将1左移pos后，那个位置自然就是1，然后和以前的数据做|，这样，那个位置就替换成1了。
		b.bits[index] |= 0x01 << pos
		// fmt.Println(2222, b.bits[index])
		// 记录曾经设置为 1 的最大位置
		if b.maxpos < num {
			b.maxpos = num
		}
	}

	// spew.Dump(b.bits)
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
func (b *BitMap) GetBit(num uint64) uint8 {
	index, pos := num/8, num%8

	if b.bitSize < num {
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
