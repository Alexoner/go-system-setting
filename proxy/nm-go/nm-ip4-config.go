/*This file is auto generate by dlib/dbus/proxyer. Don't edit it*/

package nm

import "dlib/dbus"

type IP4Config struct {
	core *dbus.Object
}

func (obj *IP4Config) SetAddresses(Addresses [][]uint32) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.IP4Config", "Addresses", Addresses)
}
func (obj IP4Config) GetAddresses() (ret [][]uint32) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.IP4Config", "Addresses").Store(&r)
	if err == nil && r.Signature().String() == "aau" {
		return r.Value().([][]uint32)
	} else {
		panic(err)
	}
	return
}

func (obj *IP4Config) SetNameservers(Nameservers []uint32) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.IP4Config", "Nameservers", Nameservers)
}
func (obj IP4Config) GetNameservers() (ret []uint32) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.IP4Config", "Nameservers").Store(&r)
	if err == nil && r.Signature().String() == "au" {
		return r.Value().([]uint32)
	} else {
		panic(err)
	}
	return
}

func (obj *IP4Config) SetWinsServers(WinsServers []uint32) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.IP4Config", "WinsServers", WinsServers)
}
func (obj IP4Config) GetWinsServers() (ret []uint32) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.IP4Config", "WinsServers").Store(&r)
	if err == nil && r.Signature().String() == "au" {
		return r.Value().([]uint32)
	} else {
		panic(err)
	}
	return
}

func (obj *IP4Config) SetDomains(Domains []string) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.IP4Config", "Domains", Domains)
}
func (obj IP4Config) GetDomains() (ret []string) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.IP4Config", "Domains").Store(&r)
	if err == nil && r.Signature().String() == "as" {
		return r.Value().([]string)
	} else {
		panic(err)
	}
	return
}

func (obj *IP4Config) SetRoutes(Routes [][]uint32) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.NetworkManager.IP4Config", "Routes", Routes)
}
func (obj IP4Config) GetRoutes() (ret [][]uint32) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.NetworkManager.IP4Config", "Routes").Store(&r)
	if err == nil && r.Signature().String() == "aau" {
		return r.Value().([][]uint32)
	} else {
		panic(err)
	}
	return
}

func GetIP4Config(path string) *IP4Config {
	return &IP4Config{getBus().Object("org.freedesktop.NetworkManager", dbus.ObjectPath(path))}
}
