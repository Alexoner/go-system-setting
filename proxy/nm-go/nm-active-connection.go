/*This file is auto generate by dlib/dbus/proxyer. Don't edit it*/

package nm

import "dlib/dbus"

type ActiveConnection struct {
	core        *dbus.Object
	signal_chan chan *dbus.Signal
}

func (obj ActiveConnection) ConnectPropertiesChanged(callback func(properties map[string]dbus.Variant)) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+obj.core.Path()+"', interface='org.freedesktop.NetworkManager.Connection.Active',sender='org.freedesktop.NetworkManager'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range obj.signal_chan {
			v = v
			/*callback(v.Body...)*/
		}
	}()

}

func (obj *ActiveConnection) SetConnection(Connection dbus.ObjectPath) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.Connection.Active", "Connection", Connection)
}
func (obj ActiveConnection) GetConnection() (ret dbus.ObjectPath) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.Connection.Active", "Connection").Store(&r)
	if err == nil && r.Signature().String() == "o" {
		return r.Value().(dbus.ObjectPath)
	} else {
		panic(err)
	}
	return
}

func (obj *ActiveConnection) SetSpecificObject(SpecificObject dbus.ObjectPath) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.Connection.Active", "SpecificObject", SpecificObject)
}
func (obj ActiveConnection) GetSpecificObject() (ret dbus.ObjectPath) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.Connection.Active", "SpecificObject").Store(&r)
	if err == nil && r.Signature().String() == "o" {
		return r.Value().(dbus.ObjectPath)
	} else {
		panic(err)
	}
	return
}

func (obj *ActiveConnection) SetUuid(Uuid string) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.Connection.Active", "Uuid", Uuid)
}
func (obj ActiveConnection) GetUuid() (ret string) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.Connection.Active", "Uuid").Store(&r)
	if err == nil && r.Signature().String() == "s" {
		return r.Value().(string)
	} else {
		panic(err)
	}
	return
}

func (obj *ActiveConnection) SetDevices(Devices []dbus.ObjectPath) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.Connection.Active", "Devices", Devices)
}
func (obj ActiveConnection) GetDevices() (ret []dbus.ObjectPath) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.Connection.Active", "Devices").Store(&r)
	if err == nil && r.Signature().String() == "ao" {
		return r.Value().([]dbus.ObjectPath)
	} else {
		panic(err)
	}
	return
}

func (obj *ActiveConnection) SetState(State uint32) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.Connection.Active", "State", State)
}
func (obj ActiveConnection) GetState() (ret uint32) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.Connection.Active", "State").Store(&r)
	if err == nil && r.Signature().String() == "u" {
		return r.Value().(uint32)
	} else {
		panic(err)
	}
	return
}

func (obj *ActiveConnection) SetDefault(Default bool) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.Connection.Active", "Default", Default)
}
func (obj ActiveConnection) GetDefault() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.Connection.Active", "Default").Store(&r)
	if err == nil && r.Signature().String() == "b" {
		return r.Value().(bool)
	} else {
		panic(err)
	}
	return
}

func (obj *ActiveConnection) SetDefault6(Default6 bool) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.Connection.Active", "Default6", Default6)
}
func (obj ActiveConnection) GetDefault6() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.Connection.Active", "Default6").Store(&r)
	if err == nil && r.Signature().String() == "b" {
		return r.Value().(bool)
	} else {
		panic(err)
	}
	return
}

func (obj *ActiveConnection) SetVpn(Vpn bool) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.Connection.Active", "Vpn", Vpn)
}
func (obj ActiveConnection) GetVpn() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.Connection.Active", "Vpn").Store(&r)
	if err == nil && r.Signature().String() == "b" {
		return r.Value().(bool)
	} else {
		panic(err)
	}
	return
}

func (obj *ActiveConnection) SetMaster(Master dbus.ObjectPath) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.Connection.Active", "Master", Master)
}
func (obj ActiveConnection) GetMaster() (ret dbus.ObjectPath) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.Connection.Active", "Master").Store(&r)
	if err == nil && r.Signature().String() == "o" {
		return r.Value().(dbus.ObjectPath)
	} else {
		panic(err)
	}
	return
}

func GetActiveConnection(path string) *ActiveConnection {
	return &ActiveConnection{getBus().Object("org.freedesktop.NetworkManager", dbus.ObjectPath(path)), make(chan *dbus.Signal)}
}
