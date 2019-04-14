package readfile

import (
	"fmt"
	"io/ioutil"

	Config "github.com/afreakk/i3statusbear/internal/config"
	Protocol "github.com/afreakk/i3statusbear/internal/protocol"
	Util "github.com/afreakk/i3statusbear/internal/util"
	"github.com/robfig/cron"
)

func Readfile(output *Protocol.Output, module Config.Module) {
	c := cron.New()
	formatString := func() string {
		data, _ := ioutil.ReadFile(module.FilePath)
		return fmt.Sprintf(module.Sprintf, string(data))
	}
	fileMsg := &Protocol.Message{
		FullText: formatString(),
	}
	Util.ApplyModuleConfigToMessage(module, fileMsg)
	c.AddFunc(module.Cron, func() {
		fileMsg.FullText = formatString()
		output.PrintMsgs()
	})
	c.Start()
	output.Messages = append(output.Messages, fileMsg)
}
