/*This file is auto generate by dlib/dbus/proxyer. Don't edit it*/

package nm

import "dlib/dbus"

type DeviceAdsl struct {
	core        *dbus.Object
	signal_chan chan *dbus.Signal
}

func (obj DeviceAdsl) ConnectPropertiesChanged(callback func(properties map[string]dbus.Variant)) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+obj.core.Path()+"', interface='org.freedesktop.NetworkManager.Device.Adsl',sender='org.freedesktop.NetworkManager'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range obj.signal_chan {
			v = v
			/*callback(v.Body...)*/
		}
	}()

}

func (obj *DeviceAdsl) SetCarrier(Carrier bool) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.Device.Adsl", "Carrier", Carrier)
}
func (obj DeviceAdsl) GetCarrier() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.Device.Adsl", "Carrier").Store(&r)
	if err == nil && r.Signature().String() == "b" {
		return r.Value().(bool)
	} else {
		panic(err)
	}
	return
}

func GetDeviceAdsl(path string) *DeviceAdsl {
	return &DeviceAdsl{getBus().Object("org.freedesktop.NetworkManager", dbus.ObjectPath(path)), make(chan *dbus.Signal)}
}
