package memory

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	Config "github.com/afreakk/i3statusbear/internal/config"
	Protocol "github.com/afreakk/i3statusbear/internal/protocol"
	Util "github.com/afreakk/i3statusbear/internal/util"
)

func Memory(output *Protocol.Output, module Config.Module) func() {
	formatString := func() string {
		f, _ := os.Open("/proc/meminfo")
		fScanner := bufio.NewScanner(f)
		var line string
		var memtotal int64
		var memavail int64
		for fScanner.Scan() {
			line = fScanner.Text()
			if strings.HasPrefix(line, "MemTotal") {
				memtotal, _ = strconv.ParseInt(strings.Fields(line)[1], 10, 32)
			} else if strings.HasPrefix(line, "MemAvailable") {
				memavail, _ = strconv.ParseInt(strings.Fields(line)[1], 10, 32)
			}
			if memtotal != 0 && memavail != 0 {
				break
			}
		}
		f.Close()
		return Util.RenderBar(module, memtotal-memavail, memtotal)
	}
	memMsg := &Protocol.Message{
		FullText: formatString(),
	}
	output.Messages = append(output.Messages, memMsg)
	Util.ApplyModuleConfigToMessage(module, memMsg)
	var lastFullText string
	return func() {
		memMsg.FullText = formatString()
		if lastFullText != memMsg.FullText {
			output.PrintMsgs()
		}
		lastFullText = memMsg.FullText
	}
}
