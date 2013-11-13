/*This file is auto generate by dlib/dbus/proxyer. Don't edit it*/

package upower
import "dlib/dbus"

type device struct {
	core *dbus.Object
	signal_chan chan *dbus.Signal
}


func (obj device) GetStatistics (type string) (data [][]interface {}) {
	obj.core.Call("org.freedesktop.UPower.Device.GetStatistics", 0, type).Store(&data)
	return
}

func (obj device) GetHistory (type string,timespan uint32,resolution uint32) (data [][]interface {}) {
	obj.core.Call("org.freedesktop.UPower.Device.GetHistory", 0, type, timespan, resolution).Store(&data)
	return
}

func (obj device) Refresh () () {
	obj.core.Call("org.freedesktop.UPower.Device.Refresh", 0).Store()
	return
}



func (obj device) ConnectChanged(callback func()) {
	__conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+obj.core.Path()+"', interface='org.freedesktop.UPower.Device',sender='org.freedesktop.UPower'")
	__conn.Signal(obj.signal_chan)
	go func() {
		for v := range(obj.signal_chan) {
			v = v
			/*callback(v.Body...)*/
		}
	}()

}



func (obj *device) SetRecallUrl(RecallUrl string) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "RecallUrl", RecallUrl)
}
func (obj device) GetRecallUrl() (ret string) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "RecallUrl").Store(&r)
	if err == nil && r.Signature().String() == "s" {
		return r.Value().(string)
	}  else {
		panic(err)
	}
	return
}

func (obj *device) SetRecallVendor(RecallVendor string) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "RecallVendor", RecallVendor)
}
func (obj device) GetRecallVendor() (ret string) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "RecallVendor").Store(&r)
	if err == nil && r.Signature().String() == "s" {
		return r.Value().(string)
	}  else {
		panic(err)
	}
	return
}

func (obj *device) SetRecallNotice(RecallNotice bool) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "RecallNotice", RecallNotice)
}
func (obj device) GetRecallNotice() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "RecallNotice").Store(&r)
	if err == nil && r.Signature().String() == "b" {
		return r.Value().(bool)
	}  else {
		panic(err)
	}
	return
}

func (obj *device) SetTechnology(Technology uint32) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "Technology", Technology)
}
func (obj device) GetTechnology() (ret uint32) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "Technology").Store(&r)
	if err == nil && r.Signature().String() == "u" {
		return r.Value().(uint32)
	}  else {
		panic(err)
	}
	return
}

func (obj *device) SetCapacity(Capacity float64) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "Capacity", Capacity)
}
func (obj device) GetCapacity() (ret float64) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "Capacity").Store(&r)
	if err == nil && r.Signature().String() == "d" {
		return r.Value().(float64)
	}  else {
		panic(err)
	}
	return
}

func (obj *device) SetIsRechargeable(IsRechargeable bool) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "IsRechargeable", IsRechargeable)
}
func (obj device) GetIsRechargeable() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "IsRechargeable").Store(&r)
	if err == nil && r.Signature().String() == "b" {
		return r.Value().(bool)
	}  else {
		panic(err)
	}
	return
}

func (obj *device) SetState(State uint32) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "State", State)
}
func (obj device) GetState() (ret uint32) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "State").Store(&r)
	if err == nil && r.Signature().String() == "u" {
		return r.Value().(uint32)
	}  else {
		panic(err)
	}
	return
}

func (obj *device) SetIsPresent(IsPresent bool) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "IsPresent", IsPresent)
}
func (obj device) GetIsPresent() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "IsPresent").Store(&r)
	if err == nil && r.Signature().String() == "b" {
		return r.Value().(bool)
	}  else {
		panic(err)
	}
	return
}

func (obj *device) SetTemperature(Temperature float64) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "Temperature", Temperature)
}
func (obj device) GetTemperature() (ret float64) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "Temperature").Store(&r)
	if err == nil && r.Signature().String() == "d" {
		return r.Value().(float64)
	}  else {
		panic(err)
	}
	return
}

func (obj *device) SetPercentage(Percentage float64) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "Percentage", Percentage)
}
func (obj device) GetPercentage() (ret float64) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "Percentage").Store(&r)
	if err == nil && r.Signature().String() == "d" {
		return r.Value().(float64)
	}  else {
		panic(err)
	}
	return
}

func (obj *device) SetTimeToFull(TimeToFull int64) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "TimeToFull", TimeToFull)
}
func (obj device) GetTimeToFull() (ret int64) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "TimeToFull").Store(&r)
	if err == nil && r.Signature().String() == "x" {
		return r.Value().(int64)
	}  else {
		panic(err)
	}
	return
}

func (obj *device) SetTimeToEmpty(TimeToEmpty int64) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "TimeToEmpty", TimeToEmpty)
}
func (obj device) GetTimeToEmpty() (ret int64) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "TimeToEmpty").Store(&r)
	if err == nil && r.Signature().String() == "x" {
		return r.Value().(int64)
	}  else {
		panic(err)
	}
	return
}

func (obj *device) SetLuminosity(Luminosity float64) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "Luminosity", Luminosity)
}
func (obj device) GetLuminosity() (ret float64) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "Luminosity").Store(&r)
	if err == nil && r.Signature().String() == "d" {
		return r.Value().(float64)
	}  else {
		panic(err)
	}
	return
}

