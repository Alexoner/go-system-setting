// +build ignore

package main

import (
	//"fmt"
	//"dlib/gio-2.0"
	"github.com/guelfey/go.dbus"
	"time"
)

const DBUS_SET_METHOD = "org.freedesktop.DBus.Properties.Set"

func main() {
	conn, err := dbus.SessionBus()
	if err != nil {
		panic(err)
	}

	obj := conn.Object("com.deepin.daemon.Audio", "/com/deepin/daemon/Audio/Sink82")
	obj.Call("org.freedesktop.DBus.Properties.Set", 0, "com.deepin.daemon.Audio.Sink", "Mute", dbus.MakeVariant(true))
	obj.Call(DBUS_SET_METHOD, 0, "com.deepin.daemon.Audio.Sink", "Volume", dbus.MakeVariant(uint32(0)))
	time.Sleep(1000)
	obj.Call("org.freedesktop.DBus.Properties.Set", 0, "com.deepin.daemon.Audio.Sink", "Mute", dbus.MakeVariant(false))
	obj.Call(DBUS_SET_METHOD, 0, "com.deepin.daemon.Audio.Sink",
		"Volume", dbus.MakeVariant(uint32(33)))
	obj.Call(DBUS_SET_METHOD, 0, "com.deepin.daemon.Audio.Sink",
		"Balance", dbus.MakeVariant(float64(-1)))
	time.Sleep(1000)
	obj.Call(DBUS_SET_METHOD, 0, "com.deepin.daemon.Audio.Sink",
		"Balance", dbus.MakeVariant(float64(0)))
	audioObj := conn.Object("com.deepin.daemon.Audio", "/com/deepin/daemon/Audio")
	audioObj.Call(DBUS_SET_METHOD, 0, "com.deepin.daemon.Audio",
		"DefaultSource", dbus.MakeVariant(int32(1)))
	audioObj.Call(DBUS_SET_METHOD, 0, "com.deepin.daemon.Audio",
		"DefaultSource", dbus.MakeVariant(int32(0)))
}
