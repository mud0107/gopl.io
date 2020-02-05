package main

import (
	"io"
	"log"
	"net"
	"os"
)

/*
	연습문제 8.3: netcat3에서 인터페이스 값  conn은 TCP 연결을 나타내는 구상타입 *.net.TCPConn이다.
				TCP 연결을은 CloseRead와 CloseWrite 메소드로 독립적으로 닫을 수 있는 두 부부분으로 구성돼 있다.
				연결의 쓰는 부분만 닫아서 표준 입력이 닫힌 후에도 프로그래밍이 reverb1서버의 마지막 에코를 표시할 때까지
				netcat3의 메인 고루틴을 수정하라
*/
func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		log.Println("done")
		done <- struct{}{} // signal the main goroutine
	}()
	mustCopy(conn, os.Stdin)
	if conn, ok := conn.(*net.TCPConn); ok {
		conn.CloseWrite()
	}
	<-done // wait for background goroutine to finish
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
