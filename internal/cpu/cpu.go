package cpu

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/afreakk/i3statusbear/internal/config"
	"github.com/afreakk/i3statusbear/internal/protocol"
	"github.com/afreakk/i3statusbear/internal/util"
)

func getCPUSample() (idle, total int64, err error) {
	f, err := os.Open("/proc/stat")
	if err != nil {
		return 0, 0, err
	}
	fScanner := bufio.NewScanner(f)
	var line string
	for fScanner.Scan() {
		line = fScanner.Text()
		if strings.HasPrefix(line, "cpu") {
			fields := strings.Fields(line)
			numFields := len(fields)
			for i := 1; i < numFields; i++ {
				val, parseErr := strconv.ParseInt(fields[i], 10, 64)
				if parseErr != nil {
					return 0, 0, err
				}
				total += val // tally up all the numbers to get total ticks
				if i == 4 {  // idle is the 5th field in the cpu line
					idle = val
				}
			}
			break
		}
	}
	f.Close()
	return
}

func Cpu(output *protocol.Output, module config.Module) func() {
	var lastIdle, lastTotal int64
	formatString := func() string {
		idle, total, err := getCPUSample()
		if err != nil {
			return err.Error()
		}
		idleTicks := idle - lastIdle
		totalTicks := total - lastTotal
		lastTotal = total
		lastIdle = idle
		return util.RenderBar(module, totalTicks-idleTicks, totalTicks)
	}
	cpuMsg := &protocol.Message{
		FullText: formatString(),
	}
	output.Messages = append(output.Messages, cpuMsg)
	util.ApplyModuleConfigToMessage(module, cpuMsg)
	var lastFullText string
	return func() {
		cpuMsg.FullText = formatString()
		if lastFullText != cpuMsg.FullText {
			output.PrintMsgs()
			lastFullText = cpuMsg.FullText
		}
	}
}
