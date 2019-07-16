package readfile

import (
	"fmt"
	"io/ioutil"

	"github.com/afreakk/i3statusbear/internal/config"
	"github.com/afreakk/i3statusbear/internal/protocol"
	"github.com/afreakk/i3statusbear/internal/util"
)

func Readfile(output *protocol.Output, module config.Module) func() {
	formatString := func() string {
		data, err := ioutil.ReadFile(module.FilePath)
		if err != nil {
			return err.Error()
		}
		return fmt.Sprintf(module.Sprintf, string(data[:len(data)-1]))
	}
	fileMsg := &protocol.Message{
		FullText: formatString(),
	}
	output.Messages = append(output.Messages, fileMsg)
	util.ApplyModuleConfigToMessage(module, fileMsg)
	var lastFullText string
	return func() {
		fileMsg.FullText = formatString()
		if lastFullText != fileMsg.FullText {
			output.PrintMsgs()
			lastFullText = fileMsg.FullText
		}
	}
}
