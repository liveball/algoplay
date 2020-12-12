package main

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestBitmapSort(t *testing.T) {
	bmap := NewBitmapSize(20)
	spew.Dump(bmap.bits)

	original := [5]uint64{4, 6, 3, 1, 7}
	expected := [5]uint64{1, 3, 4, 6, 7}
	actual := [5]uint64{}

	for _, num := range original {
		bmap.SetBit(num, 1)
	}

	var i uint64
	var num, maxpos uint64 = 0, bmap.Maxpos() + 1
	for ; num < maxpos; num++ {
		if bmap.GetBit(num) == 1 {
			actual[i] = num
			i++
		}
	}

	if expected != actual {
		t.Errorf("expected:%#v, actual:%#v", expected, actual)
	}
}

func TestBitmap(t *testing.T) {
	size := uint64(9)
	bmap := NewBitmapSize(size)

	count := 9
	a := make([]uint64, 0, count)
	for index := 0; index < count; index++ {
		a = append(a, uint64(index))
	}

	for _, num := range a {
		if num%2 == 0 {
			bmap.SetBit(num, 0)
		} else {
			bmap.SetBit(num, 1)
		}
	}

	// println(bmap.bits[0]) //128+64+32+16+8+4+2+1 =255

	// for k, v := range bmap.bits {
	// 	println(k, v)
	// }

	println(bmap.GetBit(3), bmap.Contain(1))
}


func TestAttr(t *testing.T) {
	a := setAttr(0, 1, 1)
	b := setAttr(a, 1, 2)
	c := attr(a, 1)
	d := attr(b, 2)
	fmt.Println(a, b, c, d)
	fmt.Println(attr(6, 3), attr(6, 2), attr(6, 1), attr(6, 0))
}
