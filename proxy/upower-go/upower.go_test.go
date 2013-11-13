/*This file is auto generate by dlib/dbus/proxyer. Don't edit it*/
package upower

import "testing"

func TestupowerMethodEnumerateDevices(t *testing.T) {

}

func TestupowerMethodAboutToSleep(t *testing.T) {

}

func TestupowerMethodSuspend(t *testing.T) {

}

func TestupowerMethodSuspendAllowed(t *testing.T) {

}

func TestupowerMethodHibernate(t *testing.T) {

}

func TestupowerMethodHibernateAllowed(t *testing.T) {

}

func TestupowerPropertyDaemonVersion(t *testing.T) {
	t.Log("Get the property DaemonVersion of object upower ===> ",
		Getupower("/org/freedesktop/UPower").GetDaemonVersion())
}

func TestupowerPropertyCanSuspend(t *testing.T) {
	t.Log("Get the property CanSuspend of object upower ===> ",
		Getupower("/org/freedesktop/UPower").GetCanSuspend())
}

func TestupowerPropertyCanHibernate(t *testing.T) {
	t.Log("Get the property CanHibernate of object upower ===> ",
		Getupower("/org/freedesktop/UPower").GetCanHibernate())
}

func TestupowerPropertyOnBattery(t *testing.T) {
	t.Log("Get the property OnBattery of object upower ===> ",
		Getupower("/org/freedesktop/UPower").GetOnBattery())
}

func TestupowerPropertyOnLowBattery(t *testing.T) {
	t.Log("Get the property OnLowBattery of object upower ===> ",
		Getupower("/org/freedesktop/UPower").GetOnLowBattery())
}

func TestupowerPropertyLidIsClosed(t *testing.T) {
	t.Log("Get the property LidIsClosed of object upower ===> ",
		Getupower("/org/freedesktop/UPower").GetLidIsClosed())
}

func TestupowerPropertyLidIsPresent(t *testing.T) {
	t.Log("Get the property LidIsPresent of object upower ===> ",
		Getupower("/org/freedesktop/UPower").GetLidIsPresent())
}

func TestupowerPropertyLidForceSleep(t *testing.T) {
	t.Log("Get the property LidForceSleep of object upower ===> ",
		Getupower("/org/freedesktop/UPower").GetLidForceSleep())
}

func TestupowerPropertyIsDocked(t *testing.T) {
	t.Log("Get the property IsDocked of object upower ===> ",
		Getupower("/org/freedesktop/UPower").GetIsDocked())
}

func TestupowerSignalDeviceAdded(t *testing.T) {
}

func TestupowerSignalDeviceRemoved(t *testing.T) {
}

func TestupowerSignalDeviceChanged(t *testing.T) {
}

func TestupowerSignalChanged(t *testing.T) {
}

func TestupowerSignalSleeping(t *testing.T) {
}

func TestupowerSignalNotifySleep(t *testing.T) {
}

func TestupowerSignalResuming(t *testing.T) {
}

func TestupowerSignalNotifyResume(t *testing.T) {
}
