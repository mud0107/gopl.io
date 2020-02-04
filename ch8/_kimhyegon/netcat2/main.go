package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	for {
		conn, err := net.Dial("tcp", "localhost:8000")
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		go mustCopy(os.Stdout, conn) // netcat1 과 다른부분

	}
}

func mustCopy(dst io.Writer, src io.Reader) {

	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
