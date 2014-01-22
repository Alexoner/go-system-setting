//+build ignore

package main

import (
	"dlib/dbus"
	"fmt"
	"os"
)

func main() {
	conn, err := dbus.SystemBus()

	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to connect to the system bus: ", err)
		os.Exit(1)
	}

	var s bool
	obj := conn.Object("org.freedesktop.UPower",
		"/org/freedesktop/UPower")
	/*err =obj.Call("org.freedesktop.UPower.SuspendAllowed",6).Store(&s)

	if err != nil {
		panic (err)
	}*/

	names := conn.Names()
	for _, v := range names {
		fmt.Println(v)
	}

	call := obj.Call("org.freedesktop.UPower.SuspendAllowed", 6)

	if call.Err != nil {
		panic(err)
		os.Exit(1)
	}

	call.Store(&s)

	fmt.Println("call UPower SuspendAllowed: ", s)

	variant, err := obj.GetProperty("org.freedesktop.UPower.LidIsPresent")
	if err != nil {
		panic(err)
		os.Exit(1)
	}

	fmt.Println("getproperty UPower LidIsPresent: ", variant.Value())
}
