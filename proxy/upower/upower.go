/*This file is auto generate by dlib/dbus/proxyer. Don't edit it*/

package upower

import "dlib/dbus"
import "reflect"

type Upower struct {
	Path        dbus.ObjectPath
	core        *dbus.Object
	signal_chan chan *dbus.Signal
}

func (obj Upower) EnumerateDevices() (devices []dbus.ObjectPath) {
	obj.core.Call("org.freedesktop.UPower.EnumerateDevices", 0).Store(&devices)
	return
}

func (obj Upower) AboutToSleep(action string) {
	obj.core.Call("org.freedesktop.UPower.AboutToSleep", 0, action).Store()
	return
}

func (obj Upower) Suspend() {
	obj.core.Call("org.freedesktop.UPower.Suspend", 0).Store()
	return
}

func (obj Upower) SuspendAllowed() (allowed bool) {
	obj.core.Call("org.freedesktop.UPower.SuspendAllowed", 0).Store(&allowed)
	return
}

func (obj Upower) Hibernate() {
	obj.core.Call("org.freedesktop.UPower.Hibernate", 0).Store()
	return
}

func (obj Upower) HibernateAllowed() (allowed bool) {
	obj.core.Call("org.freedesktop.UPower.HibernateAllowed", 0).Store(&allowed)
	return
}

func (obj Upower) ConnectDeviceAdded(callback func(device dbus.ObjectPath)) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+string(obj.core.Path())+"', interface='org.freedesktop.UPower',sender='org.freedesktop.UPower',member='DeviceAdded'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range obj.signal_chan {
			if v.Name != "org.freedesktop.UPower.DeviceAdded" || 1 != len(v.Body) {
				continue
			}

			if reflect.TypeOf(v.Body[0]) != reflect.TypeOf((*dbus.ObjectPath)(nil)).Elem() {
				continue
			}

			callback(v.Body[0].(dbus.ObjectPath))
		}
	}()

}

func (obj Upower) ConnectDeviceRemoved(callback func(device dbus.ObjectPath)) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+string(obj.core.Path())+"', interface='org.freedesktop.UPower',sender='org.freedesktop.UPower',member='DeviceRemoved'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range obj.signal_chan {
			if v.Name != "org.freedesktop.UPower.DeviceRemoved" || 1 != len(v.Body) {
				continue
			}

			if reflect.TypeOf(v.Body[0]) != reflect.TypeOf((*dbus.ObjectPath)(nil)).Elem() {
				continue
			}

			callback(v.Body[0].(dbus.ObjectPath))
		}
	}()

}

func (obj Upower) ConnectDeviceChanged(callback func(device dbus.ObjectPath)) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+string(obj.core.Path())+"', interface='org.freedesktop.UPower',sender='org.freedesktop.UPower',member='DeviceChanged'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range obj.signal_chan {
			if v.Name != "org.freedesktop.UPower.DeviceChanged" || 1 != len(v.Body) {
				continue
			}

			if reflect.TypeOf(v.Body[0]) != reflect.TypeOf((*dbus.ObjectPath)(nil)).Elem() {
				continue
			}

			callback(v.Body[0].(dbus.ObjectPath))
		}
	}()

}

func (obj Upower) ConnectChanged(callback func()) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+string(obj.core.Path())+"', interface='org.freedesktop.UPower',sender='org.freedesktop.UPower',member='Changed'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range obj.signal_chan {
			if v.Name != "org.freedesktop.UPower.Changed" || 0 != len(v.Body) {
				continue
			}
			_ = reflect.TypeOf /*prevent compile error*/

			callback()
		}
	}()

}

func (obj Upower) ConnectSleeping(callback func()) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+string(obj.core.Path())+"', interface='org.freedesktop.UPower',sender='org.freedesktop.UPower',member='Sleeping'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range obj.signal_chan {
			if v.Name != "org.freedesktop.UPower.Sleeping" || 0 != len(v.Body) {
				continue
			}
			_ = reflect.TypeOf /*prevent compile error*/

			callback()
		}
	}()

}

func (obj Upower) ConnectNotifySleep(callback func(action string)) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+string(obj.core.Path())+"', interface='org.freedesktop.UPower',sender='org.freedesktop.UPower',member='NotifySleep'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range obj.signal_chan {
			if v.Name != "org.freedesktop.UPower.NotifySleep" || 1 != len(v.Body) {
				continue
			}

			if reflect.TypeOf(v.Body[0]) != reflect.TypeOf((*string)(nil)).Elem() {
				continue
			}

			callback(v.Body[0].(string))
		}
	}()

}

