package goer

import (
	"fmt"
	"sync"
	"time"
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

func SendRecive() {
	syncChan := make(chan int, 3)
	chan1 := make(chan struct{}, 1)
	chan2 := make(chan struct{}, 2)
	go func() {
		nums := []int{1, 2, 3, 4, 5}
		for _, v := range nums {
			syncChan <- v
			if v == 3 {
				chan1 <- struct{}{}
				fmt.Println("[sender] send a sync signal")
			}
			fmt.Println("[sender] send value", v)
		}
		time.Sleep(2 * time.Second)
		close(syncChan)
		chan2 <- struct{}{}
	}()
	go func() {
		<-chan1
		fmt.Println("[reciver] Recived a sync signal from sender")

		for {
			if v, ok := <-syncChan; ok {
				fmt.Println("[reciver] recive value", v)
			} else {
				break
			}
		}
		chan2 <- struct{}{}
	}()
	<-chan2
	<-chan2
}
