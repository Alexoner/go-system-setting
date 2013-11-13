/*This file is auto generate by dlib/dbus/proxyer. Don't edit it*/

package nm

import "dlib/dbus"

type Manager struct {
	core        *dbus.Object
	signal_chan chan *dbus.Signal
}

func (obj Manager) GetDevices() (devices []dbus.ObjectPath) {
	obj.core.Call("org.freedesktop.NetworkManager.GetDevices", 0).Store(&devices)
	return
}

func (obj Manager) GetDeviceByIpIface(iface string) (device dbus.ObjectPath) {
	obj.core.Call("org.freedesktop.NetworkManager.GetDeviceByIpIface", 0, iface).Store(&device)
	return
}

func (obj Manager) ActivateConnection(connection dbus.ObjectPath, device dbus.ObjectPath, specific_object dbus.ObjectPath) (active_connection dbus.ObjectPath) {
	obj.core.Call("org.freedesktop.NetworkManager.ActivateConnection", 0, connection, device, specific_object).Store(&active_connection)
	return
}

func (obj Manager) AddAndActivateConnection(connection map[string]map[string]dbus.Variant, device dbus.ObjectPath, specific_object dbus.ObjectPath) (path dbus.ObjectPath, active_connection dbus.ObjectPath) {
	obj.core.Call("org.freedesktop.NetworkManager.AddAndActivateConnection", 0, connection, device, specific_object).Store(&path, &active_connection)
	return
}

func (obj Manager) DeactivateConnection(active_connection dbus.ObjectPath) {
	obj.core.Call("org.freedesktop.NetworkManager.DeactivateConnection", 0, active_connection).Store()
	return
}

func (obj Manager) Sleep(sleep bool) {
	obj.core.Call("org.freedesktop.NetworkManager.Sleep", 0, sleep).Store()
	return
}

func (obj Manager) Enable(enable bool) {
	obj.core.Call("org.freedesktop.NetworkManager.Enable", 0, enable).Store()
	return
}

func (obj Manager) GetPermissions() (permissions map[string]string) {
	obj.core.Call("org.freedesktop.NetworkManager.GetPermissions", 0).Store(&permissions)
	return
}

func (obj Manager) SetLogging(level string, domains string) {
	obj.core.Call("org.freedesktop.NetworkManager.SetLogging", 0, level, domains).Store()
	return
}

func (obj Manager) GetLogging() (level string, domains string) {
	obj.core.Call("org.freedesktop.NetworkManager.GetLogging", 0).Store(&level, &domains)
	return
}

func (obj Manager) state() (state uint32) {
	obj.core.Call("org.freedesktop.NetworkManager.state", 0).Store(&state)
	return
}

func (obj Manager) ConnectCheckPermissions(callback func()) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+obj.core.Path()+"', interface='org.freedesktop.NetworkManager',sender='org.freedesktop.NetworkManager'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range obj.signal_chan {
			v = v
			/*callback(v.Body...)*/
		}
	}()

}

func (obj Manager) ConnectStateChanged(callback func(state uint32)) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+obj.core.Path()+"', interface='org.freedesktop.NetworkManager',sender='org.freedesktop.NetworkManager'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range obj.signal_chan {
			v = v
			/*callback(v.Body...)*/
		}
	}()

}

func (obj Manager) ConnectPropertiesChanged(callback func(properties map[string]dbus.Variant)) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+obj.core.Path()+"', interface='org.freedesktop.NetworkManager',sender='org.freedesktop.NetworkManager'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range obj.signal_chan {
			v = v
			/*callback(v.Body...)*/
		}
	}()

}

func (obj Manager) ConnectDeviceAdded(callback func(device_path dbus.ObjectPath)) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+obj.core.Path()+"', interface='org.freedesktop.NetworkManager',sender='org.freedesktop.NetworkManager'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range obj.signal_chan {
			v = v
			/*callback(v.Body...)*/
		}
	}()

}

func (obj Manager) ConnectDeviceRemoved(callback func(device_path dbus.ObjectPath)) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+obj.core.Path()+"', interface='org.freedesktop.NetworkManager',sender='org.freedesktop.NetworkManager'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range obj.signal_chan {
			v = v
			/*callback(v.Body...)*/
		}
	}()

}

