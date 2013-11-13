/*This file is auto generate by dlib/dbus/proxyer. Don't edit it*/

package nm

import "dlib/dbus"

type DHCP4Config struct {
	core        *dbus.Object
	signal_chan chan *dbus.Signal
}

func (obj DHCP4Config) ConnectPropertiesChanged(callback func(properties map[string]dbus.Variant)) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+obj.core.Path()+"', interface='org.freedesktop.NetworkManager.DHCP4Config',sender='org.freedesktop.NetworkManager'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range obj.signal_chan {
			v = v
			/*callback(v.Body...)*/
		}
	}()

}

func (obj *DHCP4Config) SetOptions(Options map[string]dbus.Variant) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.DHCP4Config", "Options", Options)
}
func (obj DHCP4Config) GetOptions() (ret map[string]dbus.Variant) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.DHCP4Config", "Options").Store(&r)
	if err == nil && r.Signature().String() == "a{sv}" {
		return r.Value().(map[string]dbus.Variant)
	} else {
		panic(err)
	}
	return
}

func GetDHCP4Config(path string) *DHCP4Config {
	return &DHCP4Config{getBus().Object("org.freedesktop.NetworkManager", dbus.ObjectPath(path)), make(chan *dbus.Signal)}
}
