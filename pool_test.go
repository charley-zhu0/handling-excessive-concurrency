/*
 * @Author: charley zhu
 * @Date: 2023-10-24 09:32:49
 * @LastEditTime: 2023-10-24 09:53:34
 * @LastEditors: charley zhu
 * @Description:
 */
package main

import (
	"sync"
	"testing"
	"time"
)

func BenchmarkBase(b *testing.B) {
	testNums := 100000

	wg := sync.WaitGroup{}
	for i := 0; i < testNums; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			testFunc()
		}()
	}
	wg.Wait()
	b.Log("all done")
}

func BenchmarkChannel(b *testing.B) {
	testNums := 100000

	wg := sync.WaitGroup{}
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
	b.Log("all done, cost: ", cost)
}
