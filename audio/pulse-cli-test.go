//+build ignore

package main

import (
	"../system-settings"
	//"fmt"
)

func main() {
	p := pulse.Pulse{}
	p.PulseInit()
	p.PulseConnect()
	//pulse.PulseCliShell(p.Path)
	line, _ := p.PulseCtl("list-modules")
	println(line)
	line, _ = p.PulseCtl("list-sinks")
	println(line)
	line, _ = p.PulseCtl("list-sources")
	println(line)
}
