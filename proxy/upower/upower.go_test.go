/*This file is auto generate by dlib/dbus/proxyer. Don't edit it*/
package upower

import "testing"

func TestUpowerMethodEnumerateDevices(t *testing.T) {

}

func TestUpowerMethodAboutToSleep(t *testing.T) {

}

func TestUpowerMethodSuspend(t *testing.T) {

}

func TestUpowerMethodSuspendAllowed(t *testing.T) {

}

func TestUpowerMethodHibernate(t *testing.T) {

}

func TestUpowerMethodHibernateAllowed(t *testing.T) {

}

func TestUpowerPropertyDaemonVersion(t *testing.T) {
	t.Log("Get the property DaemonVersion of object Upower ===> ",
		GetUpower("/org/freedesktop/UPower").GetDaemonVersion())
}

func TestUpowerPropertyCanSuspend(t *testing.T) {
	t.Log("Get the property CanSuspend of object Upower ===> ",
		GetUpower("/org/freedesktop/UPower").GetCanSuspend())
}

func TestUpowerPropertyCanHibernate(t *testing.T) {
	t.Log("Get the property CanHibernate of object Upower ===> ",
		GetUpower("/org/freedesktop/UPower").GetCanHibernate())
}

func TestUpowerPropertyOnBattery(t *testing.T) {
	t.Log("Get the property OnBattery of object Upower ===> ",
		GetUpower("/org/freedesktop/UPower").GetOnBattery())
}

func TestUpowerPropertyOnLowBattery(t *testing.T) {
	t.Log("Get the property OnLowBattery of object Upower ===> ",
		GetUpower("/org/freedesktop/UPower").GetOnLowBattery())
}

func TestUpowerPropertyLidIsClosed(t *testing.T) {
	t.Log("Get the property LidIsClosed of object Upower ===> ",
		GetUpower("/org/freedesktop/UPower").GetLidIsClosed())
}

func TestUpowerPropertyLidIsPresent(t *testing.T) {
	t.Log("Get the property LidIsPresent of object Upower ===> ",
		GetUpower("/org/freedesktop/UPower").GetLidIsPresent())
}

func TestUpowerPropertyLidForceSleep(t *testing.T) {
	t.Log("Get the property LidForceSleep of object Upower ===> ",
		GetUpower("/org/freedesktop/UPower").GetLidForceSleep())
}

func TestUpowerPropertyIsDocked(t *testing.T) {
	t.Log("Get the property IsDocked of object Upower ===> ",
		GetUpower("/org/freedesktop/UPower").GetIsDocked())
}

func TestUpowerSignalDeviceAdded(t *testing.T) {
}

func TestUpowerSignalDeviceRemoved(t *testing.T) {
}

func TestUpowerSignalDeviceChanged(t *testing.T) {
}

func TestUpowerSignalChanged(t *testing.T) {
}

func TestUpowerSignalSleeping(t *testing.T) {
}

func TestUpowerSignalNotifySleep(t *testing.T) {
}

func TestUpowerSignalResuming(t *testing.T) {
}

func TestUpowerSignalNotifyResume(t *testing.T) {
}
