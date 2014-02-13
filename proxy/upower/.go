/*This file is auto generate by dlib/dbus/proxyer. Don't edit it*/
package upower
import "dlib/dbus"
import "reflect"
var _ = reflect.TypeOf /*prevent compile error*/

type Upower struct {
	Path dbus.ObjectPath
	core *dbus.Object
	signal_chan chan *dbus.Signal
}


func (obj Upower) EnumerateDevices () (devices []dbus.ObjectPath) {
	obj.core.Call("org.freedesktop.UPower.EnumerateDevices", 0).Store(&devices)
	return
}

func (obj Upower) AboutToSleep (action string) () {
	obj.core.Call("org.freedesktop.UPower.AboutToSleep", 0, action).Store()
	return
}

func (obj Upower) Suspend () () {
	obj.core.Call("org.freedesktop.UPower.Suspend", 0).Store()
	return
}

func (obj Upower) SuspendAllowed () (allowed bool) {
	obj.core.Call("org.freedesktop.UPower.SuspendAllowed", 0).Store(&allowed)
	return
}

func (obj Upower) Hibernate () () {
	obj.core.Call("org.freedesktop.UPower.Hibernate", 0).Store()
	return
}

func (obj Upower) HibernateAllowed () (allowed bool) {
	obj.core.Call("org.freedesktop.UPower.HibernateAllowed", 0).Store(&allowed)
	return
}



func (obj Upower) ConnectDeviceAdded(callback func(device dbus.ObjectPath)) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+string(obj.core.Path())+"', interface='org.freedesktop.UPower',sender='org.freedesktop.UPower',member='DeviceAdded'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range(obj.signal_chan) {
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
		for v := range(obj.signal_chan) {
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
		for v := range(obj.signal_chan) {
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
		for v := range(obj.signal_chan) {
			if v.Name != "org.freedesktop.UPower.Changed" || 0 != len(v.Body) {
				continue
			}
			

			callback()
		}
	}()

}

func (obj Upower) ConnectSleeping(callback func()) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+string(obj.core.Path())+"', interface='org.freedesktop.UPower',sender='org.freedesktop.UPower',member='Sleeping'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range(obj.signal_chan) {
			if v.Name != "org.freedesktop.UPower.Sleeping" || 0 != len(v.Body) {
				continue
			}
			

			callback()
		}
	}()

}

func (obj Upower) ConnectNotifySleep(callback func(action string)) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+string(obj.core.Path())+"', interface='org.freedesktop.UPower',sender='org.freedesktop.UPower',member='NotifySleep'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range(obj.signal_chan) {
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
		for v := range(obj.signal_chan) {
			if v.Name != "org.freedesktop.UPower.Resuming" || 0 != len(v.Body) {
				continue
			}
			

			callback()
		}
	}()

}

func (obj Upower) ConnectNotifyResume(callback func(action string)) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+string(obj.core.Path())+"', interface='org.freedesktop.UPower',sender='org.freedesktop.UPower',member='NotifyResume'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range(obj.signal_chan) {
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





func (obj Upower) GetDaemonVersion() (ret string) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower", "DaemonVersion").Store(&r)
	if err == nil && r.Signature().String() == "s" { 
		return r.Value().(string)
	}  else {
		panic(err)
	}
	return
}



func (obj Upower) GetCanSuspend() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower", "CanSuspend").Store(&r)
	if err == nil && r.Signature().String() == "b" { 
		return r.Value().(bool)
	}  else {
		panic(err)
	}
	return
}



func (obj Upower) GetCanHibernate() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower", "CanHibernate").Store(&r)
	if err == nil && r.Signature().String() == "b" { 
		return r.Value().(bool)
	}  else {
		panic(err)
	}
	return
}



func (obj Upower) GetOnBattery() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower", "OnBattery").Store(&r)
	if err == nil && r.Signature().String() == "b" { 
		return r.Value().(bool)
	}  else {
		panic(err)
	}
	return
}



func (obj Upower) GetOnLowBattery() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower", "OnLowBattery").Store(&r)
	if err == nil && r.Signature().String() == "b" { 
		return r.Value().(bool)
	}  else {
		panic(err)
	}
	return
}



func (obj Upower) GetLidIsClosed() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower", "LidIsClosed").Store(&r)
	if err == nil && r.Signature().String() == "b" { 
		return r.Value().(bool)
	}  else {
		panic(err)
	}
	return
}



