package cpu

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	Config "github.com/afreakk/i3statusbear/internal/config"
	Protocol "github.com/afreakk/i3statusbear/internal/protocol"
	Util "github.com/afreakk/i3statusbear/internal/util"
)

func getCPUSample() (idle, total int64) {
	f, _ := os.Open("/proc/stat")
	fScanner := bufio.NewScanner(f)
	var line string
	for fScanner.Scan() {
		line = fScanner.Text()
		if strings.HasPrefix(line, "cpu") {
			fields := strings.Fields(line)
			numFields := len(fields)
			for i := 1; i < numFields; i++ {
				val, _ := strconv.ParseInt(fields[i], 10, 64)
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

func Cpu(output *Protocol.Output, module Config.Module) func() {
	var lastIdle, lastTotal int64
	formatString := func() string {
		idle, total := getCPUSample()
		idleTicks := idle - lastIdle
		totalTicks := total - lastTotal
		lastTotal = total
		lastIdle = idle
		return Util.RenderBar(module, totalTicks-idleTicks, totalTicks)
	}
	cpuMsg := &Protocol.Message{
		FullText: formatString(),
	}
	output.Messages = append(output.Messages, cpuMsg)
	Util.ApplyModuleConfigToMessage(module, cpuMsg)
	return func() {
		cpuMsg.FullText = formatString()
		output.PrintMsgs()
	}
}
