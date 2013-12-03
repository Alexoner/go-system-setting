//+build ignore

package main

import "upower"
import "fmt"

func main() {
	a := upower.GetUpower("/org/freedesktop/UPower")
	b := upower.GetDevice("/org/freedesktop/UPower/devices/battery_BAT0")
	fmt.Println("can suspend: ", a.GetCanSuspend())
	fmt.Println("Model: ", upower.GetDevice("/org/freedesktop/UPower/devices/battery_BAT0").GetModel())
	fmt.Println("Model: ", b.GetModel(), "Vendor: ", b.GetVendor())
}