func (obj Upower) GetLidIsPresent() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower", "LidIsPresent").Store(&r)
	if err == nil && r.Signature().String() == "b" { 
		return r.Value().(bool)
	}  else {
		panic(err)
	}
	return
}



func (obj Upower) GetLidForceSleep() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower", "LidForceSleep").Store(&r)
	if err == nil && r.Signature().String() == "b" { 
		return r.Value().(bool)
	}  else {
		panic(err)
	}
	return
}



func (obj Upower) GetIsDocked() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower", "IsDocked").Store(&r)
	if err == nil && r.Signature().String() == "b" { 
		return r.Value().(bool)
	}  else {
		panic(err)
	}
	return
}


func GetUpower(path string) *Upower {
	return  &Upower{dbus.ObjectPath(path), getBus().Object("org.freedesktop.UPower", dbus.ObjectPath(path)),make(chan *dbus.Signal)}
}


type QoS struct {
	Path dbus.ObjectPath
	core *dbus.Object
	signal_chan chan *dbus.Signal
}


func (obj QoS) GetLatencyRequests () (requests [][]interface {}) {
	obj.core.Call("org.freedesktop.UPower.QoS.GetLatencyRequests", 0).Store(&requests)
	return
}

func (obj QoS) GetLatency (type_ string) (value int32) {
	obj.core.Call("org.freedesktop.UPower.QoS.GetLatency", 0, type_).Store(&value)
	return
}

func (obj QoS) CancelRequest (cookie uint32) () {
	obj.core.Call("org.freedesktop.UPower.QoS.CancelRequest", 0, cookie).Store()
	return
}

func (obj QoS) RequestLatency (type_ string,value int32,persistent bool) (cookie uint32) {
	obj.core.Call("org.freedesktop.UPower.QoS.RequestLatency", 0, type_, value, persistent).Store(&cookie)
	return
}

func (obj QoS) SetMinimumLatency (type_ string,value int32) () {
	obj.core.Call("org.freedesktop.UPower.QoS.SetMinimumLatency", 0, type_, value).Store()
	return
}



func (obj QoS) ConnectRequestsChanged(callback func()) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+string(obj.core.Path())+"', interface='org.freedesktop.UPower.QoS',sender='org.freedesktop.UPower',member='RequestsChanged'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range(obj.signal_chan) {
			if v.Name != "org.freedesktop.UPower.QoS.RequestsChanged" || 0 != len(v.Body) {
				continue
			}
			

			callback()
		}
	}()

}

func (obj QoS) ConnectLatencyChanged(callback func(arg0 string,arg1 int32)) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+string(obj.core.Path())+"', interface='org.freedesktop.UPower.QoS',sender='org.freedesktop.UPower',member='LatencyChanged'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range(obj.signal_chan) {
			if v.Name != "org.freedesktop.UPower.QoS.LatencyChanged" || 2 != len(v.Body) {
				continue
			}
			if reflect.TypeOf(v.Body[0]) != reflect.TypeOf((*string)(nil)).Elem() {
				continue
			}
			if reflect.TypeOf(v.Body[0]) != reflect.TypeOf((*int32)(nil)).Elem() {
				continue
			}
			

			callback(v.Body[0].(string),v.Body[1].(int32))
		}
	}()

}




func GetQoS(path string) *QoS {
	return  &QoS{dbus.ObjectPath(path), getBus().Object("org.freedesktop.UPower", dbus.ObjectPath(path)),make(chan *dbus.Signal)}
}


type Wakeups struct {
	Path dbus.ObjectPath
	core *dbus.Object
	signal_chan chan *dbus.Signal
}


func (obj Wakeups) GetData () (data [][]interface {}) {
	obj.core.Call("org.freedesktop.UPower.Wakeups.GetData", 0).Store(&data)
	return
}

func (obj Wakeups) GetTotal () (value uint32) {
	obj.core.Call("org.freedesktop.UPower.Wakeups.GetTotal", 0).Store(&value)
	return
}



func (obj Wakeups) ConnectDataChanged(callback func()) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+string(obj.core.Path())+"', interface='org.freedesktop.UPower.Wakeups',sender='org.freedesktop.UPower',member='DataChanged'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range(obj.signal_chan) {
			if v.Name != "org.freedesktop.UPower.Wakeups.DataChanged" || 0 != len(v.Body) {
				continue
			}
			

			callback()
		}
	}()

}

