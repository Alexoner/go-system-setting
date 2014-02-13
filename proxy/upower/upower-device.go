/*This file is auto generate by dlib/dbus/proxyer. Don't edit it*/

package upower

import "dlib/dbus"
import "reflect"

type Device struct {
	Path        dbus.ObjectPath
	core        *dbus.Object
	signal_chan chan *dbus.Signal
}

func (obj Device) GetStatistics(type_ string) (data [][]interface{}) {
	obj.core.Call("org.freedesktop.UPower.Device.GetStatistics", 0, type_).Store(&data)
	return
}

func (obj Device) GetHistory(type_ string, timespan uint32, resolution uint32) (data [][]interface{}) {
	obj.core.Call("org.freedesktop.UPower.Device.GetHistory", 0, type_, timespan, resolution).Store(&data)
	return
}

func (obj Device) Refresh() {
	obj.core.Call("org.freedesktop.UPower.Device.Refresh", 0).Store()
	return
}

func (obj Device) ConnectChanged(callback func()) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+string(obj.core.Path())+"', interface='org.freedesktop.UPower.Device',sender='org.freedesktop.UPower',member='Changed'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range obj.signal_chan {
			if v.Name != "org.freedesktop.UPower.Device.Changed" || 0 != len(v.Body) {
				continue
			}
			_ = reflect.TypeOf /*prevent compile error*/

			callback()
		}
	}()

}

func (obj *Device) SetRecallUrl(RecallUrl string) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "RecallUrl", RecallUrl)
}
func (obj Device) GetRecallUrl() (ret string) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "RecallUrl").Store(&r)
	if err == nil && r.Signature().String() == "s" {
		return r.Value().(string)
	} else {
		panic(err)
	}
	return
}

func (obj *Device) SetRecallVendor(RecallVendor string) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "RecallVendor", RecallVendor)
}
func (obj Device) GetRecallVendor() (ret string) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "RecallVendor").Store(&r)
	if err == nil && r.Signature().String() == "s" {
		return r.Value().(string)
	} else {
		panic(err)
	}
	return
}

func (obj *Device) SetRecallNotice(RecallNotice bool) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "RecallNotice", RecallNotice)
}
func (obj Device) GetRecallNotice() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "RecallNotice").Store(&r)
	if err == nil && r.Signature().String() == "b" {
		return r.Value().(bool)
	} else {
		panic(err)
	}
	return
}

func (obj *Device) SetTechnology(Technology uint32) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "Technology", Technology)
}
func (obj Device) GetTechnology() (ret uint32) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "Technology").Store(&r)
	if err == nil && r.Signature().String() == "u" {
		return r.Value().(uint32)
	} else {
		panic(err)
	}
	return
}

func (obj *Device) SetCapacity(Capacity float64) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "Capacity", Capacity)
}
func (obj Device) GetCapacity() (ret float64) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "Capacity").Store(&r)
	if err == nil && r.Signature().String() == "d" {
		return r.Value().(float64)
	} else {
		panic(err)
	}
	return
}

func (obj *Device) SetIsRechargeable(IsRechargeable bool) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "IsRechargeable", IsRechargeable)
}
func (obj Device) GetIsRechargeable() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "IsRechargeable").Store(&r)
	if err == nil && r.Signature().String() == "b" {
		return r.Value().(bool)
	} else {
		panic(err)
	}
	return
}

func (obj *Device) SetState(State uint32) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "State", State)
}
func (obj Device) GetState() (ret uint32) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "State").Store(&r)
	if err == nil && r.Signature().String() == "u" {
		return r.Value().(uint32)
	} else {
		panic(err)
	}
	return
}

func (obj *Device) SetIsPresent(IsPresent bool) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "IsPresent", IsPresent)
}
func (obj Device) GetIsPresent() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "IsPresent").Store(&r)
	if err == nil && r.Signature().String() == "b" {
		return r.Value().(bool)
	} else {
		panic(err)
	}
	return
}

func (obj *Device) SetTemperature(Temperature float64) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "Temperature", Temperature)
}
func (obj Device) GetTemperature() (ret float64) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "Temperature").Store(&r)
	if err == nil && r.Signature().String() == "d" {
		return r.Value().(float64)
	} else {
		panic(err)
	}
	return
}

func (obj *Device) SetPercentage(Percentage float64) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "Percentage", Percentage)
}
func (obj Device) GetPercentage() (ret float64) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "Percentage").Store(&r)
	if err == nil && r.Signature().String() == "d" {
		return r.Value().(float64)
	} else {
		panic(err)
	}
	return
}

func (obj *Device) SetTimeToFull(TimeToFull int64) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "TimeToFull", TimeToFull)
}
func (obj Device) GetTimeToFull() (ret int64) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "TimeToFull").Store(&r)
	if err == nil && r.Signature().String() == "x" {
		return r.Value().(int64)
	} else {
		panic(err)
	}
	return
}

func (obj *Device) SetTimeToEmpty(TimeToEmpty int64) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "TimeToEmpty", TimeToEmpty)
}
func (obj Device) GetTimeToEmpty() (ret int64) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "TimeToEmpty").Store(&r)
	if err == nil && r.Signature().String() == "x" {
		return r.Value().(int64)
	} else {
		panic(err)
	}
	return
}

func (obj *Device) SetLuminosity(Luminosity float64) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "Luminosity", Luminosity)
}
func (obj Device) GetLuminosity() (ret float64) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "Luminosity").Store(&r)
	if err == nil && r.Signature().String() == "d" {
		return r.Value().(float64)
	} else {
		panic(err)
	}
	return
}

