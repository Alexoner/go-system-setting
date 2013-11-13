/*This file is auto generate by dlib/dbus/proxyer. Don't edit it*/

package nm

import "dlib/dbus"

type SecretAgent struct {
	core *dbus.Object
}

func (obj SecretAgent) GetSecrets(connection map[string]map[string]dbus.Variant, connection_path dbus.ObjectPath, setting_name string, hints []string, flags uint32) (secrets map[string]map[string]dbus.Variant) {
	obj.core.Call("org.freedesktop.NetworkManager.SecretAgent.GetSecrets", 0, connection, connection_path, setting_name, hints, flags).Store(&secrets)
	return
}

func (obj SecretAgent) CancelGetSecrets(connection_path dbus.ObjectPath, setting_name string) {
	obj.core.Call("org.freedesktop.NetworkManager.SecretAgent.CancelGetSecrets", 0, connection_path, setting_name).Store()
	return
}

func (obj SecretAgent) SaveSecrets(connection map[string]map[string]dbus.Variant, connection_path dbus.ObjectPath) {
	obj.core.Call("org.freedesktop.NetworkManager.SecretAgent.SaveSecrets", 0, connection, connection_path).Store()
	return
}

func (obj SecretAgent) DeleteSecrets(connection map[string]map[string]dbus.Variant, connection_path dbus.ObjectPath) {
	obj.core.Call("org.freedesktop.NetworkManager.SecretAgent.DeleteSecrets", 0, connection, connection_path).Store()
	return
}

func GetSecretAgent(path string) *SecretAgent {
	return &SecretAgent{getBus().Object("org.freedesktop.NetworkManager", dbus.ObjectPath(path))}
}
