//+build ignore

package main

import (
	"dlib/dbus"
	"dlib/dbus/introspect"
	"encoding/json"
	"os"
)

func main() {
	conn, err := dbus.SystemBus()
	if err != nil {
		panic(err)
	}
	node, err := introspect.Call(conn.Object("org.freedesktop.UPower",
		"/org/freedesktop/UPower"))
	if err != nil {
		panic(err)
	}
	data, _ := json.MarshalIndent(node, "", "    ")
	os.Stdout.Write(data)
}
