/*This file is auto generate by dlib/dbus/proxyer. Don't edit it*/

package upower

import "dlib/dbus"

type wakeups struct {
	core        *dbus.Object
	signal_chan chan *dbus.Signal
}

func (obj wakeups) GetData() (data [][]interface{}) {
	obj.core.Call("org.freedesktop.UPower.Wakeups.GetData", 0).Store(&data)
	return
}

func (obj wakeups) GetTotal() (value uint32) {
	obj.core.Call("org.freedesktop.UPower.Wakeups.GetTotal", 0).Store(&value)
	return
}

func (obj wakeups) ConnectDataChanged(callback func()) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+obj.core.Path()+"', interface='org.freedesktop.UPower.Wakeups',sender='org.freedesktop.UPower'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range obj.signal_chan {
			v = v
			/*callback(v.Body...)*/
		}
	}()

}

func (obj wakeups) ConnectTotalChanged(callback func(uint32)) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+obj.core.Path()+"', interface='org.freedesktop.UPower.Wakeups',sender='org.freedesktop.UPower'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range obj.signal_chan {
			v = v
			/*callback(v.Body...)*/
		}
	}()

}

func (obj *wakeups) SetHasCapability(HasCapability bool) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Wakeups", "HasCapability", HasCapability)
}
func (obj wakeups) GetHasCapability() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Wakeups", "HasCapability").Store(&r)
	if err == nil && r.Signature().String() == "b" {
		return r.Value().(bool)
	} else {
		panic(err)
	}
	return
}

func Getwakeups(path string) *wakeups {
	return &wakeups{getBus().Object("org.freedesktop.UPower", dbus.ObjectPath(path)), make(chan *dbus.Signal)}
}
