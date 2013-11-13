/*This file is auto generate by dlib/dbus/proxyer. Don't edit it*/

package nm

import "dlib/dbus"

type Settings struct {
	core        *dbus.Object
	signal_chan chan *dbus.Signal
}

func (obj Settings) ListConnections() (connections []dbus.ObjectPath) {
	obj.core.Call("org.freedesktop.NetworkManager.Settings.ListConnections", 0).Store(&connections)
	return
}

func (obj Settings) GetConnectionByUuid(uuid string) (connection dbus.ObjectPath) {
	obj.core.Call("org.freedesktop.NetworkManager.Settings.GetConnectionByUuid", 0, uuid).Store(&connection)
	return
}

func (obj Settings) AddConnection(connection map[string]map[string]dbus.Variant) (path dbus.ObjectPath) {
	obj.core.Call("org.freedesktop.NetworkManager.Settings.AddConnection", 0, connection).Store(&path)
	return
}

func (obj Settings) SaveHostname(hostname string) {
	obj.core.Call("org.freedesktop.NetworkManager.Settings.SaveHostname", 0, hostname).Store()
	return
}

func (obj Settings) ConnectPropertiesChanged(callback func(properties map[string]dbus.Variant)) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+obj.core.Path()+"', interface='org.freedesktop.NetworkManager.Settings',sender='org.freedesktop.NetworkManager'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range obj.signal_chan {
			v = v
			/*callback(v.Body...)*/
		}
	}()

}

func (obj Settings) ConnectNewConnection(callback func(connection dbus.ObjectPath)) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+obj.core.Path()+"', interface='org.freedesktop.NetworkManager.Settings',sender='org.freedesktop.NetworkManager'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range obj.signal_chan {
			v = v
			/*callback(v.Body...)*/
		}
	}()

}

func (obj *Settings) SetHostname(Hostname string) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.Settings", "Hostname", Hostname)
}
func (obj Settings) GetHostname() (ret string) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.Settings", "Hostname").Store(&r)
	if err == nil && r.Signature().String() == "s" {
		return r.Value().(string)
	} else {
		panic(err)
	}
	return
}

func (obj *Settings) SetCanModify(CanModify bool) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.Settings", "CanModify", CanModify)
}
func (obj Settings) GetCanModify() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.Settings", "CanModify").Store(&r)
	if err == nil && r.Signature().String() == "b" {
		return r.Value().(bool)
	} else {
		panic(err)
	}
	return
}

func GetSettings(path string) *Settings {
	return &Settings{getBus().Object("org.freedesktop.NetworkManager", dbus.ObjectPath(path)), make(chan *dbus.Signal)}
}