func (obj *Manager) SetNetworkingEnabled(NetworkingEnabled bool) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager", "NetworkingEnabled", NetworkingEnabled)
}
func (obj Manager) GetNetworkingEnabled() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager", "NetworkingEnabled").Store(&r)
	if err == nil && r.Signature().String() == "b" {
		return r.Value().(bool)
	} else {
		panic(err)
	}
	return
}

func (obj *Manager) SetWirelessEnabled(WirelessEnabled bool) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager", "WirelessEnabled", WirelessEnabled)
}
func (obj Manager) GetWirelessEnabled() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager", "WirelessEnabled").Store(&r)
	if err == nil && r.Signature().String() == "b" {
		return r.Value().(bool)
	} else {
		panic(err)
	}
	return
}

func (obj *Manager) SetWirelessHardwareEnabled(WirelessHardwareEnabled bool) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager", "WirelessHardwareEnabled", WirelessHardwareEnabled)
}
func (obj Manager) GetWirelessHardwareEnabled() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager", "WirelessHardwareEnabled").Store(&r)
	if err == nil && r.Signature().String() == "b" {
		return r.Value().(bool)
	} else {
		panic(err)
	}
	return
}

func (obj *Manager) SetWwanEnabled(WwanEnabled bool) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager", "WwanEnabled", WwanEnabled)
}
func (obj Manager) GetWwanEnabled() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager", "WwanEnabled").Store(&r)
	if err == nil && r.Signature().String() == "b" {
		return r.Value().(bool)
	} else {
		panic(err)
	}
	return
}

func (obj *Manager) SetWwanHardwareEnabled(WwanHardwareEnabled bool) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager", "WwanHardwareEnabled", WwanHardwareEnabled)
}
func (obj Manager) GetWwanHardwareEnabled() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager", "WwanHardwareEnabled").Store(&r)
	if err == nil && r.Signature().String() == "b" {
		return r.Value().(bool)
	} else {
		panic(err)
	}
	return
}

func (obj *Manager) SetWimaxEnabled(WimaxEnabled bool) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager", "WimaxEnabled", WimaxEnabled)
}
func (obj Manager) GetWimaxEnabled() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager", "WimaxEnabled").Store(&r)
	if err == nil && r.Signature().String() == "b" {
		return r.Value().(bool)
	} else {
		panic(err)
	}
	return
}

func (obj *Manager) SetWimaxHardwareEnabled(WimaxHardwareEnabled bool) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager", "WimaxHardwareEnabled", WimaxHardwareEnabled)
}
func (obj Manager) GetWimaxHardwareEnabled() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager", "WimaxHardwareEnabled").Store(&r)
	if err == nil && r.Signature().String() == "b" {
		return r.Value().(bool)
	} else {
		panic(err)
	}
	return
}

func (obj *Manager) SetActiveConnections(ActiveConnections []dbus.ObjectPath) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager", "ActiveConnections", ActiveConnections)
}
func (obj Manager) GetActiveConnections() (ret []dbus.ObjectPath) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager", "ActiveConnections").Store(&r)
	if err == nil && r.Signature().String() == "ao" {
		return r.Value().([]dbus.ObjectPath)
	} else {
		panic(err)
	}
	return
}

func (obj *Manager) SetVersion(Version string) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager", "Version", Version)
}
func (obj Manager) GetVersion() (ret string) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager", "Version").Store(&r)
	if err == nil && r.Signature().String() == "s" {
		return r.Value().(string)
	} else {
		panic(err)
	}
	return
}

func (obj *Manager) SetState(State uint32) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager", "State", State)
}
func (obj Manager) GetState() (ret uint32) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager", "State").Store(&r)
	if err == nil && r.Signature().String() == "u" {
		return r.Value().(uint32)
	} else {
		panic(err)
	}
	return
}

func GetManager(path string) *Manager {
	return &Manager{getBus().Object("org.freedesktop.NetworkManager", dbus.ObjectPath(path)), make(chan *dbus.Signal)}
}
