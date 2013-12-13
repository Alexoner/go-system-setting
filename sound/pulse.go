// +build ignore

package main

// #cgo amd64 386 CFLAGS: -g -Wall
// #cgo LDFLAGS: -L. -ldde-pulse -lpulse -lc
// #include "dde-pulse.h"
import "C"

import (
	"dlib/dbus"
	"fmt"
	"os"
	"strconv"
)

type Audio struct {
	NumDevics int32
	pa        *C.pa
	//Change func(int32)
}

type Sink struct {
	Index          int32
	Name           string
	Description    string
	Driver         string
	Mute           string
	N_volume_steps int32
	Card           int32
	N_ports        int32
	N_formats      int32
	Cvolume        Volume
}

//Only capitalized first character in Capitalized structure can be exposed

type Source struct {
	Index          int32
	Name           string
	Description    string
	Driver         string
	Mute           int32
	N_volume_steps int32
	Card           int32
	C_ports        int32
	N_formates     int32
	Cvolume        Volume
}

type Sink_input struct {
	Index           int32
	Name            string
	Owner_module    int32
	Client          int32
	Sink            int32
	Cvolume         Volume
	Resample_method string
	Driver          string
	Mute            int32
	Corked          int32
	Has_volume      int32
	Volume_writable int32
}

type Source_output struct {
	Index           int32
	Name            string
	Owner_module    int32
	Client          int32
	Sink            int32
	Cvolume         Volume
	Resample_method string
	Driver          string
	Mute            int32
	Corked          int32
	Has_volume      int32
	Volume_writable int32
}

type Client struct {
	Index        int32
	Name         string
	Owner_module int32
	Driver       string
	//pa_proplist *proplist
	Prop map[string]string
}

type Card struct {
	Index        int32
	Name         string
	Owner_module int32
	Driver       string
	N_profiles   int32
	//pa_card_profile_info* profiles
	//pa_card_profile_info*  active_profile
	//pa_proplist *proplist
	n_ports int32
}

type Module struct {
}

type Volume struct {
	Channels uint32
	Values   [2]uint32 `access:"read"`
}

func NewAudio() (*Audio, error) {
	audio := Audio{}
	audio.pa = C.pa_new()
	if audio.pa == nil {
		fmt.Fprintln(os.Stderr,
			"unable to connect to the pulseaudio server,exiting\n")
		os.Exit(-1)
	}
	fmt.Println("module started\n")
	return &audio, nil
}

func (o *Audio) GetDBusInfo() dbus.DBusInfo {
	return dbus.DBusInfo{
		"com.deepin.daemon.Audio",
		"/com/deepin/daemon/Audio",
		"com.deepin.daemon.Audio",
	}
}

func (audio *Audio) GetServerInfo() string {
	C.pa_get_server_info(audio.pa)
	return C.GoString(audio.pa.server_info.host_name)
}

func (audio *Audio) GetSinks() []*Sink {
	sinks := make([]*Sink, 2)
	sinks[0] = &Sink{}
	sinks[0].Index = 2
	sinks[1] = &Sink{}
	sinks[1].Index = 3

	/*data, err := p.PulseCtl("list-sinks")
	if err != nil {
		println(err.Error())
		return nil
	}*/

	return sinks
}

func (audio *Audio) GetSources() *Source {
	return &Source{}
}

func (audio *Audio) GetSinkInputs() *Sink_input {
	return &Sink_input{}
}

func (audio *Audio) GetSourceOutput() *Source_output {
	return &Source_output{}
}

func (audio *Audio) GetClients() [2]*Client {
	var clients = [2]*Client{}
	clients[0] = new(Client)
	clients[0].Index = 0
	clients[1] = new(Client)
	clients[1].Index = 1
	return clients
}

func (audio *Audio) GetVolume() *Volume {
	return &Volume{}
}

func (o *Audio) IsRunning() uint32 {
	a := 3999
	return uint32(a)
}

func (sink *Sink) GetDBusInfo() dbus.DBusInfo {
	return dbus.DBusInfo{
		"com.deepin.daemon.Audio",
		"/com/deepin/daemon/Audio/Output" + strconv.FormatInt(int64(sink.Index), 10),
		"com.deepin.daemon.Audio.Output",
	}
}

func (sink *Sink) GetSinkVolume() Volume {
	return Volume{}
}

func (sink *Sink) SetSinkVolume() Volume {
	return Volume{}
}

func (sink *Sink) SetSinkMute(mute int32) int32 {
	return 1
}

func (source *Source) GetDBusInfo() dbus.DBusInfo {
	return dbus.DBusInfo{
		"com.deepin.daemon.Audio",
		"/com/deepin/daemon/Audio/Input" + strconv.FormatInt(int64(source.Index), 10),
		"com.deepin.daemon.Audio.Input",
	}
}

func (source *Source) GetSourceVolume() Volume {
	return Volume{}
}

func (source *Source) SetSourceVolume(volume Volume) Volume {
	return Volume{}
}

func (source *Source) SetSourceMute(mute int32) int32 {
	return 1
}

func (sink_input *Sink_input) GetDBusInfo() dbus.DBusInfo {
	return dbus.DBusInfo{
		"com.deepin.daemon.Audio",
		"/com/deepin/daemon/Audio/Application" + strconv.FormatInt(int64(sink_input.Index), 10),
		"com.deepin.daemon.Audio.Application",
	}
}

func (sink_input *Sink_input) GetSink_input_volume() Volume {
	return Volume{}
}

func (sink_input *Sink_input) SetSink_input_volume(volume Volume) int32 {
	return 1
}

func (sink_input *Sink_input) SetSink_input_mute(mute int32) int32 {
	return 1
}

func (source_output *Source_output) GetDBusInfo() dbus.DBusInfo {
	return dbus.DBusInfo{
		"com.deepin.daemon.Audio",
		"/com/deepin/daemon/Audio/Application" + strconv.FormatInt(int64(source_output.Index), 10),
		"com.deepin.daemon.Audio.Application",
	}
}

func (source_output *Source_output) GetSource_ouput_volume() Volume {
	return Volume{}
}

func (source_output *Source_output) SetSource_output_volume(volume Volume) Volume {
	return Volume{}
}

func (source_output *Source_output) SetSource_output_mute(mute int32) int32 {
	return 1
}

func (client *Client) GetDBusInfo() dbus.DBusInfo {
	return dbus.DBusInfo{
		"com.deepin.daemon.Audio",
		"/com/deepin/daemon/Audio/Client" + strconv.FormatInt(int64(client.Index), 10),
		"com.deepin.daemon.Audio.Client",
	}
}

func (card *Card) GetDBusInfo() dbus.DBusInfo {
	return dbus.DBusInfo{
		"com.deepin.daemon.Card",
		"/com/deepin/daemon/Audio/Card" + strconv.FormatInt(int64(card.Index), 10),
		"com.deepin.daemon.Card",
	}
}

func main() {
	pulse, err := NewAudio()
	if err != nil {
		panic(err)
		os.Exit(-1)
	}
	dbus.InstallOnSession(pulse)
	select {}
}
