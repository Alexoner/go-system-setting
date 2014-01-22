//+build ignore

package main

import "fmt"
import "net"

func main() {
	conn, err := net.Dial("tcp", "www.google.com.hk:80")

	if err != nil {
		panic(err)
	}

	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	b := make([]byte, 10000)
	n, err := conn.Read(b)
	fmt.Println(n, "\n", string(b))
}
