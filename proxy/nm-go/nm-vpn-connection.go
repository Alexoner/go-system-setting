/*This file is auto generate by dlib/dbus/proxyer. Don't edit it*/

package nm

import "dlib/dbus"

type VPNConnection struct {
	core        *dbus.Object
	signal_chan chan *dbus.Signal
}

func (obj VPNConnection) ConnectPropertiesChanged(callback func(properties map[string]dbus.Variant)) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+obj.core.Path()+"', interface='org.freedesktop.NetworkManager.VPN.Connection',sender='org.freedesktop.NetworkManager'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range obj.signal_chan {
			v = v
			/*callback(v.Body...)*/
		}
	}()

}

func (obj VPNConnection) ConnectVpnStateChanged(callback func(state uint32, reason uint32)) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+obj.core.Path()+"', interface='org.freedesktop.NetworkManager.VPN.Connection',sender='org.freedesktop.NetworkManager'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range obj.signal_chan {
			v = v
			/*callback(v.Body...)*/
		}
	}()

}

func (obj *VPNConnection) SetVpnState(VpnState uint32) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.VPN.Connection", "VpnState", VpnState)
}
func (obj VPNConnection) GetVpnState() (ret uint32) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.VPN.Connection", "VpnState").Store(&r)
	if err == nil && r.Signature().String() == "u" {
		return r.Value().(uint32)
	} else {
		panic(err)
	}
	return
}

func (obj *VPNConnection) SetBanner(Banner string) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.VPN.Connection", "Banner", Banner)
}
func (obj VPNConnection) GetBanner() (ret string) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.VPN.Connection", "Banner").Store(&r)
	if err == nil && r.Signature().String() == "s" {
		return r.Value().(string)
	} else {
		panic(err)
	}
	return
}

func GetVPNConnection(path string) *VPNConnection {
	return &VPNConnection{getBus().Object("org.freedesktop.NetworkManager", dbus.ObjectPath(path)), make(chan *dbus.Signal)}
}
