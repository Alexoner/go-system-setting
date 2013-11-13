/*This file is auto generate by dlib/dbus/proxyer. Don't edit it*/
package nm

import "testing"

func TestManagerMethodGetDevices(t *testing.T) {

}

func TestManagerMethodGetDeviceByIpIface(t *testing.T) {

}

func TestManagerMethodActivateConnection(t *testing.T) {

}

func TestManagerMethodAddAndActivateConnection(t *testing.T) {

}

func TestManagerMethodDeactivateConnection(t *testing.T) {

}

func TestManagerMethodSleep(t *testing.T) {

}

func TestManagerMethodEnable(t *testing.T) {

}

func TestManagerMethodGetPermissions(t *testing.T) {

}

func TestManagerMethodSetLogging(t *testing.T) {

}

func TestManagerMethodGetLogging(t *testing.T) {

}

func TestManagerMethodstate(t *testing.T) {

}

func TestManagerPropertyNetworkingEnabled(t *testing.T) {
	t.Log("Get the property NetworkingEnabled of object Manager ===> ",
		GetManager("/org/freedesktop/NetworkManager").GetNetworkingEnabled())
}

func TestManagerPropertyWirelessEnabled(t *testing.T) {
	t.Log("Get the property WirelessEnabled of object Manager ===> ",
		GetManager("/org/freedesktop/NetworkManager").GetWirelessEnabled())
}

func TestManagerPropertyWirelessHardwareEnabled(t *testing.T) {
	t.Log("Get the property WirelessHardwareEnabled of object Manager ===> ",
		GetManager("/org/freedesktop/NetworkManager").GetWirelessHardwareEnabled())
}

func TestManagerPropertyWwanEnabled(t *testing.T) {
	t.Log("Get the property WwanEnabled of object Manager ===> ",
		GetManager("/org/freedesktop/NetworkManager").GetWwanEnabled())
}

func TestManagerPropertyWwanHardwareEnabled(t *testing.T) {
	t.Log("Get the property WwanHardwareEnabled of object Manager ===> ",
		GetManager("/org/freedesktop/NetworkManager").GetWwanHardwareEnabled())
}

func TestManagerPropertyWimaxEnabled(t *testing.T) {
	t.Log("Get the property WimaxEnabled of object Manager ===> ",
		GetManager("/org/freedesktop/NetworkManager").GetWimaxEnabled())
}

func TestManagerPropertyWimaxHardwareEnabled(t *testing.T) {
	t.Log("Get the property WimaxHardwareEnabled of object Manager ===> ",
		GetManager("/org/freedesktop/NetworkManager").GetWimaxHardwareEnabled())
}

func TestManagerPropertyActiveConnections(t *testing.T) {
	t.Log("Get the property ActiveConnections of object Manager ===> ",
		GetManager("/org/freedesktop/NetworkManager").GetActiveConnections())
}

func TestManagerPropertyVersion(t *testing.T) {
	t.Log("Get the property Version of object Manager ===> ",
		GetManager("/org/freedesktop/NetworkManager").GetVersion())
}

func TestManagerPropertyState(t *testing.T) {
	t.Log("Get the property State of object Manager ===> ",
		GetManager("/org/freedesktop/NetworkManager").GetState())
}

func TestManagerSignalCheckPermissions(t *testing.T) {
}

func TestManagerSignalStateChanged(t *testing.T) {
}

func TestManagerSignalPropertiesChanged(t *testing.T) {
}

func TestManagerSignalDeviceAdded(t *testing.T) {
}

func TestManagerSignalDeviceRemoved(t *testing.T) {
}
