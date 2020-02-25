package main

import (
	"fmt"
	"testing"
)

func Test_List(t *testing.T) {
	list := NewList()

	fmt.Println("rpush=>start")
	list.RPush("a")
	list.RPush("b")
	list.RPush("c")
	list.RPush("d")
	fmt.Println("rpush=>end")

	fmt.Println(list.Len())

	fmt.Println(list.Index(1))

	fmt.Println("lpop<=start")
	fmt.Println(list.LPop().value)
	fmt.Println(list.LPop().value)
	fmt.Println(list.LPop().value)
	fmt.Println(list.LPop().value)
	fmt.Println("lpop<=end")

	nodes := list.Range(0, list.Len())
	fmt.Println(nodes)

	list.RPush("1")
	list.RPush("2")
	list.RPush("3")

	nodes2 := list.Range(0, list.Len())
	for _, node := range nodes2 {
		fmt.Println("node=>", node)
	}
}
