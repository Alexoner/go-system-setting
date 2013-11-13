/*This file is auto generate by dlib/dbus/proxyer. Don't edit it*/

package nm

import "dlib/dbus"

type Device struct {
	core        *dbus.Object
	signal_chan chan *dbus.Signal
}

func (obj Device) Disconnect() {
	obj.core.Call("org.freedesktop.NetworkManager.Device.Disconnect", 0).Store()
	return
}

func (obj Device) ConnectStateChanged(callback func(new_state uint32, old_state uint32, reason uint32)) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+obj.core.Path()+"', interface='org.freedesktop.NetworkManager.Device',sender='org.freedesktop.NetworkManager'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range obj.signal_chan {
			v = v
			/*callback(v.Body...)*/
		}
	}()

}

func (obj *Device) SetUdi(Udi string) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.Device", "Udi", Udi)
}
func (obj Device) GetUdi() (ret string) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.Device", "Udi").Store(&r)
	if err == nil && r.Signature().String() == "s" {
		return r.Value().(string)
	} else {
		panic(err)
	}
	return
}

func (obj *Device) SetInterface(Interface string) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.Device", "Interface", Interface)
}
func (obj Device) GetInterface() (ret string) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.Device", "Interface").Store(&r)
	if err == nil && r.Signature().String() == "s" {
		return r.Value().(string)
	} else {
		panic(err)
	}
	return
}

func (obj *Device) SetIpInterface(IpInterface string) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.Device", "IpInterface", IpInterface)
}
func (obj Device) GetIpInterface() (ret string) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.Device", "IpInterface").Store(&r)
	if err == nil && r.Signature().String() == "s" {
		return r.Value().(string)
	} else {
		panic(err)
	}
	return
}

func (obj *Device) SetDriver(Driver string) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.Device", "Driver", Driver)
}
func (obj Device) GetDriver() (ret string) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.Device", "Driver").Store(&r)
	if err == nil && r.Signature().String() == "s" {
		return r.Value().(string)
	} else {
		panic(err)
	}
	return
}

func (obj *Device) SetDriverVersion(DriverVersion string) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.Device", "DriverVersion", DriverVersion)
}
func (obj Device) GetDriverVersion() (ret string) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.Device", "DriverVersion").Store(&r)
	if err == nil && r.Signature().String() == "s" {
		return r.Value().(string)
	} else {
		panic(err)
	}
	return
}

func (obj *Device) SetFirmwareVersion(FirmwareVersion string) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.Device", "FirmwareVersion", FirmwareVersion)
}
func (obj Device) GetFirmwareVersion() (ret string) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.Device", "FirmwareVersion").Store(&r)
	if err == nil && r.Signature().String() == "s" {
		return r.Value().(string)
	} else {
		panic(err)
	}
	return
}

func (obj *Device) SetCapabilities(Capabilities uint32) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.Device", "Capabilities", Capabilities)
}
func (obj Device) GetCapabilities() (ret uint32) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.Device", "Capabilities").Store(&r)
	if err == nil && r.Signature().String() == "u" {
		return r.Value().(uint32)
	} else {
		panic(err)
	}
	return
}

func (obj *Device) SetIp4Address(Ip4Address int32) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.Device", "Ip4Address", Ip4Address)
}
func (obj Device) GetIp4Address() (ret int32) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.Device", "Ip4Address").Store(&r)
	if err == nil && r.Signature().String() == "i" {
		return r.Value().(int32)
	} else {
		panic(err)
	}
	return
}

func (obj *Device) SetState(State uint32) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.Device", "State", State)
}
func (obj Device) GetState() (ret uint32) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.Device", "State").Store(&r)
	if err == nil && r.Signature().String() == "u" {
		return r.Value().(uint32)
	} else {
		panic(err)
	}
	return
}

func (obj *Device) SetStateReason(StateReason []interface{}) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.Device", "StateReason", StateReason)
}
func (obj Device) GetStateReason() (ret []interface{}) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.Device", "StateReason").Store(&r)
	if err == nil && r.Signature().String() == "(uu)" {
		return r.Value().([]interface{})
	} else {
		panic(err)
	}
	return
}

