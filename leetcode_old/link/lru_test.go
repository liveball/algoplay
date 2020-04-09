package link

import (
	"fmt"
	"testing"
)

func Test_lru(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1)) // 返回  1

	cache.Put(3, 3)           // 该操作会使得密钥 2 作废
	fmt.Println(cache.Get(2)) // 返回 -1 (未找到)

	cache.Put(4, 4)           // 该操作会使得密钥 1 作废
	fmt.Println(cache.Get(1)) // 返回 -1 (未找到)
	fmt.Println(cache.Get(3)) // 返回  3
	fmt.Println(cache.Get(4)) // 返回  4
}
