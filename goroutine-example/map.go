package main

import (
	"fmt"
	"sync"
)

// map 的并发安全

var mp = make(map[int]int)
var mwg sync.WaitGroup
var smp sync.Map

func main() {
	for i := 0; i < 10; i++ {
		mwg.Add(1)
		go func() {
			// mp[i] = i
			smp.Store(i, i)
			mwg.Done()
		}()
	}

	mwg.Wait()
	// fmt.Println(mp)
	smp.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})
}
