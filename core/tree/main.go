package main

import (
	"fmt"
	"time"

	"algoplay/core/tree/bplustree"
)

func main() {
	bplustree_insert()
}

func bplustree_insert() {
	testCount := 1000000
	bt := bplustree.NewBTree()
	start := time.Now()
	for i := testCount; i > 0; i-- {
		bt.Insert(i, "")
	}

	fmt.Println(time.Now().Sub(start))
}
