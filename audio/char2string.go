// +build ignore

package main

//#include <stdio.h>
// const char a[16]="hello,world";
import "C"
import "fmt"

import "unsafe"

func main() {
	var b *C.char = (*C.char)(unsafe.Pointer(&C.a[0]))
	//var b *C.char = (*C.char)(&C.a[0])
	fmt.Println(C.GoString(b))
	fmt.Println(C.GoString(&C.a[0]))
}
