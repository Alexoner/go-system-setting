// +build ignore

package main

import (
	"dlib/dbus"
	"upower"
)

type battery struct {
	bus_name    string
	object_path string

	//device upower.device
}

const (
	operation_suspend   = "suspend"
	operation_poweroff  = "poweroff"
	operation_hibernate = "hibernate"
)

const (
	suspend_time_0  = 0 //don't suspend
	suspend_time_5  = 5
	suspend_time_10 = 10
	suspend_time_30 = 30
	suspend_time_60 = 60 //after one hour
)

type Power struct {
	BatteryIsPresent  bool     `access:"read"`       //battery present
	BatteryPercentage float64  `access:"read"`       //batter voltage
	charging          int32    `access:"read"`       //charging or discharging
	PlugedIn          int32    `access:"read"`       //power pluged in
	TimeToEmpty       int64    `access:"read"`       //
	TimeToFull        int64    `access:"read"`       //time to fully charged
	SuspendTime       []int32  `access:"read/write"` //with or without battery
	HandleLowPower    []string `access:"read/write"`
	HandleClosedLid   []string
}

func (p *Power) GetDBusInfo() dbus.DBusInfo {
	return dbus.DBusInfo{
		"com.deepin.daemon.Power",  //bus name
		"/com/deepin/daemon/Power", //object path
		"com.deepin.daemon.Power",
	}
}

func (p *Power) Refresh() int32 {
	device := upower.Getdevice("/org/freedesktop/UPower/devices/battery_BAT0")
	p.BatteryPercentage = device.GetTemperature()
	p.BatteryPercentage = device.GetPercentage()
	//p.charging=
	p.Plugedin = device.GetPowerSupply()
	p.TimeToEmpty = device.GetTimeToEmpty()
	p.TimeToFull = device.GetTimeToFull()
}

func main() {
	dbus.InstallOnSession(&Power{})
	select {}
}
