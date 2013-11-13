//+build ignore

package main

import "./upower-go"
import "fmt"

func main() {
	a := upower.Getupower("/org/freedesktop/UPower")
	fmt.Println("can suspend: ", a.GetCanSuspend())
}