func (obj Upower) ConnectResuming(callback func()) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+string(obj.core.Path())+"', interface='org.freedesktop.UPower',sender='org.freedesktop.UPower',member='Resuming'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range obj.signal_chan {
			if v.Name != "org.freedesktop.UPower.Resuming" || 0 != len(v.Body) {
				continue
			}
			_ = reflect.TypeOf /*prevent compile error*/

			callback()
		}
	}()

}

func (obj Upower) ConnectNotifyResume(callback func(action string)) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+string(obj.core.Path())+"', interface='org.freedesktop.UPower',sender='org.freedesktop.UPower',member='NotifyResume'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range obj.signal_chan {
			if v.Name != "org.freedesktop.UPower.NotifyResume" || 1 != len(v.Body) {
				continue
			}

			if reflect.TypeOf(v.Body[0]) != reflect.TypeOf((*string)(nil)).Elem() {
				continue
			}

			callback(v.Body[0].(string))
		}
	}()

}

func (obj *Upower) SetDaemonVersion(DaemonVersion string) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower", "DaemonVersion", DaemonVersion)
}
func (obj Upower) GetDaemonVersion() (ret string) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower", "DaemonVersion").Store(&r)
	if err == nil && r.Signature().String() == "s" {
		return r.Value().(string)
	} else {
		panic(err)
	}
	return
}

func (obj *Upower) SetCanSuspend(CanSuspend bool) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower", "CanSuspend", CanSuspend)
}
func (obj Upower) GetCanSuspend() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower", "CanSuspend").Store(&r)
	if err == nil && r.Signature().String() == "b" {
		return r.Value().(bool)
	} else {
		panic(err)
	}
	return
}

func (obj *Upower) SetCanHibernate(CanHibernate bool) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower", "CanHibernate", CanHibernate)
}
func (obj Upower) GetCanHibernate() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower", "CanHibernate").Store(&r)
	if err == nil && r.Signature().String() == "b" {
		return r.Value().(bool)
	} else {
		panic(err)
	}
	return
}

func (obj *Upower) SetOnBattery(OnBattery bool) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower", "OnBattery", OnBattery)
}
func (obj Upower) GetOnBattery() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower", "OnBattery").Store(&r)
	if err == nil && r.Signature().String() == "b" {
		return r.Value().(bool)
	} else {
		panic(err)
	}
	return
}

func (obj *Upower) SetOnLowBattery(OnLowBattery bool) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower", "OnLowBattery", OnLowBattery)
}
func (obj Upower) GetOnLowBattery() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower", "OnLowBattery").Store(&r)
	if err == nil && r.Signature().String() == "b" {
		return r.Value().(bool)
	} else {
		panic(err)
	}
	return
}

func (obj *Upower) SetLidIsClosed(LidIsClosed bool) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower", "LidIsClosed", LidIsClosed)
}
func (obj Upower) GetLidIsClosed() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower", "LidIsClosed").Store(&r)
	if err == nil && r.Signature().String() == "b" {
		return r.Value().(bool)
	} else {
		panic(err)
	}
	return
}

func (obj *Upower) SetLidIsPresent(LidIsPresent bool) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower", "LidIsPresent", LidIsPresent)
}
func (obj Upower) GetLidIsPresent() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower", "LidIsPresent").Store(&r)
	if err == nil && r.Signature().String() == "b" {
		return r.Value().(bool)
	} else {
		panic(err)
	}
	return
}

func (obj *Upower) SetLidForceSleep(LidForceSleep bool) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower", "LidForceSleep", LidForceSleep)
}
func (obj Upower) GetLidForceSleep() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower", "LidForceSleep").Store(&r)
	if err == nil && r.Signature().String() == "b" {
		return r.Value().(bool)
	} else {
		panic(err)
	}
	return
}

func (obj *Upower) SetIsDocked(IsDocked bool) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower", "IsDocked", IsDocked)
}
func (obj Upower) GetIsDocked() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower", "IsDocked").Store(&r)
	if err == nil && r.Signature().String() == "b" {
		return r.Value().(bool)
	} else {
		panic(err)
	}
	return
}

func GetUpower(path string) *Upower {
	return &Upower{dbus.ObjectPath(path), getBus().Object("org.freedesktop.UPower", dbus.ObjectPath(path)), make(chan *dbus.Signal)}
}
