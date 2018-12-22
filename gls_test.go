package gls

import (
	"fmt"
	"sync"
	"testing"
)

func TestGls(t *testing.T) {
	wg := sync.WaitGroup{}
	for i := 0; i < 12; i++ {
		wg.Add(1)
		go func(i int) {
			defer func() {
				wg.Done()
				//Delete()
			}()
			Set(i, fmt.Sprintf("routine #%d", i))
			fmt.Printf("gid:%v, index:%v, v:%v\n", GoId(), i, Get(i))
		}(i)
	}
	wg.Wait()
	fmt.Printf("=---=-=-=-=-=-=-=-\n")
	for shard, mp := range globalMap {
		mp.Range(func(key, value interface{}) bool {
			if imp, ok := value.(*sync.Map); ok {
				imp.Range(func(key2, value2 interface{}) bool {
					fmt.Printf("shard:%v gid:%v index:%v, v:%v\n", shard, key, key2, value2)
					return true
				})
			}
			return true
		})
	}
}
