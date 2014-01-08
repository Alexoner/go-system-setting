//+build ignore

package main

import (
	"dlib"
	"fmt"
)

func TestGSettings() {
	go dlib.StartLoop() //gtk_main()

	s := dlib.NewSettings("org.gnome.settings-daemon.plugins.power")
	fmt.Println(s.ListKeys())
	v := s.GetBoolean("active")
	defer func() {
		fmt.Println("Power.active: %v", v)
	}()
	changed_times := 0
	//s.Connect("changed::active-mini-mode", func(s *Settings, name string) {
	//	changed_times++
	//})

	//s.SetBoolean("active-mini-mode", true)
	//if s.GetBoolean("active-mini-mode") != true {
	//	t.Error("SetBoolean failed")
	//}
	//s.SetBoolean("active-mini-mode", false)
	//if s.GetBoolean("active-mini-mode") != false {
	//	t.Error("SetBoolean failed")
	//}

	if changed_times != 2 {
		fmt.Println("changed_times", changed_times)
		//t.Error("ChangedError")
	}
}

func main() {
	TestGSettings()
}
