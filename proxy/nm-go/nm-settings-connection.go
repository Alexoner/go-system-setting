/*This file is auto generate by dlib/dbus/proxyer. Don't edit it*/

package nm

import "dlib/dbus"

type SettingsConnection struct {
	core        *dbus.Object
	signal_chan chan *dbus.Signal
}

func (obj SettingsConnection) Update(properties map[string]map[string]dbus.Variant) {
	obj.core.Call("org.freedesktop.NetworkManager.Settings.Connection.Update", 0, properties).Store()
	return
}

func (obj SettingsConnection) Delete() {
	obj.core.Call("org.freedesktop.NetworkManager.Settings.Connection.Delete", 0).Store()
	return
}

func (obj SettingsConnection) GetSettings() (settings map[string]map[string]dbus.Variant) {
	obj.core.Call("org.freedesktop.NetworkManager.Settings.Connection.GetSettings", 0).Store(&settings)
	return
}

func (obj SettingsConnection) GetSecrets(setting_name string) (secrets map[string]map[string]dbus.Variant) {
	obj.core.Call("org.freedesktop.NetworkManager.Settings.Connection.GetSecrets", 0, setting_name).Store(&secrets)
	return
}

func (obj SettingsConnection) ConnectUpdated(callback func()) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+obj.core.Path()+"', interface='org.freedesktop.NetworkManager.Settings.Connection',sender='org.freedesktop.NetworkManager'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range obj.signal_chan {
			v = v
			/*callback(v.Body...)*/
		}
	}()

}

func (obj SettingsConnection) ConnectRemoved(callback func()) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+obj.core.Path()+"', interface='org.freedesktop.NetworkManager.Settings.Connection',sender='org.freedesktop.NetworkManager'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range obj.signal_chan {
			v = v
			/*callback(v.Body...)*/
		}
	}()

}

func GetSettingsConnection(path string) *SettingsConnection {
	return &SettingsConnection{getBus().Object("org.freedesktop.NetworkManager", dbus.ObjectPath(path)), make(chan *dbus.Signal)}
}
