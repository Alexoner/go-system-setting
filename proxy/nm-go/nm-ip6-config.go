/*This file is auto generate by dlib/dbus/proxyer. Don't edit it*/

package nm

import "dlib/dbus"

type IP6Config struct {
	core *dbus.Object
}

func (obj *IP6Config) SetAddresses(Addresses [][]interface{}) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.IP6Config", "Addresses", Addresses)
}
func (obj IP6Config) GetAddresses() (ret [][]interface{}) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.IP6Config", "Addresses").Store(&r)
	if err == nil && r.Signature().String() == "a(ayuay)" {
		return r.Value().([][]interface{})
	} else {
		panic(err)
	}
	return
}

func (obj *IP6Config) SetNameservers(Nameservers [][]uint8) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.IP6Config", "Nameservers", Nameservers)
}
func (obj IP6Config) GetNameservers() (ret [][]uint8) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.IP6Config", "Nameservers").Store(&r)
	if err == nil && r.Signature().String() == "aay" {
		return r.Value().([][]uint8)
	} else {
		panic(err)
	}
	return
}

func (obj *IP6Config) SetDomains(Domains []string) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.IP6Config", "Domains", Domains)
}
func (obj IP6Config) GetDomains() (ret []string) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.IP6Config", "Domains").Store(&r)
	if err == nil && r.Signature().String() == "as" {
		return r.Value().([]string)
	} else {
		panic(err)
	}
	return
}

func (obj *IP6Config) SetRoutes(Routes [][]interface{}) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.IP6Config", "Routes", Routes)
}
func (obj IP6Config) GetRoutes() (ret [][]interface{}) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.IP6Config", "Routes").Store(&r)
	if err == nil && r.Signature().String() == "a(ayuayu)" {
		return r.Value().([][]interface{})
	} else {
		panic(err)
	}
	return
}

func GetIP6Config(path string) *IP6Config {
	return &IP6Config{getBus().Object("org.freedesktop.NetworkManager", dbus.ObjectPath(path))}
}
