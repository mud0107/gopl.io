package main

import (
	"fmt"
	"time"
)

/*
	8.1 고루틴은 호출 함수 앞에 go 키워드 사용.
	예)
		f() // 일반 함수 호출
		go f() 고루틴
		동시성과 병렬성의 차이
		동시성은 코드이고 병렬성의 실제 물리적이 코어에 실행된는 것
*/

// 샐행방법 : go run main.go

func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}

	return fib(x-1) + fib(x-2)
}
