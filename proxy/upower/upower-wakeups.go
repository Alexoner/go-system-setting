/*This file is auto generate by dlib/dbus/proxyer. Don't edit it*/

package upower

import "dlib/dbus"
import "reflect"

type Wakeups struct {
	Path        dbus.ObjectPath
	core        *dbus.Object
	signal_chan chan *dbus.Signal
}

func (obj Wakeups) GetData() (data [][]interface{}) {
	obj.core.Call("org.freedesktop.UPower.Wakeups.GetData", 0).Store(&data)
	return
}

func (obj Wakeups) GetTotal() (value uint32) {
	obj.core.Call("org.freedesktop.UPower.Wakeups.GetTotal", 0).Store(&value)
	return
}

func (obj Wakeups) ConnectDataChanged(callback func()) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+string(obj.core.Path())+"', interface='org.freedesktop.UPower.Wakeups',sender='org.freedesktop.UPower',member='DataChanged'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range obj.signal_chan {
			if v.Name != "org.freedesktop.UPower.Wakeups.DataChanged" || 0 != len(v.Body) {
				continue
			}
			_ = reflect.TypeOf /*prevent compile error*/

			callback()
		}
	}()

}

func (obj Wakeups) ConnectTotalChanged(callback func(uint32)) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+string(obj.core.Path())+"', interface='org.freedesktop.UPower.Wakeups',sender='org.freedesktop.UPower',member='TotalChanged'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range obj.signal_chan {
			if v.Name != "org.freedesktop.UPower.Wakeups.TotalChanged" || 1 != len(v.Body) {
				continue
			}

			if reflect.TypeOf(v.Body[0]) != reflect.TypeOf((*uint32)(nil)).Elem() {
				continue
			}

			callback(v.Body[0].(uint32))
		}
	}()

}

func (obj *Wakeups) SetHasCapability(HasCapability bool) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Wakeups", "HasCapability", HasCapability)
}
func (obj Wakeups) GetHasCapability() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Wakeups", "HasCapability").Store(&r)
	if err == nil && r.Signature().String() == "b" {
		return r.Value().(bool)
	} else {
		panic(err)
	}
	return
}

func GetWakeups(path string) *Wakeups {
	return &Wakeups{dbus.ObjectPath(path), getBus().Object("org.freedesktop.UPower", dbus.ObjectPath(path)), make(chan *dbus.Signal)}
}
