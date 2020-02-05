package main

import "fmt"

/*
	100개까지 돌고 naturals 채널을 닫는다.
	전달 하는 Squarer 고루틴을 그 개수 만큰 일고 squares 채널을 닫는다.
	프린터 하는 루프는 버퍼에 있는 값까지 앍고 끝난다.
*/
func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// Squarer
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// Printer (in main goroutine)
	for x := range squares {
		fmt.Println(x)
	}
}
