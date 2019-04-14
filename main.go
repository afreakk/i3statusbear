package main

import (
	"os"

	Config "github.com/afreakk/i3statusbear/internal/config"
	Datetime "github.com/afreakk/i3statusbear/internal/datetime"
	Protocol "github.com/afreakk/i3statusbear/internal/protocol"
	Pulseaudio "github.com/afreakk/i3statusbear/internal/pulseaudio"
	Readfile "github.com/afreakk/i3statusbear/internal/readfile"
)

func main() {
	configFilePath := os.Args[1]
	config := Config.GetConfigFromPath(configFilePath)

	go Protocol.HandleInput()

	output := Protocol.Output{}
	output.Init()
	for _, module := range config.Modules {
		switch module.Name {
		case "datetime":
			Datetime.Datetime(&output, module)
		case "pulseaudio":
			Pulseaudio.Pulseaudio(&output, module)
		case "readfile":
			Readfile.Readfile(&output, module)
		}
	}
	output.PrintMsgs()
	//hacky way of blocking forever..
	select {}
}
