package main

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestBitmapSort(t *testing.T) {
	bmap := NewBitmapSize(20)
	spew.Dump(bmap.bits)

	original := [5]uint64{4, 6, 3, 1, 7}
	expected := [5]uint64{1, 3, 4, 6, 7}
	actual := [5]uint64{}

	for index := 0; index < 1000; index++ {

	}

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
	size := uint64(10)
	bmap := NewBitmapSize(size)
	spew.Dump(bmap.bits)

	count := 100
	a := make([]uint64, 0, count)
	for index := 0; index < count; index++ {
		a = append(a, uint64(index))
	}

	for _, num := range a {
		if num%2 == 0 {
			bmap.SetBit(num, 1)
		} else {
			bmap.SetBit(num, 0)
		}
	}

	println(bmap.Contain(0), bmap.Contain(1))
}