func (obj Wakeups) ConnectTotalChanged(callback func(arg0 uint32)) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+string(obj.core.Path())+"', interface='org.freedesktop.UPower.Wakeups',sender='org.freedesktop.UPower',member='TotalChanged'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range(obj.signal_chan) {
			if v.Name != "org.freedesktop.UPower.Wakeups.TotalChanged" || 1 != len(v.Body) {
				continue
			}
			if reflect.TypeOf(v.Body[0]) != reflect.TypeOf((*uint32)(nil)).Elem() {
				continue
			}
			

			callback(v.Body[0].(uint32))
		}
	}()

}





func (obj Wakeups) GetHasCapability() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Wakeups", "HasCapability").Store(&r)
	if err == nil && r.Signature().String() == "b" { 
		return r.Value().(bool)
	}  else {
		panic(err)
	}
	return
}


func GetWakeups(path string) *Wakeups {
	return  &Wakeups{dbus.ObjectPath(path), getBus().Object("org.freedesktop.UPower", dbus.ObjectPath(path)),make(chan *dbus.Signal)}
}


type Device struct {
	Path dbus.ObjectPath
	core *dbus.Object
	signal_chan chan *dbus.Signal
}


func (obj Device) GetStatistics (type_ string) (data [][]interface {}) {
	obj.core.Call("org.freedesktop.UPower.Device.GetStatistics", 0, type_).Store(&data)
	return
}

func (obj Device) GetHistory (type_ string,timespan uint32,resolution uint32) (data [][]interface {}) {
	obj.core.Call("org.freedesktop.UPower.Device.GetHistory", 0, type_, timespan, resolution).Store(&data)
	return
}

func (obj Device) Refresh () () {
	obj.core.Call("org.freedesktop.UPower.Device.Refresh", 0).Store()
	return
}



func (obj Device) ConnectChanged(callback func()) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+string(obj.core.Path())+"', interface='org.freedesktop.UPower.Device',sender='org.freedesktop.UPower',member='Changed'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range(obj.signal_chan) {
			if v.Name != "org.freedesktop.UPower.Device.Changed" || 0 != len(v.Body) {
				continue
			}
			

			callback()
		}
	}()

}





func (obj Device) GetRecallUrl() (ret string) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "RecallUrl").Store(&r)
	if err == nil && r.Signature().String() == "s" { 
		return r.Value().(string)
	}  else {
		panic(err)
	}
	return
}



func (obj Device) GetRecallVendor() (ret string) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "RecallVendor").Store(&r)
	if err == nil && r.Signature().String() == "s" { 
		return r.Value().(string)
	}  else {
		panic(err)
	}
	return
}



func (obj Device) GetRecallNotice() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "RecallNotice").Store(&r)
	if err == nil && r.Signature().String() == "b" { 
		return r.Value().(bool)
	}  else {
		panic(err)
	}
	return
}



func (obj Device) GetTechnology() (ret uint32) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "Technology").Store(&r)
	if err == nil && r.Signature().String() == "u" { 
		return r.Value().(uint32)
	}  else {
		panic(err)
	}
	return
}



func (obj Device) GetCapacity() (ret float64) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "Capacity").Store(&r)
	if err == nil && r.Signature().String() == "d" { 
		return r.Value().(float64)
	}  else {
		panic(err)
	}
	return
}



func (obj Device) GetIsRechargeable() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "IsRechargeable").Store(&r)
	if err == nil && r.Signature().String() == "b" { 
		return r.Value().(bool)
	}  else {
		panic(err)
	}
	return
}



func (obj Device) GetState() (ret uint32) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "State").Store(&r)
	if err == nil && r.Signature().String() == "u" { 
		return r.Value().(uint32)
	}  else {
		panic(err)
	}
	return
}



func (obj Device) GetIsPresent() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "IsPresent").Store(&r)
	if err == nil && r.Signature().String() == "b" { 
		return r.Value().(bool)
	}  else {
		panic(err)
	}
	return
}



func (obj Device) GetTemperature() (ret float64) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "Temperature").Store(&r)
	if err == nil && r.Signature().String() == "d" { 
		return r.Value().(float64)
	}  else {
		panic(err)
	}
	return
}



func (obj Device) GetPercentage() (ret float64) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "Percentage").Store(&r)
	if err == nil && r.Signature().String() == "d" { 
		return r.Value().(float64)
	}  else {
		panic(err)
	}
	return
}



