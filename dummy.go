package main

import "fmt"

func dummy() {
	ch := make(chan int, 10)
	defer close(ch)
	for i := 0; i < 20; i++ {
		ch <- i
	}
	for i := 0; i < 20; i++ {
		fmt.Println(<-ch)
	}

}
