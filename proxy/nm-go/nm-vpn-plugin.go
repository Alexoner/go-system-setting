/*This file is auto generate by dlib/dbus/proxyer. Don't edit it*/

package nm

import "dlib/dbus"

type VPNPlugin struct {
	core        *dbus.Object
	signal_chan chan *dbus.Signal
}

func (obj VPNPlugin) Connect(connection map[string]map[string]dbus.Variant) {
	obj.core.Call("org.freedesktop.NetworkManager.VPN.Plugin.Connect", 0, connection).Store()
	return
}

func (obj VPNPlugin) NeedSecrets(settings map[string]map[string]dbus.Variant) (setting_name string) {
	obj.core.Call("org.freedesktop.NetworkManager.VPN.Plugin.NeedSecrets", 0, settings).Store(&setting_name)
	return
}

func (obj VPNPlugin) Disconnect() {
	obj.core.Call("org.freedesktop.NetworkManager.VPN.Plugin.Disconnect", 0).Store()
	return
}

func (obj VPNPlugin) SetConfig(config map[string]dbus.Variant) {
	obj.core.Call("org.freedesktop.NetworkManager.VPN.Plugin.SetConfig", 0, config).Store()
	return
}

func (obj VPNPlugin) SetIp4Config(config map[string]dbus.Variant) {
	obj.core.Call("org.freedesktop.NetworkManager.VPN.Plugin.SetIp4Config", 0, config).Store()
	return
}

func (obj VPNPlugin) SetIp6Config(config map[string]dbus.Variant) {
	obj.core.Call("org.freedesktop.NetworkManager.VPN.Plugin.SetIp6Config", 0, config).Store()
	return
}

func (obj VPNPlugin) SetFailure(reason string) {
	obj.core.Call("org.freedesktop.NetworkManager.VPN.Plugin.SetFailure", 0, reason).Store()
	return
}

func (obj VPNPlugin) ConnectStateChanged(callback func(state uint32)) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+obj.core.Path()+"', interface='org.freedesktop.NetworkManager.VPN.Plugin',sender='org.freedesktop.NetworkManager'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range obj.signal_chan {
			v = v
			/*callback(v.Body...)*/
		}
	}()

}

func (obj VPNPlugin) ConnectConfig(callback func(config map[string]dbus.Variant)) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+obj.core.Path()+"', interface='org.freedesktop.NetworkManager.VPN.Plugin',sender='org.freedesktop.NetworkManager'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range obj.signal_chan {
			v = v
			/*callback(v.Body...)*/
		}
	}()

}

func (obj VPNPlugin) ConnectIp4Config(callback func(ip4config map[string]dbus.Variant)) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+obj.core.Path()+"', interface='org.freedesktop.NetworkManager.VPN.Plugin',sender='org.freedesktop.NetworkManager'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range obj.signal_chan {
			v = v
			/*callback(v.Body...)*/
		}
	}()

}

func (obj VPNPlugin) ConnectIp6Config(callback func(ip6config map[string]dbus.Variant)) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+obj.core.Path()+"', interface='org.freedesktop.NetworkManager.VPN.Plugin',sender='org.freedesktop.NetworkManager'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range obj.signal_chan {
			v = v
			/*callback(v.Body...)*/
		}
	}()

}

func (obj VPNPlugin) ConnectLoginBanner(callback func(banner string)) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+obj.core.Path()+"', interface='org.freedesktop.NetworkManager.VPN.Plugin',sender='org.freedesktop.NetworkManager'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range obj.signal_chan {
			v = v
			/*callback(v.Body...)*/
		}
	}()

}

func (obj VPNPlugin) ConnectFailure(callback func(reason uint32)) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+obj.core.Path()+"', interface='org.freedesktop.NetworkManager.VPN.Plugin',sender='org.freedesktop.NetworkManager'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range obj.signal_chan {
			v = v
			/*callback(v.Body...)*/
		}
	}()

}

func (obj *VPNPlugin) SetState(State uint32) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.VPN.Plugin", "State", State)
}
func (obj VPNPlugin) GetState() (ret uint32) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.VPN.Plugin", "State").Store(&r)
	if err == nil && r.Signature().String() == "u" {
		return r.Value().(uint32)
	} else {
		panic(err)
	}
	return
}

func GetVPNPlugin(path string) *VPNPlugin {
	return &VPNPlugin{getBus().Object("org.freedesktop.NetworkManager", dbus.ObjectPath(path)), make(chan *dbus.Signal)}
}