func (obj Device) GetTimeToFull() (ret int64) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "TimeToFull").Store(&r)
	if err == nil && r.Signature().String() == "x" { 
		return r.Value().(int64)
	}  else {
		panic(err)
	}
	return
}



func (obj Device) GetTimeToEmpty() (ret int64) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "TimeToEmpty").Store(&r)
	if err == nil && r.Signature().String() == "x" { 
		return r.Value().(int64)
	}  else {
		panic(err)
	}
	return
}



func (obj Device) GetLuminosity() (ret float64) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "Luminosity").Store(&r)
	if err == nil && r.Signature().String() == "d" { 
		return r.Value().(float64)
	}  else {
		panic(err)
	}
	return
}



func (obj Device) GetVoltage() (ret float64) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "Voltage").Store(&r)
	if err == nil && r.Signature().String() == "d" { 
		return r.Value().(float64)
	}  else {
		panic(err)
	}
	return
}



func (obj Device) GetEnergyRate() (ret float64) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "EnergyRate").Store(&r)
	if err == nil && r.Signature().String() == "d" { 
		return r.Value().(float64)
	}  else {
		panic(err)
	}
	return
}



func (obj Device) GetEnergyFullDesign() (ret float64) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "EnergyFullDesign").Store(&r)
	if err == nil && r.Signature().String() == "d" { 
		return r.Value().(float64)
	}  else {
		panic(err)
	}
	return
}



func (obj Device) GetEnergyFull() (ret float64) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "EnergyFull").Store(&r)
	if err == nil && r.Signature().String() == "d" { 
		return r.Value().(float64)
	}  else {
		panic(err)
	}
	return
}



func (obj Device) GetEnergyEmpty() (ret float64) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "EnergyEmpty").Store(&r)
	if err == nil && r.Signature().String() == "d" { 
		return r.Value().(float64)
	}  else {
		panic(err)
	}
	return
}



func (obj Device) GetEnergy() (ret float64) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "Energy").Store(&r)
	if err == nil && r.Signature().String() == "d" { 
		return r.Value().(float64)
	}  else {
		panic(err)
	}
	return
}



func (obj Device) GetOnline() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "Online").Store(&r)
	if err == nil && r.Signature().String() == "b" { 
		return r.Value().(bool)
	}  else {
		panic(err)
	}
	return
}



func (obj Device) GetHasStatistics() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "HasStatistics").Store(&r)
	if err == nil && r.Signature().String() == "b" { 
		return r.Value().(bool)
	}  else {
		panic(err)
	}
	return
}



func (obj Device) GetHasHistory() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "HasHistory").Store(&r)
	if err == nil && r.Signature().String() == "b" { 
		return r.Value().(bool)
	}  else {
		panic(err)
	}
	return
}



func (obj Device) GetPowerSupply() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "PowerSupply").Store(&r)
	if err == nil && r.Signature().String() == "b" { 
		return r.Value().(bool)
	}  else {
		panic(err)
	}
	return
}



func (obj Device) GetType() (ret uint32) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "Type").Store(&r)
	if err == nil && r.Signature().String() == "u" { 
		return r.Value().(uint32)
	}  else {
		panic(err)
	}
	return
}



func (obj Device) GetUpdateTime() (ret uint64) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "UpdateTime").Store(&r)
	if err == nil && r.Signature().String() == "t" { 
		return r.Value().(uint64)
	}  else {
		panic(err)
	}
	return
}



func (obj Device) GetSerial() (ret string) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "Serial").Store(&r)
	if err == nil && r.Signature().String() == "s" { 
		return r.Value().(string)
	}  else {
		panic(err)
	}
	return
}



func (obj Device) GetModel() (ret string) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "Model").Store(&r)
	if err == nil && r.Signature().String() == "s" { 
		return r.Value().(string)
	}  else {
		panic(err)
	}
	return
}



func (obj Device) GetVendor() (ret string) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "Vendor").Store(&r)
	if err == nil && r.Signature().String() == "s" { 
		return r.Value().(string)
	}  else {
		panic(err)
	}
	return
}



func (obj Device) GetNativePath() (ret string) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "NativePath").Store(&r)
	if err == nil && r.Signature().String() == "s" { 
		return r.Value().(string)
	}  else {
		panic(err)
	}
	return
}


func GetDevice(path string) *Device {
	return  &Device{dbus.ObjectPath(path), getBus().Object("org.freedesktop.UPower", dbus.ObjectPath(path)),make(chan *dbus.Signal)}
}

