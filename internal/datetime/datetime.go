package datetime

import (
	"time"

	Config "github.com/afreakk/i3statusbear/internal/config"
	Protocol "github.com/afreakk/i3statusbear/internal/protocol"
	Util "github.com/afreakk/i3statusbear/internal/util"
	"github.com/robfig/cron"
)

func Datetime(output *Protocol.Output, module Config.Module) {

	c := cron.New()
	formatDateTimeMsg := func() string {
		return time.Now().Format(module.DateTimeFormat)
	}
	dateTimeMsg := &Protocol.Message{
		FullText: formatDateTimeMsg(),
	}
	Util.ApplyModuleConfigToMessage(module, dateTimeMsg)
	c.AddFunc("0 * * * * *", func() {
		dateTimeMsg.FullText = formatDateTimeMsg()
		output.PrintMsgs()
	})
	c.Start()
	output.Messages = append(output.Messages, dateTimeMsg)
}
