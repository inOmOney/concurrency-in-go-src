package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}
func main() {

	newRandStream := func(done <-chan interface{}) <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream closure exited.") // <1>
			defer close(randStream)

			for {
				select {
				case <-done:
					return
				default: // 可以缩短select
				}
				randStream <- rand.Int()
				//Do something interesting
			}
		}()
		return randStream
	}

	done := make(chan interface{})
	randStream := newRandStream(done)
	for i := 0; i < 3; i++ {
		fmt.Printf("%d\n", <-randStream)
	}
	close(done)
	<-randStream
	fmt.Println("Done")
}
