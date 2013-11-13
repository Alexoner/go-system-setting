/*This file is auto generate by dlib/dbus/proxyer. Don't edit it*/

package nm

import "dlib/dbus"

type DeviceWiMax struct {
	core        *dbus.Object
	signal_chan chan *dbus.Signal
}

func (obj DeviceWiMax) GetNspList() (nsps []dbus.ObjectPath) {
	obj.core.Call("org.freedesktop.NetworkManager.Device.WiMax.GetNspList", 0).Store(&nsps)
	return
}

func (obj DeviceWiMax) ConnectPropertiesChanged(callback func(properties map[string]dbus.Variant)) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+obj.core.Path()+"', interface='org.freedesktop.NetworkManager.Device.WiMax',sender='org.freedesktop.NetworkManager'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range obj.signal_chan {
			v = v
			/*callback(v.Body...)*/
		}
	}()

}

func (obj DeviceWiMax) ConnectNspAdded(callback func(nsp dbus.ObjectPath)) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+obj.core.Path()+"', interface='org.freedesktop.NetworkManager.Device.WiMax',sender='org.freedesktop.NetworkManager'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range obj.signal_chan {
			v = v
			/*callback(v.Body...)*/
		}
	}()

}

func (obj DeviceWiMax) ConnectNspRemoved(callback func(nsp dbus.ObjectPath)) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+obj.core.Path()+"', interface='org.freedesktop.NetworkManager.Device.WiMax',sender='org.freedesktop.NetworkManager'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range obj.signal_chan {
			v = v
			/*callback(v.Body...)*/
		}
	}()

}

func (obj *DeviceWiMax) SetHwAddress(HwAddress string) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.Device.WiMax", "HwAddress", HwAddress)
}
func (obj DeviceWiMax) GetHwAddress() (ret string) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.Device.WiMax", "HwAddress").Store(&r)
	if err == nil && r.Signature().String() == "s" {
		return r.Value().(string)
	} else {
		panic(err)
	}
	return
}

func (obj *DeviceWiMax) SetCenterFrequency(CenterFrequency uint32) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.Device.WiMax", "CenterFrequency", CenterFrequency)
}
func (obj DeviceWiMax) GetCenterFrequency() (ret uint32) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.Device.WiMax", "CenterFrequency").Store(&r)
	if err == nil && r.Signature().String() == "u" {
		return r.Value().(uint32)
	} else {
		panic(err)
	}
	return
}

func (obj *DeviceWiMax) SetRssi(Rssi int32) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.Device.WiMax", "Rssi", Rssi)
}
func (obj DeviceWiMax) GetRssi() (ret int32) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.Device.WiMax", "Rssi").Store(&r)
	if err == nil && r.Signature().String() == "i" {
		return r.Value().(int32)
	} else {
		panic(err)
	}
	return
}

func (obj *DeviceWiMax) SetCinr(Cinr int32) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.Device.WiMax", "Cinr", Cinr)
}
func (obj DeviceWiMax) GetCinr() (ret int32) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.Device.WiMax", "Cinr").Store(&r)
	if err == nil && r.Signature().String() == "i" {
		return r.Value().(int32)
	} else {
		panic(err)
	}
	return
}

func (obj *DeviceWiMax) SetTxPower(TxPower int32) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.Device.WiMax", "TxPower", TxPower)
}
func (obj DeviceWiMax) GetTxPower() (ret int32) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.Device.WiMax", "TxPower").Store(&r)
	if err == nil && r.Signature().String() == "i" {
		return r.Value().(int32)
	} else {
		panic(err)
	}
	return
}

func (obj *DeviceWiMax) SetBsid(Bsid string) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.Device.WiMax", "Bsid", Bsid)
}
func (obj DeviceWiMax) GetBsid() (ret string) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.Device.WiMax", "Bsid").Store(&r)
	if err == nil && r.Signature().String() == "s" {
		return r.Value().(string)
	} else {
		panic(err)
	}
	return
}

func (obj *DeviceWiMax) SetActiveNsp(ActiveNsp dbus.ObjectPath) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.Device.WiMax", "ActiveNsp", ActiveNsp)
}
func (obj DeviceWiMax) GetActiveNsp() (ret dbus.ObjectPath) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.Device.WiMax", "ActiveNsp").Store(&r)
	if err == nil && r.Signature().String() == "o" {
		return r.Value().(dbus.ObjectPath)
	} else {
		panic(err)
	}
	return
}

func GetDeviceWiMax(path string) *DeviceWiMax {
	return &DeviceWiMax{getBus().Object("org.freedesktop.NetworkManager", dbus.ObjectPath(path)), make(chan *dbus.Signal)}
}
