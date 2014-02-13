//+build ignore

package main

import "dlib/dbus"

func main() {
	conn, err := dbus.SessionBus()

	if err != nil {
		panic(err)
	}

	obj := conn.Object("org.freedesktop.Notifications",
		"/org/freedesktop/Notifications")

	call := obj.Call("org.freedesktop.Notifications.Notify", 6, "dlib", uint32(6),
		"hello world", "Test", "This is a test of the DBus bindings for go.",
		[]string{},
		map[string]dbus.Variant{}, int32(5006))
	if call.Err != nil {
		panic(call.Err)
	}
}
