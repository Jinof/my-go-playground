package goer

import (
	"fmt"
	"sync"
)

func SeqPrint() {
	wg := sync.WaitGroup{}
	wg.Add(3)

	nums := []int{1, 2, 3}
	for _, v := range nums {
		ch := make(chan int)
		go func(ch <-chan int) {
			key := <-ch
			fmt.Println(key)
		}(ch)
		ch <- v
	}
	wg.Wait()
}
