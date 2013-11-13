/*This file is auto generate by dlib/dbus/proxyer. Don't edit it*/

package nm

import "dlib/dbus"

type DeviceWireless struct {
	core        *dbus.Object
	signal_chan chan *dbus.Signal
}

func (obj DeviceWireless) GetAccessPoints() (access_points []dbus.ObjectPath) {
	obj.core.Call("org.freedesktop.NetworkManager.Device.Wireless.GetAccessPoints", 0).Store(&access_points)
	return
}

func (obj DeviceWireless) RequestScan(options map[string]dbus.Variant) {
	obj.core.Call("org.freedesktop.NetworkManager.Device.Wireless.RequestScan", 0, options).Store()
	return
}

func (obj DeviceWireless) ConnectPropertiesChanged(callback func(properties map[string]dbus.Variant)) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+obj.core.Path()+"', interface='org.freedesktop.NetworkManager.Device.Wireless',sender='org.freedesktop.NetworkManager'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range obj.signal_chan {
			v = v
			/*callback(v.Body...)*/
		}
	}()

}

func (obj DeviceWireless) ConnectAccessPointAdded(callback func(access_point dbus.ObjectPath)) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+obj.core.Path()+"', interface='org.freedesktop.NetworkManager.Device.Wireless',sender='org.freedesktop.NetworkManager'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range obj.signal_chan {
			v = v
			/*callback(v.Body...)*/
		}
	}()

}

func (obj DeviceWireless) ConnectAccessPointRemoved(callback func(access_point dbus.ObjectPath)) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+obj.core.Path()+"', interface='org.freedesktop.NetworkManager.Device.Wireless',sender='org.freedesktop.NetworkManager'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range obj.signal_chan {
			v = v
			/*callback(v.Body...)*/
		}
	}()

}

func (obj *DeviceWireless) SetHwAddress(HwAddress string) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.Device.Wireless", "HwAddress", HwAddress)
}
func (obj DeviceWireless) GetHwAddress() (ret string) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.Device.Wireless", "HwAddress").Store(&r)
	if err == nil && r.Signature().String() == "s" {
		return r.Value().(string)
	} else {
		panic(err)
	}
	return
}

func (obj *DeviceWireless) SetPermHwAddress(PermHwAddress string) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.Device.Wireless", "PermHwAddress", PermHwAddress)
}
func (obj DeviceWireless) GetPermHwAddress() (ret string) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.Device.Wireless", "PermHwAddress").Store(&r)
	if err == nil && r.Signature().String() == "s" {
		return r.Value().(string)
	} else {
		panic(err)
	}
	return
}

func (obj *DeviceWireless) SetMode(Mode uint32) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.Device.Wireless", "Mode", Mode)
}
func (obj DeviceWireless) GetMode() (ret uint32) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.Device.Wireless", "Mode").Store(&r)
	if err == nil && r.Signature().String() == "u" {
		return r.Value().(uint32)
	} else {
		panic(err)
	}
	return
}

func (obj *DeviceWireless) SetBitrate(Bitrate uint32) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.Device.Wireless", "Bitrate", Bitrate)
}
func (obj DeviceWireless) GetBitrate() (ret uint32) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.Device.Wireless", "Bitrate").Store(&r)
	if err == nil && r.Signature().String() == "u" {
		return r.Value().(uint32)
	} else {
		panic(err)
	}
	return
}

func (obj *DeviceWireless) SetActiveAccessPoint(ActiveAccessPoint dbus.ObjectPath) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.Device.Wireless", "ActiveAccessPoint", ActiveAccessPoint)
}
func (obj DeviceWireless) GetActiveAccessPoint() (ret dbus.ObjectPath) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.Device.Wireless", "ActiveAccessPoint").Store(&r)
	if err == nil && r.Signature().String() == "o" {
		return r.Value().(dbus.ObjectPath)
	} else {
		panic(err)
	}
	return
}

func (obj *DeviceWireless) SetWirelessCapabilities(WirelessCapabilities uint32) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.Device.Wireless", "WirelessCapabilities", WirelessCapabilities)
}
func (obj DeviceWireless) GetWirelessCapabilities() (ret uint32) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.Device.Wireless", "WirelessCapabilities").Store(&r)
	if err == nil && r.Signature().String() == "u" {
		return r.Value().(uint32)
	} else {
		panic(err)
	}
	return
}

func GetDeviceWireless(path string) *DeviceWireless {
	return &DeviceWireless{getBus().Object("org.freedesktop.NetworkManager", dbus.ObjectPath(path)), make(chan *dbus.Signal)}
}
