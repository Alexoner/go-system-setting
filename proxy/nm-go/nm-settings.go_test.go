/*This file is auto generate by dlib/dbus/proxyer. Don't edit it*/
package nm

import "testing"

func TestSettingsMethodListConnections(t *testing.T) {

}

func TestSettingsMethodGetConnectionByUuid(t *testing.T) {

}

func TestSettingsMethodAddConnection(t *testing.T) {

}

func TestSettingsMethodSaveHostname(t *testing.T) {

}

func TestSettingsPropertyHostname(t *testing.T) {
	t.Log("Get the property Hostname of object Settings ===> ",
		GetSettings("/org/freedesktop/NetworkManager/Settings").GetHostname())
}

func TestSettingsPropertyCanModify(t *testing.T) {
	t.Log("Get the property CanModify of object Settings ===> ",
		GetSettings("/org/freedesktop/NetworkManager/Settings").GetCanModify())
}

func TestSettingsSignalPropertiesChanged(t *testing.T) {
}

func TestSettingsSignalNewConnection(t *testing.T) {
}
