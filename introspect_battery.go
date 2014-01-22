//+build ignore

package main

import (
	"dlib/dbus"
	"fmt"
)

func main() {
	conn, err := dbus.SystemBus()

	if err != nil {
		panic(err)
	}

	obj := conn.Object("org.freedesktop.UPower",
		"/org/freedesktop/UPower/devices/battery_BAT0")

	variant, err := obj.GetProperty("org.freedesktop.UPower.Device.Percentage")
	if err != nil {
		panic(err)
	}

	fmt.Println("percentage: ", variant.Value())
}
