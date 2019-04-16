package main

import (
	"os"

	i3 "github.com/afreakk/go-i3"
	ActiveWindow "github.com/afreakk/i3statusbear/internal/activewindow"
	Command "github.com/afreakk/i3statusbear/internal/command"
	Config "github.com/afreakk/i3statusbear/internal/config"
	Cpu "github.com/afreakk/i3statusbear/internal/cpu"
	Datetime "github.com/afreakk/i3statusbear/internal/datetime"
	Memory "github.com/afreakk/i3statusbear/internal/memory"
	Protocol "github.com/afreakk/i3statusbear/internal/protocol"
	Pulseaudio "github.com/afreakk/i3statusbear/internal/pulseaudio"
	Readfile "github.com/afreakk/i3statusbear/internal/readfile"
	"github.com/robfig/cron"
)

func main() {
	configFilePath := os.Args[1]
	config := Config.GetConfigFromPath(configFilePath)

	go Protocol.HandleInput()

	if config.WMClient == "sway" {
		i3.WMClient = i3.WMTypeSway
	}

	output := Protocol.Output{}
	output.Init(config)

	c := cron.New()
	for _, module := range config.Modules {
		switch module.Name {
		case "datetime":
			c.AddFunc(module.Cron, Datetime.Datetime(&output, module))
		case "pulseaudio":
			Pulseaudio.Pulseaudio(&output, module)
		case "readfile":
			c.AddFunc(module.Cron, Readfile.Readfile(&output, module))
		case "memory":
			c.AddFunc(module.Cron, Memory.Memory(&output, module))
		case "cpu":
			c.AddFunc(module.Cron, Cpu.Cpu(&output, module))
		case "command":
			c.AddFunc(module.Cron, Command.Command(&output, module))
		case "activewindow":
			ActiveWindow.ActiveWindow(&output, module)
		}
	}
	output.PrintMsgs()
	c.Start()
	//hacky way of blocking forever..
	select {}
}