func (obj *Device) SetVoltage(Voltage float64) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "Voltage", Voltage)
}
func (obj Device) GetVoltage() (ret float64) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "Voltage").Store(&r)
	if err == nil && r.Signature().String() == "d" {
		return r.Value().(float64)
	} else {
		panic(err)
	}
	return
}

func (obj *Device) SetEnergyRate(EnergyRate float64) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "EnergyRate", EnergyRate)
}
func (obj Device) GetEnergyRate() (ret float64) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "EnergyRate").Store(&r)
	if err == nil && r.Signature().String() == "d" {
		return r.Value().(float64)
	} else {
		panic(err)
	}
	return
}

func (obj *Device) SetEnergyFullDesign(EnergyFullDesign float64) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "EnergyFullDesign", EnergyFullDesign)
}
func (obj Device) GetEnergyFullDesign() (ret float64) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "EnergyFullDesign").Store(&r)
	if err == nil && r.Signature().String() == "d" {
		return r.Value().(float64)
	} else {
		panic(err)
	}
	return
}

func (obj *Device) SetEnergyFull(EnergyFull float64) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "EnergyFull", EnergyFull)
}
func (obj Device) GetEnergyFull() (ret float64) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "EnergyFull").Store(&r)
	if err == nil && r.Signature().String() == "d" {
		return r.Value().(float64)
	} else {
		panic(err)
	}
	return
}

func (obj *Device) SetEnergyEmpty(EnergyEmpty float64) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "EnergyEmpty", EnergyEmpty)
}
func (obj Device) GetEnergyEmpty() (ret float64) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "EnergyEmpty").Store(&r)
	if err == nil && r.Signature().String() == "d" {
		return r.Value().(float64)
	} else {
		panic(err)
	}
	return
}

func (obj *Device) SetEnergy(Energy float64) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "Energy", Energy)
}
func (obj Device) GetEnergy() (ret float64) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "Energy").Store(&r)
	if err == nil && r.Signature().String() == "d" {
		return r.Value().(float64)
	} else {
		panic(err)
	}
	return
}

func (obj *Device) SetOnline(Online bool) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "Online", Online)
}
func (obj Device) GetOnline() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "Online").Store(&r)
	if err == nil && r.Signature().String() == "b" {
		return r.Value().(bool)
	} else {
		panic(err)
	}
	return
}

func (obj *Device) SetHasStatistics(HasStatistics bool) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "HasStatistics", HasStatistics)
}
func (obj Device) GetHasStatistics() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "HasStatistics").Store(&r)
	if err == nil && r.Signature().String() == "b" {
		return r.Value().(bool)
	} else {
		panic(err)
	}
	return
}

func (obj *Device) SetHasHistory(HasHistory bool) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "HasHistory", HasHistory)
}
func (obj Device) GetHasHistory() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "HasHistory").Store(&r)
	if err == nil && r.Signature().String() == "b" {
		return r.Value().(bool)
	} else {
		panic(err)
	}
	return
}

func (obj *Device) SetPowerSupply(PowerSupply bool) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "PowerSupply", PowerSupply)
}
func (obj Device) GetPowerSupply() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "PowerSupply").Store(&r)
	if err == nil && r.Signature().String() == "b" {
		return r.Value().(bool)
	} else {
		panic(err)
	}
	return
}

func (obj *Device) SetType(Type uint32) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "Type", Type)
}
func (obj Device) GetType() (ret uint32) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "Type").Store(&r)
	if err == nil && r.Signature().String() == "u" {
		return r.Value().(uint32)
	} else {
		panic(err)
	}
	return
}

func (obj *Device) SetUpdateTime(UpdateTime uint64) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "UpdateTime", UpdateTime)
}
func (obj Device) GetUpdateTime() (ret uint64) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "UpdateTime").Store(&r)
	if err == nil && r.Signature().String() == "t" {
		return r.Value().(uint64)
	} else {
		panic(err)
	}
	return
}

func (obj *Device) SetSerial(Serial string) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "Serial", Serial)
}
func (obj Device) GetSerial() (ret string) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "Serial").Store(&r)
	if err == nil && r.Signature().String() == "s" {
		return r.Value().(string)
	} else {
		panic(err)
	}
	return
}

func (obj *Device) SetModel(Model string) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "Model", Model)
}
func (obj Device) GetModel() (ret string) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "Model").Store(&r)
	if err == nil && r.Signature().String() == "s" {
		return r.Value().(string)
	} else {
		panic(err)
	}
	return
}

func (obj *Device) SetVendor(Vendor string) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "Vendor", Vendor)
}
func (obj Device) GetVendor() (ret string) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "Vendor").Store(&r)
	if err == nil && r.Signature().String() == "s" {
		return r.Value().(string)
	} else {
		panic(err)
	}
	return
}

func (obj *Device) SetNativePath(NativePath string) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "NativePath", NativePath)
}
func (obj Device) GetNativePath() (ret string) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "NativePath").Store(&r)
	if err == nil && r.Signature().String() == "s" {
		return r.Value().(string)
	} else {
		panic(err)
	}
	return
}

func GetDevice(path string) *Device {
	return &Device{dbus.ObjectPath(path), getBus().Object("org.freedesktop.UPower", dbus.ObjectPath(path)), make(chan *dbus.Signal)}
}
