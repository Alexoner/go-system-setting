/*This file is auto generate by dlib/dbus/proxyer. Don't edit it*/

package upower
import "dlib/dbus"

type qos struct {
	core *dbus.Object
	signal_chan chan *dbus.Signal
}


func (obj qos) GetLatencyRequests () (requests [][]interface {}) {
	obj.core.Call("org.freedesktop.UPower.QoS.GetLatencyRequests", 0).Store(&requests)
	return
}

func (obj qos) GetLatency (type string) (value int32) {
	obj.core.Call("org.freedesktop.UPower.QoS.GetLatency", 0, type).Store(&value)
	return
}

func (obj qos) CancelRequest (cookie uint32) () {
	obj.core.Call("org.freedesktop.UPower.QoS.CancelRequest", 0, cookie).Store()
	return
}

func (obj qos) RequestLatency (type string,value int32,persistent bool) (cookie uint32) {
	obj.core.Call("org.freedesktop.UPower.QoS.RequestLatency", 0, type, value, persistent).Store(&cookie)
	return
}

func (obj qos) SetMinimumLatency (type string,value int32) () {
	obj.core.Call("org.freedesktop.UPower.QoS.SetMinimumLatency", 0, type, value).Store()
	return
}



func (obj qos) ConnectRequestsChanged(callback func()) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+obj.core.Path()+"', interface='org.freedesktop.UPower.QoS',sender='org.freedesktop.UPower'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range(obj.signal_chan) {
			v = v
			/*callback(v.Body...)*/
		}
	}()

}

func (obj qos) ConnectLatencyChanged(callback func( string, int32)) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+obj.core.Path()+"', interface='org.freedesktop.UPower.QoS',sender='org.freedesktop.UPower'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range(obj.signal_chan) {
			v = v
			/*callback(v.Body...)*/
		}
	}()

}




func Getqos(path string) *qos {
	return  &qos{getBus().Object("org.freedesktop.UPower", dbus.ObjectPath(path)),make(chan *dbus.Signal)}
}

