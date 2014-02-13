/*This file is auto generate by dlib/dbus/proxyer. Don't edit it*/

package upower

import "dlib/dbus"
import "reflect"

type QoS struct {
	Path        dbus.ObjectPath
	core        *dbus.Object
	signal_chan chan *dbus.Signal
}

func (obj QoS) GetLatencyRequests() (requests [][]interface{}) {
	obj.core.Call("org.freedesktop.UPower.QoS.GetLatencyRequests", 0).Store(&requests)
	return
}

func (obj QoS) GetLatency(type_ string) (value int32) {
	obj.core.Call("org.freedesktop.UPower.QoS.GetLatency", 0, type_).Store(&value)
	return
}

func (obj QoS) CancelRequest(cookie uint32) {
	obj.core.Call("org.freedesktop.UPower.QoS.CancelRequest", 0, cookie).Store()
	return
}

func (obj QoS) RequestLatency(type_ string, value int32, persistent bool) (cookie uint32) {
	obj.core.Call("org.freedesktop.UPower.QoS.RequestLatency", 0, type_, value, persistent).Store(&cookie)
	return
}

func (obj QoS) SetMinimumLatency(type_ string, value int32) {
	obj.core.Call("org.freedesktop.UPower.QoS.SetMinimumLatency", 0, type_, value).Store()
	return
}

func (obj QoS) ConnectRequestsChanged(callback func()) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+string(obj.core.Path())+"', interface='org.freedesktop.UPower.QoS',sender='org.freedesktop.UPower',member='RequestsChanged'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range obj.signal_chan {
			if v.Name != "org.freedesktop.UPower.QoS.RequestsChanged" || 0 != len(v.Body) {
				continue
			}
			_ = reflect.TypeOf /*prevent compile error*/

			callback()
		}
	}()

}

func (obj QoS) ConnectLatencyChanged(callback func(string, int32)) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+string(obj.core.Path())+"', interface='org.freedesktop.UPower.QoS',sender='org.freedesktop.UPower',member='LatencyChanged'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range obj.signal_chan {
			if v.Name != "org.freedesktop.UPower.QoS.LatencyChanged" || 2 != len(v.Body) {
				continue
			}

			if reflect.TypeOf(v.Body[0]) != reflect.TypeOf((*string)(nil)).Elem() {
				continue
			}
			if reflect.TypeOf(v.Body[0]) != reflect.TypeOf((*int32)(nil)).Elem() {
				continue
			}

			callback(v.Body[0].(string), v.Body[1].(int32))
		}
	}()

}

func GetQoS(path string) *QoS {
	return &QoS{dbus.ObjectPath(path), getBus().Object("org.freedesktop.UPower", dbus.ObjectPath(path)), make(chan *dbus.Signal)}
}
