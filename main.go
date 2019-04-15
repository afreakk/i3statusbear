package main

import (
	"os"

	Config "github.com/afreakk/i3statusbear/internal/config"
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
		}
	}
	output.PrintMsgs()
	c.Start()
	//hacky way of blocking forever..
	select {}
}
