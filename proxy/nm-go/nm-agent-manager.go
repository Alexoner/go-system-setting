/*This file is auto generate by dlib/dbus/proxyer. Don't edit it*/

package nm

import "dlib/dbus"

type AgentManager struct {
	core *dbus.Object
}

func (obj AgentManager) Register(identifier string) {
	obj.core.Call("org.freedesktop.NetworkManager.AgentManager.Register", 0, identifier).Store()
	return
}

func (obj AgentManager) Unregister() {
	obj.core.Call("org.freedesktop.NetworkManager.AgentManager.Unregister", 0).Store()
	return
}

func GetAgentManager(path string) *AgentManager {
	return &AgentManager{getBus().Object("org.freedesktop.NetworkManager", dbus.ObjectPath(path))}
}
