//+build ignore

package main

import (
	"bufio"
	//"io"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	conn, err := net.Dial("unix", "/run/user/1000/pulse/cli")

	if err != nil {
		panic(err)
	}
	lineRead := make([]byte, 65536)
	n, err := conn.Read(lineRead)
	if err != nil {
		println(err.Error())
	}
	fmt.Printf(string(lineRead[0 : n-2]))

	i := 0
	for {
		bio := bufio.NewReader(os.Stdin)
		lineWrite, _, err := bio.ReadLine()

		if err != nil {
			if err.Error() == "EOF" {
				os.Exit(0)
			}
		}

		i = i + 1

		_, err = conn.Write(lineWrite)
		conn.Write([]byte("\r\n"))

		if err != nil {
			println(err.Error())
			break
		}

		n, err := conn.Read(lineRead)
		//println(n, ": \n")
		if err != nil {
			println(err.Error())
		}
		fmt.Printf(string(lineRead[0 : n-2]))
		time.Sleep(0)
	}
}
