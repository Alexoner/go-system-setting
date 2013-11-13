/*This file is auto generate by dlib/dbus/proxyer. Don't edit it*/

package nm

import "dlib/dbus"

type AccessPoint struct {
	core        *dbus.Object
	signal_chan chan *dbus.Signal
}

func (obj AccessPoint) ConnectPropertiesChanged(callback func(properties map[string]dbus.Variant)) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+obj.core.Path()+"', interface='org.freedesktop.NetworkManager.AccessPoint',sender='org.freedesktop.NetworkManager'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range obj.signal_chan {
			v = v
			/*callback(v.Body...)*/
		}
	}()

}

func (obj *AccessPoint) SetFlags(Flags uint32) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.AccessPoint", "Flags", Flags)
}
func (obj AccessPoint) GetFlags() (ret uint32) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.AccessPoint", "Flags").Store(&r)
	if err == nil && r.Signature().String() == "u" {
		return r.Value().(uint32)
	} else {
		panic(err)
	}
	return
}

func (obj *AccessPoint) SetWpaFlags(WpaFlags uint32) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.AccessPoint", "WpaFlags", WpaFlags)
}
func (obj AccessPoint) GetWpaFlags() (ret uint32) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.AccessPoint", "WpaFlags").Store(&r)
	if err == nil && r.Signature().String() == "u" {
		return r.Value().(uint32)
	} else {
		panic(err)
	}
	return
}

func (obj *AccessPoint) SetRsnFlags(RsnFlags uint32) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.AccessPoint", "RsnFlags", RsnFlags)
}
func (obj AccessPoint) GetRsnFlags() (ret uint32) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.AccessPoint", "RsnFlags").Store(&r)
	if err == nil && r.Signature().String() == "u" {
		return r.Value().(uint32)
	} else {
		panic(err)
	}
	return
}

func (obj *AccessPoint) SetSsid(Ssid []uint8) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.AccessPoint", "Ssid", Ssid)
}
func (obj AccessPoint) GetSsid() (ret []uint8) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.AccessPoint", "Ssid").Store(&r)
	if err == nil && r.Signature().String() == "ay" {
		return r.Value().([]uint8)
	} else {
		panic(err)
	}
	return
}

func (obj *AccessPoint) SetFrequency(Frequency uint32) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.AccessPoint", "Frequency", Frequency)
}
func (obj AccessPoint) GetFrequency() (ret uint32) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.AccessPoint", "Frequency").Store(&r)
	if err == nil && r.Signature().String() == "u" {
		return r.Value().(uint32)
	} else {
		panic(err)
	}
	return
}

func (obj *AccessPoint) SetHwAddress(HwAddress string) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.AccessPoint", "HwAddress", HwAddress)
}
func (obj AccessPoint) GetHwAddress() (ret string) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.AccessPoint", "HwAddress").Store(&r)
	if err == nil && r.Signature().String() == "s" {
		return r.Value().(string)
	} else {
		panic(err)
	}
	return
}

func (obj *AccessPoint) SetMode(Mode uint32) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.AccessPoint", "Mode", Mode)
}
func (obj AccessPoint) GetMode() (ret uint32) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.AccessPoint", "Mode").Store(&r)
	if err == nil && r.Signature().String() == "u" {
		return r.Value().(uint32)
	} else {
		panic(err)
	}
	return
}

func (obj *AccessPoint) SetMaxBitrate(MaxBitrate uint32) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.AccessPoint", "MaxBitrate", MaxBitrate)
}
func (obj AccessPoint) GetMaxBitrate() (ret uint32) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.AccessPoint", "MaxBitrate").Store(&r)
	if err == nil && r.Signature().String() == "u" {
		return r.Value().(uint32)
	} else {
		panic(err)
	}
	return
}

func (obj *AccessPoint) SetStrength(Strength uint8) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.AccessPoint", "Strength", Strength)
}
func (obj AccessPoint) GetStrength() (ret uint8) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.AccessPoint", "Strength").Store(&r)
	if err == nil && r.Signature().String() == "y" {
		return r.Value().(uint8)
	} else {
		panic(err)
	}
	return
}

func GetAccessPoint(path string) *AccessPoint {
	return &AccessPoint{getBus().Object("org.freedesktop.NetworkManager", dbus.ObjectPath(path)), make(chan *dbus.Signal)}
}
