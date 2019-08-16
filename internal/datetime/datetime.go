package datetime

import (
	"fmt"
	"time"

	"github.com/afreakk/i3statusbear/internal/config"
	"github.com/afreakk/i3statusbear/internal/protocol"
	"github.com/afreakk/i3statusbear/internal/util"
)

func Datetime(output *protocol.Output, module config.Module) func() {
	formatDateTimeMsg := func() string {
		return fmt.Sprintf(module.Sprintf, time.Now().Format(module.DateTimeFormat))
	}
	dateTimeMsg := &protocol.Message{
		FullText: formatDateTimeMsg(),
	}
	output.Messages = append(output.Messages, dateTimeMsg)
	util.ApplyModuleConfigToMessage(module, dateTimeMsg)
	return func() {
		dateTimeMsg.FullText = formatDateTimeMsg()
		output.PrintMsgs()
	}
}
