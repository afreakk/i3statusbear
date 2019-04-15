package readfile

import (
	"fmt"
	"io/ioutil"

	Config "github.com/afreakk/i3statusbear/internal/config"
	Protocol "github.com/afreakk/i3statusbear/internal/protocol"
	Util "github.com/afreakk/i3statusbear/internal/util"
)

func Readfile(output *Protocol.Output, module Config.Module) func() {
	formatString := func() string {
		data, _ := ioutil.ReadFile(module.FilePath)
		return fmt.Sprintf(module.Sprintf, string(data))
	}
	fileMsg := &Protocol.Message{
		FullText: formatString(),
	}
	output.Messages = append(output.Messages, fileMsg)
	Util.ApplyModuleConfigToMessage(module, fileMsg)
	return func() {
		fileMsg.FullText = formatString()
		output.PrintMsgs()
	}
}
