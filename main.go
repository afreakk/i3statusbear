package main

import (
	"fmt"
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

func errorExit(str string, code int) {
	fmt.Fprintln(os.Stderr, str)
	os.Exit(code)
}

func main() {
	osArgsLen := len(os.Args)

	if osArgsLen < 2 {
		errorExit("Please provide config as argument", 1)
	}
	configFilePath := os.Args[1]

	cfg := config.GetConfigFromPath(configFilePath)
	if osArgsLen > 2 {
		cfg.WMClient = os.Args[2]
	}

	go protocol.HandleInput()

	if cfg.WMClient == "sway" {
		i3.WMClient = i3.WMTypeSway
	}

	output := protocol.Output{}
	output.Init(cfg)

	c := cron.New()
	for _, module := range cfg.Modules {
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
