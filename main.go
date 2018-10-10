package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/liveball/algoplay/common"
)

//GOMAXPROCS=1
//go build -o main main.go && GODEBUG=schedtrace=1000,scheddetail=1 ./main

func main() {

	pool := common.New(2)

	var wg sync.WaitGroup
	wg.Add(1)
	pool.Go(func() {
		fmt.Println("hello alg")
		wg.Done()
	})
	wg.Wait()
	pool.Close() //让工作池停止工作， 等待所有现有的工作完成

	time.Sleep(1 * time.Minute)
}