func (obj *device) SetVoltage(Voltage float64) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "Voltage", Voltage)
}
func (obj device) GetVoltage() (ret float64) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "Voltage").Store(&r)
	if err == nil && r.Signature().String() == "d" {
		return r.Value().(float64)
	}  else {
		panic(err)
	}
	return
}

func (obj *device) SetEnergyRate(EnergyRate float64) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "EnergyRate", EnergyRate)
}
func (obj device) GetEnergyRate() (ret float64) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "EnergyRate").Store(&r)
	if err == nil && r.Signature().String() == "d" {
		return r.Value().(float64)
	}  else {
		panic(err)
	}
	return
}

func (obj *device) SetEnergyFullDesign(EnergyFullDesign float64) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "EnergyFullDesign", EnergyFullDesign)
}
func (obj device) GetEnergyFullDesign() (ret float64) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "EnergyFullDesign").Store(&r)
	if err == nil && r.Signature().String() == "d" {
		return r.Value().(float64)
	}  else {
		panic(err)
	}
	return
}

func (obj *device) SetEnergyFull(EnergyFull float64) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "EnergyFull", EnergyFull)
}
func (obj device) GetEnergyFull() (ret float64) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "EnergyFull").Store(&r)
	if err == nil && r.Signature().String() == "d" {
		return r.Value().(float64)
	}  else {
		panic(err)
	}
	return
}

func (obj *device) SetEnergyEmpty(EnergyEmpty float64) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "EnergyEmpty", EnergyEmpty)
}
func (obj device) GetEnergyEmpty() (ret float64) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "EnergyEmpty").Store(&r)
	if err == nil && r.Signature().String() == "d" {
		return r.Value().(float64)
	}  else {
		panic(err)
	}
	return
}

func (obj *device) SetEnergy(Energy float64) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "Energy", Energy)
}
func (obj device) GetEnergy() (ret float64) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "Energy").Store(&r)
	if err == nil && r.Signature().String() == "d" {
		return r.Value().(float64)
	}  else {
		panic(err)
	}
	return
}

func (obj *device) SetOnline(Online bool) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "Online", Online)
}
func (obj device) GetOnline() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "Online").Store(&r)
	if err == nil && r.Signature().String() == "b" {
		return r.Value().(bool)
	}  else {
		panic(err)
	}
	return
}

func (obj *device) SetHasStatistics(HasStatistics bool) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "HasStatistics", HasStatistics)
}
func (obj device) GetHasStatistics() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "HasStatistics").Store(&r)
	if err == nil && r.Signature().String() == "b" {
		return r.Value().(bool)
	}  else {
		panic(err)
	}
	return
}

func (obj *device) SetHasHistory(HasHistory bool) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "HasHistory", HasHistory)
}
func (obj device) GetHasHistory() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "HasHistory").Store(&r)
	if err == nil && r.Signature().String() == "b" {
		return r.Value().(bool)
	}  else {
		panic(err)
	}
	return
}

func (obj *device) SetPowerSupply(PowerSupply bool) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "PowerSupply", PowerSupply)
}
func (obj device) GetPowerSupply() (ret bool) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "PowerSupply").Store(&r)
	if err == nil && r.Signature().String() == "b" {
		return r.Value().(bool)
	}  else {
		panic(err)
	}
	return
}

func (obj *device) SetType(Type uint32) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "Type", Type)
}
func (obj device) GetType() (ret uint32) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "Type").Store(&r)
	if err == nil && r.Signature().String() == "u" {
		return r.Value().(uint32)
	}  else {
		panic(err)
	}
	return
}

func (obj *device) SetUpdateTime(UpdateTime uint64) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "UpdateTime", UpdateTime)
}
func (obj device) GetUpdateTime() (ret uint64) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "UpdateTime").Store(&r)
	if err == nil && r.Signature().String() == "t" {
		return r.Value().(uint64)
	}  else {
		panic(err)
	}
	return
}

func (obj *device) SetSerial(Serial string) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "Serial", Serial)
}
func (obj device) GetSerial() (ret string) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "Serial").Store(&r)
	if err == nil && r.Signature().String() == "s" {
		return r.Value().(string)
	}  else {
		panic(err)
	}
	return
}

func (obj *device) SetModel(Model string) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "Model", Model)
}
func (obj device) GetModel() (ret string) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "Model").Store(&r)
	if err == nil && r.Signature().String() == "s" {
		return r.Value().(string)
	}  else {
		panic(err)
	}
	return
}

func (obj *device) SetVendor(Vendor string) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "Vendor", Vendor)
}
func (obj device) GetVendor() (ret string) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "Vendor").Store(&r)
	if err == nil && r.Signature().String() == "s" {
		return r.Value().(string)
	}  else {
		panic(err)
	}
	return
}

func (obj *device) SetNativePath(NativePath string) {
	obj.core.Call("org.freedesktop.DBus.Properties.Set", 0, "org.freedesktop.UPower.Device", "NativePath", NativePath)
}
func (obj device) GetNativePath() (ret string) {
	var r dbus.Variant
	err := obj.core.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "NativePath").Store(&r)
	if err == nil && r.Signature().String() == "s" {
		return r.Value().(string)
	}  else {
		panic(err)
	}
	return
}


func Getdevice(path string) *device {
	return  &device{getBus().Object("org.freedesktop.UPower", dbus.ObjectPath(path)),make(chan *dbus.Signal)}
}