func (obj *Device) SetActiveConnection(ActiveConnection dbus.ObjectPath) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.Device", "ActiveConnection", ActiveConnection)
}
func (obj Device) GetActiveConnection() (ret dbus.ObjectPath) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.Device", "ActiveConnection").Store(&r)
	if err == nil && r.Signature().String() == "o" {
		return r.Value().(dbus.ObjectPath)
	} else {
		panic(err)
	}
	return
}

func (obj *Device) SetIp4Config(Ip4Config dbus.ObjectPath) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.Device", "Ip4Config", Ip4Config)
}
func (obj Device) GetIp4Config() (ret dbus.ObjectPath) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.Device", "Ip4Config").Store(&r)
	if err == nil && r.Signature().String() == "o" {
		return r.Value().(dbus.ObjectPath)
	} else {
		panic(err)
	}
	return
}

func (obj *Device) SetDhcp4Config(Dhcp4Config dbus.ObjectPath) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.Device", "Dhcp4Config", Dhcp4Config)
}
func (obj Device) GetDhcp4Config() (ret dbus.ObjectPath) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.Device", "Dhcp4Config").Store(&r)
	if err == nil && r.Signature().String() == "o" {
		return r.Value().(dbus.ObjectPath)
	} else {
		panic(err)
	}
	return
}

func (obj *Device) SetIp6Config(Ip6Config dbus.ObjectPath) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.Device", "Ip6Config", Ip6Config)
}
func (obj Device) GetIp6Config() (ret dbus.ObjectPath) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.Device", "Ip6Config").Store(&r)
	if err == nil && r.Signature().String() == "o" {
		return r.Value().(dbus.ObjectPath)
	} else {
		panic(err)
	}
	return
}

func (obj *Device) SetDhcp6Config(Dhcp6Config dbus.ObjectPath) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.Device", "Dhcp6Config", Dhcp6Config)
}
func (obj Device) GetDhcp6Config() (ret dbus.ObjectPath) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.Device", "Dhcp6Config").Store(&r)
	if err == nil && r.Signature().String() == "o" {
		return r.Value().(dbus.ObjectPath)
	} else {
		panic(err)
	}
	return
}

func (obj *Device) SetManaged(Managed bool) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.Device", "Managed", Managed)
}
func (obj Device) GetManaged() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.Device", "Managed").Store(&r)
	if err == nil && r.Signature().String() == "b" {
		return r.Value().(bool)
	} else {
		panic(err)
	}
	return
}

func (obj *Device) SetAutoconnect(Autoconnect bool) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.Device", "Autoconnect", Autoconnect)
}
func (obj Device) GetAutoconnect() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.Device", "Autoconnect").Store(&r)
	if err == nil && r.Signature().String() == "b" {
		return r.Value().(bool)
	} else {
		panic(err)
	}
	return
}

func (obj *Device) SetFirmwareMissing(FirmwareMissing bool) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.Device", "FirmwareMissing", FirmwareMissing)
}
func (obj Device) GetFirmwareMissing() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.Device", "FirmwareMissing").Store(&r)
	if err == nil && r.Signature().String() == "b" {
		return r.Value().(bool)
	} else {
		panic(err)
	}
	return
}

func (obj *Device) SetDeviceType(DeviceType uint32) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.Device", "DeviceType", DeviceType)
}
func (obj Device) GetDeviceType() (ret uint32) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.Device", "DeviceType").Store(&r)
	if err == nil && r.Signature().String() == "u" {
		return r.Value().(uint32)
	} else {
		panic(err)
	}
	return
}

func (obj *Device) SetAvailableConnections(AvailableConnections []dbus.ObjectPath) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.Device", "AvailableConnections", AvailableConnections)
}
func (obj Device) GetAvailableConnections() (ret []dbus.ObjectPath) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.Device", "AvailableConnections").Store(&r)
	if err == nil && r.Signature().String() == "ao" {
		return r.Value().([]dbus.ObjectPath)
	} else {
		panic(err)
	}
	return
}

func GetDevice(path string) *Device {
	return &Device{getBus().Object("org.freedesktop.NetworkManager", dbus.ObjectPath(path)), make(chan *dbus.Signal)}
}
