package main

import (
	"fmt"
	"sync"
)

// go generator.
var do_once_range sync.Once

func grange(start int, after int) chan int {
	var ch = make(chan int, 1)
	if after < start {
		start, after = after, start
	}
	f1 := func() {

		go func(start int, after int) {
			for i := start; i < after; i++ {
				ch <- i
			}
			close(ch)
		}(start, after)
	}

	do_once_range.Do(f1)
	return ch
}

func test_grange() {
	fmt.Println("vim-go")
	fmt.Println("[TEST][GRANGE] 0 -5")
	for i := range grange(0, 5) {
		fmt.Println("[TEST][GRANGE][STEP]", i)
	}
}
