//+build ignore

package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	reg, err := regexp.Compile("(?:index:)([\\S\\s]+)(?:index)")
	//reg, err := regexp.Compile("(?:index:)([\\S\\s]+)")
	//reg, err := regexp.Compile("(?:index)([[:alnum:]\\s]+)(?:index)?")
	if err != nil {
		panic(err)
	}
	file, err := os.Open("./sound/sinks.log")
	if err != nil {
		panic(err)
	}
	b := make([]byte, 65536)
	n, err := file.Read(b)
	s := string(b[0:n]) + "1 sinks(s) available.\n\t* index31443helllo3\r\n23index"
	//loc := reg.FindIndex([]byte("1 sinks(s) available.\n\t* index31443helllo3\r\n23index"))
	fmt.Println(reg.FindAllString("1 sinks(s) available.\n\t* index31443helllo3\r\n23index", -1))
	fmt.Println(reg.FindAllString(s, -1))
	//fmt.Print(loc)
}
