/*This file is auto generate by dlib/dbus/proxyer. Don't edit it*/

package nm

import "dlib/dbus"

type DeviceBridge struct {
	core        *dbus.Object
	signal_chan chan *dbus.Signal
}

func (obj DeviceBridge) ConnectPropertiesChanged(callback func(properties map[string]dbus.Variant)) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+obj.core.Path()+"', interface='org.freedesktop.NetworkManager.Device.Bridge',sender='org.freedesktop.NetworkManager'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range obj.signal_chan {
			v = v
			/*callback(v.Body...)*/
		}
	}()

}

func (obj *DeviceBridge) SetHwAddress(HwAddress string) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.Device.Bridge", "HwAddress", HwAddress)
}
func (obj DeviceBridge) GetHwAddress() (ret string) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.Device.Bridge", "HwAddress").Store(&r)
	if err == nil && r.Signature().String() == "s" {
		return r.Value().(string)
	} else {
		panic(err)
	}
	return
}

func (obj *DeviceBridge) SetCarrier(Carrier bool) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.Device.Bridge", "Carrier", Carrier)
}
func (obj DeviceBridge) GetCarrier() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.Device.Bridge", "Carrier").Store(&r)
	if err == nil && r.Signature().String() == "b" {
		return r.Value().(bool)
	} else {
		panic(err)
	}
	return
}

func (obj *DeviceBridge) SetSlaves(Slaves []dbus.ObjectPath) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.Device.Bridge", "Slaves", Slaves)
}
func (obj DeviceBridge) GetSlaves() (ret []dbus.ObjectPath) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.Device.Bridge", "Slaves").Store(&r)
	if err == nil && r.Signature().String() == "ao" {
		return r.Value().([]dbus.ObjectPath)
	} else {
		panic(err)
	}
	return
}

func GetDeviceBridge(path string) *DeviceBridge {
	return &DeviceBridge{getBus().Object("org.freedesktop.NetworkManager", dbus.ObjectPath(path)), make(chan *dbus.Signal)}
}
