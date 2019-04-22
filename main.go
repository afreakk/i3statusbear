package main

import (
	"os"

	"github.com/afreakk/go-i3"
	"github.com/afreakk/i3statusbear/internal/activewindow"
	"github.com/afreakk/i3statusbear/internal/command"
	"github.com/afreakk/i3statusbear/internal/config"
	"github.com/afreakk/i3statusbear/internal/cpu"
	"github.com/afreakk/i3statusbear/internal/datetime"
	"github.com/afreakk/i3statusbear/internal/memory"
	"github.com/afreakk/i3statusbear/internal/protocol"
	"github.com/afreakk/i3statusbear/internal/pulseaudio"
	"github.com/afreakk/i3statusbear/internal/readfile"
	"github.com/robfig/cron"
)

func main() {
	configFilePath := os.Args[1]
	config := config.GetConfigFromPath(configFilePath)
	if len(os.Args) > 2 {
		config.WMClient = os.Args[2]
	}

	go protocol.HandleInput()

	if config.WMClient == "sway" {
		i3.WMClient = i3.WMTypeSway
	}

	output := protocol.Output{}
	output.Init(config)

	c := cron.New()
	for _, module := range config.Modules {
		switch module.Name {
		case "datetime":
			c.AddFunc(module.Cron, datetime.Datetime(&output, module))
		case "pulseaudio":
			pulseaudio.Pulseaudio(&output, module)
		case "readfile":
			c.AddFunc(module.Cron, readfile.Readfile(&output, module))
		case "memory":
			c.AddFunc(module.Cron, memory.Memory(&output, module))
		case "cpu":
			c.AddFunc(module.Cron, cpu.Cpu(&output, module))
		case "command":
			c.AddFunc(module.Cron, command.Command(&output, module))
		case "activewindow":
			activewindow.ActiveWindow(&output, module)
		}
	}
	output.PrintMsgs()
	c.Start()
	//hacky way of blocking forever..
	select {}
}
