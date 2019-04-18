package command

import (
	"fmt"
	"os/exec"

	Config "github.com/afreakk/i3statusbear/internal/config"
	Protocol "github.com/afreakk/i3statusbear/internal/protocol"
	Util "github.com/afreakk/i3statusbear/internal/util"
)

func Command(output *Protocol.Output, module Config.Module) func() {
	formatString := func() string {
		cmd := exec.Command(module.CommandName, module.CommandArgs...)
		out, err := cmd.Output()
		if err != nil {
			return err.Error()
		}

		return fmt.Sprintf(module.Sprintf, string(out[:len(out)-1]))
	}
	cmdMsg := &Protocol.Message{
		FullText: formatString(),
	}
	output.Messages = append(output.Messages, cmdMsg)
	Util.ApplyModuleConfigToMessage(module, cmdMsg)
	var lastFullText string
	return func() {
		cmdMsg.FullText = formatString()
		if lastFullText != cmdMsg.FullText {
			output.PrintMsgs()
		}
		lastFullText = cmdMsg.FullText
	}
}
