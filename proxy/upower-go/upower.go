/*This file is auto generate by dlib/dbus/proxyer. Don't edit it*/

package upower

import "dlib/dbus"

type upower struct {
	core        *dbus.Object
	signal_chan chan *dbus.Signal
}

func (obj upower) EnumerateDevices() (devices []dbus.ObjectPath) {
	obj.core.Call("org.freedesktop.UPower.EnumerateDevices", 0).Store(&devices)
	return
}

func (obj upower) AboutToSleep(action string) {
	obj.core.Call("org.freedesktop.UPower.AboutToSleep", 0, action).Store()
	return
}

func (obj upower) Suspend() {
	obj.core.Call("org.freedesktop.UPower.Suspend", 0).Store()
	return
}

func (obj upower) SuspendAllowed() (allowed bool) {
	obj.core.Call("org.freedesktop.UPower.SuspendAllowed", 0).Store(&allowed)
	return
}

func (obj upower) Hibernate() {
	obj.core.Call("org.freedesktop.UPower.Hibernate", 0).Store()
	return
}

func (obj upower) HibernateAllowed() (allowed bool) {
	obj.core.Call("org.freedesktop.UPower.HibernateAllowed", 0).Store(&allowed)
	return
}

func (obj upower) ConnectDeviceAdded(callback func(device dbus.ObjectPath)) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+obj.core.Path()+"', interface='org.freedesktop.UPower',sender='org.freedesktop.UPower'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range obj.signal_chan {
			v = v
			/*callback(v.Body...)*/
		}
	}()

}

func (obj upower) ConnectDeviceRemoved(callback func(device dbus.ObjectPath)) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+obj.core.Path()+"', interface='org.freedesktop.UPower',sender='org.freedesktop.UPower'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range obj.signal_chan {
			v = v
			/*callback(v.Body...)*/
		}
	}()

}

func (obj upower) ConnectDeviceChanged(callback func(device dbus.ObjectPath)) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+obj.core.Path()+"', interface='org.freedesktop.UPower',sender='org.freedesktop.UPower'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range obj.signal_chan {
			v = v
			/*callback(v.Body...)*/
		}
	}()

}

func (obj upower) ConnectChanged(callback func()) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+obj.core.Path()+"', interface='org.freedesktop.UPower',sender='org.freedesktop.UPower'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range obj.signal_chan {
			v = v
			/*callback(v.Body...)*/
		}
	}()

}

func (obj upower) ConnectSleeping(callback func()) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+obj.core.Path()+"', interface='org.freedesktop.UPower',sender='org.freedesktop.UPower'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range obj.signal_chan {
			v = v
			/*callback(v.Body...)*/
		}
	}()

}

func (obj upower) ConnectNotifySleep(callback func()) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+obj.core.Path()+"', interface='org.freedesktop.UPower',sender='org.freedesktop.UPower'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range obj.signal_chan {
			v = v
			/*callback(v.Body...)*/
		}
	}()

}

func (obj upower) ConnectResuming(callback func()) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+obj.core.Path()+"', interface='org.freedesktop.UPower',sender='org.freedesktop.UPower'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range obj.signal_chan {
			v = v
			/*callback(v.Body...)*/
		}
	}()

}

func (obj upower) ConnectNotifyResume(callback func()) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+obj.core.Path()+"', interface='org.freedesktop.UPower',sender='org.freedesktop.UPower'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range obj.signal_chan {
			v = v
			/*callback(v.Body...)*/
		}
	}()

}

func (obj *upower) SetDaemonVersion(DaemonVersion string) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower", "DaemonVersion", DaemonVersion)
}
func (obj upower) GetDaemonVersion() (ret string) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower", "DaemonVersion").Store(&r)
	if err == nil && r.Signature().String() == "s" {
		return r.Value().(string)
	} else {
		panic(err)
	}
	return
}

func (obj *upower) SetCanSuspend(CanSuspend bool) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower", "CanSuspend", CanSuspend)
}
func (obj upower) GetCanSuspend() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower", "CanSuspend").Store(&r)
	if err == nil && r.Signature().String() == "b" {
		return r.Value().(bool)
	} else {
		panic(err)
	}
	return
}

func (obj *upower) SetCanHibernate(CanHibernate bool) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower", "CanHibernate", CanHibernate)
}
func (obj upower) GetCanHibernate() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower", "CanHibernate").Store(&r)
	if err == nil && r.Signature().String() == "b" {
		return r.Value().(bool)
	} else {
		panic(err)
	}
	return
}

func (obj *upower) SetOnBattery(OnBattery bool) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower", "OnBattery", OnBattery)
}
func (obj upower) GetOnBattery() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower", "OnBattery").Store(&r)
	if err == nil && r.Signature().String() == "b" {
		return r.Value().(bool)
	} else {
		panic(err)
	}
	return
}

func (obj *upower) SetOnLowBattery(OnLowBattery bool) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower", "OnLowBattery", OnLowBattery)
}
func (obj upower) GetOnLowBattery() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower", "OnLowBattery").Store(&r)
	if err == nil && r.Signature().String() == "b" {
		return r.Value().(bool)
	} else {
		panic(err)
	}
	return
}

func (obj *upower) SetLidIsClosed(LidIsClosed bool) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower", "LidIsClosed", LidIsClosed)
}
func (obj upower) GetLidIsClosed() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower", "LidIsClosed").Store(&r)
	if err == nil && r.Signature().String() == "b" {
		return r.Value().(bool)
	} else {
		panic(err)
	}
	return
}

func (obj *upower) SetLidIsPresent(LidIsPresent bool) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower", "LidIsPresent", LidIsPresent)
}
func (obj upower) GetLidIsPresent() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower", "LidIsPresent").Store(&r)
	if err == nil && r.Signature().String() == "b" {
		return r.Value().(bool)
	} else {
		panic(err)
	}
	return
}

func (obj *upower) SetLidForceSleep(LidForceSleep bool) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower", "LidForceSleep", LidForceSleep)
}
func (obj upower) GetLidForceSleep() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower", "LidForceSleep").Store(&r)
	if err == nil && r.Signature().String() == "b" {
		return r.Value().(bool)
	} else {
		panic(err)
	}
	return
}

func (obj *upower) SetIsDocked(IsDocked bool) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower", "IsDocked", IsDocked)
}
func (obj upower) GetIsDocked() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower", "IsDocked").Store(&r)
	if err == nil && r.Signature().String() == "b" {
		return r.Value().(bool)
	} else {
		panic(err)
	}
	return
}

func Getupower(path string) *upower {
	return &upower{getBus().Object("org.freedesktop.UPower", dbus.ObjectPath(path)), make(chan *dbus.Signal)}
}
