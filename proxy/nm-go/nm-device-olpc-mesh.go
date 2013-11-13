/*This file is auto generate by dlib/dbus/proxyer. Don't edit it*/

package nm

import "dlib/dbus"

type DeviceOlpcMesh struct {
	core        *dbus.Object
	signal_chan chan *dbus.Signal
}

func (obj DeviceOlpcMesh) ConnectPropertiesChanged(callback func(properties map[string]dbus.Variant)) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+obj.core.Path()+"', interface='org.freedesktop.NetworkManager.Device.OlpcMesh',sender='org.freedesktop.NetworkManager'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range obj.signal_chan {
			v = v
			/*callback(v.Body...)*/
		}
	}()

}

func (obj *DeviceOlpcMesh) SetHwAddress(HwAddress string) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.Device.OlpcMesh", "HwAddress", HwAddress)
}
func (obj DeviceOlpcMesh) GetHwAddress() (ret string) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.Device.OlpcMesh", "HwAddress").Store(&r)
	if err == nil && r.Signature().String() == "s" {
		return r.Value().(string)
	} else {
		panic(err)
	}
	return
}

func (obj *DeviceOlpcMesh) SetCompanion(Companion dbus.ObjectPath) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.Device.OlpcMesh", "Companion", Companion)
}
func (obj DeviceOlpcMesh) GetCompanion() (ret dbus.ObjectPath) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.Device.OlpcMesh", "Companion").Store(&r)
	if err == nil && r.Signature().String() == "o" {
		return r.Value().(dbus.ObjectPath)
	} else {
		panic(err)
	}
	return
}

func (obj *DeviceOlpcMesh) SetActiveChannel(ActiveChannel uint32) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.Device.OlpcMesh", "ActiveChannel", ActiveChannel)
}
func (obj DeviceOlpcMesh) GetActiveChannel() (ret uint32) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.Device.OlpcMesh", "ActiveChannel").Store(&r)
	if err == nil && r.Signature().String() == "u" {
		return r.Value().(uint32)
	} else {
		panic(err)
	}
	return
}

func GetDeviceOlpcMesh(path string) *DeviceOlpcMesh {
	return &DeviceOlpcMesh{getBus().Object("org.freedesktop.NetworkManager", dbus.ObjectPath(path)), make(chan *dbus.Signal)}
}
