//+build ignore

package main

import "upower"
import "fmt"

func main() {
	a := upower.Getupower("/org/freedesktop/UPower")
	//b := upower.Getdevice("org/freedesktop/UPower/devices/battery_BAT0")
	//fmt.Println("Model: ", upower.Getdevice("/org/freedesktop/UPower/devices/battery_BAT0").GetModel())
	fmt.Println("can suspend: ", b.GetModel())
}
