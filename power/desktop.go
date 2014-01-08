//+build ignore

package main

import (
	"dlib"
	"dlib/dbus"
	"dlib/dbus/property"
	"dlib/gio-2.0"
)

const (
	_DESKTOP_DEST = "com.deepin.daemon.Desktop"
	_DESKTOP_PATH = "/com/deepin/daemon/Desktop"
	_DESKTOP_IFC  = "com.deepin.daemon.Desktop"

	_DESKTOP_SCHEMA = "org.gnome.settings-daemon.plugins.power"
	_DOCK_SCHEMA    = "com.deepin.dde.dock"

	_COMPIZ_INTEGRATED_SCHEMA = "org.compiz.integrated"
	_COMPIZ_COMMANDS_SCHEMA   = "org.compiz.commands"
	_COMPIZ_SCALE_SCHEMA      = "org.compiz.scale"
	_COMPIZ_COMMAND_PATH      = "/org/compiz/profiles/deepin/plugins/commands/"
	_COMPIZ_SCALE_PATH        = "/org/compiz/profiles/deepin/plugins/scale/"

	_LAUNCHER_CMD = "launcher"
)

const (
	ACTION_NONE           = int32(0)
	ACTION_OPENED_WINDOWS = int32(1)
	ACTION_LAUNCHER       = int32(2)
)

type DesktopManager struct {
	BatteryIsPresent  bool     `access:"read"`       //battery present
	BatteryPercentage float64  `access:"read"`       //batter voltage
	charging          int32    `access:"read"`       //charging or discharging
	PlugedIn          int32    `access:"read"`       //1 for in,2 for out
	TimeToEmpty       int64    `access:"read"`       //
	TimeToFull        int64    `access:"read"`       //time to fully charged
	SuspendTime       []int32  `access:"read/write"` //with or without battery
	HandleLowPower    []string `access:"read/write"`
	HandleClosedLid   []string

	CurrentPlan dbus.Property
}

func (desk *DesktopManager) GetDBusInfo() dbus.DBusInfo {
	return dbus.DBusInfo{_DESKTOP_DEST, _DESKTOP_PATH, _DESKTOP_IFC}
}

func NewDesktopManager() (*DesktopManager, error) {
	desk := DesktopManager{}
	busConn, err := dbus.SessionBus()
	if err != nil {
		return nil, err
	}

	deskSettings := gio.NewSettings(_DESKTOP_SCHEMA)
	desk.CurrentPlan = property.NewGSettingsPropertyFull(
		deskSettings, "current-plan", "", busConn,
		_DESKTOP_PATH, _DESKTOP_IFC, "CurrentPlan")

	return &desk, nil
}

func main() {
	desk, err := NewDesktopManager()
	if err != nil {
		return
	}
	dbus.InstallOnSession(desk)
	dlib.StartLoop()
}
