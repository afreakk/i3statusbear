package datetime

import (
	"time"

	Config "github.com/afreakk/i3statusbear/internal/config"
	Protocol "github.com/afreakk/i3statusbear/internal/protocol"
	Util "github.com/afreakk/i3statusbear/internal/util"
)

func Datetime(output *Protocol.Output, module Config.Module) func() {
	formatDateTimeMsg := func() string {
		return time.Now().Format(module.DateTimeFormat)
	}
	dateTimeMsg := &Protocol.Message{
		FullText: formatDateTimeMsg(),
	}
	output.Messages = append(output.Messages, dateTimeMsg)
	Util.ApplyModuleConfigToMessage(module, dateTimeMsg)
	return func() {
		dateTimeMsg.FullText = formatDateTimeMsg()
		output.PrintMsgs()
	}
}
