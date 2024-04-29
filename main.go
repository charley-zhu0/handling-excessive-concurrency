/*
 * @Author: charley zhu
 * @Date: 2023-10-23 13:25:08
 * @LastEditTime: 2023-10-24 10:00:36
 * @LastEditors: charley zhu
 * @Description:
 */
package main

import (
	"fmt"
	"sync"
	"time"
)

func testFunc() []int {
	// do something
	time.Sleep(time.Second)
	temp := make([]int, 10)
	return temp
	// percent, _ := cpu.Percent(time.Second, false)
	// memInfo, _ := mem.VirtualMemory()
	// fmt.Printf("cpu: %v, mem: %v\n", percent[0], memInfo.UsedPercent)
}

func main() {
	// start 1000000 goroutines
	testNums := 100000

	wg := sync.WaitGroup{}
	// for i := 0; i < testNums; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()
	// 		testFunc()
	// 	}()
	// }
	// wg.Wait()
	// fmt.Println("all done")

	// use channel
	start := time.Now()
	ch := make(chan struct{}, 1000)
	for i := 0; i < testNums; i++ {
		ch <- struct{}{}
		wg.Add(1)
		go func() {
			defer func() { <-ch }()
			defer wg.Done()
			testFunc()
		}()
	}
	wg.Wait()
	cost := time.Since(start)
	fmt.Printf("all done, cost: %v\n", cost)
}
