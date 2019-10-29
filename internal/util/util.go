package util

import (
	"fmt"
	"strings"

	"github.com/afreakk/i3statusbear/internal/config"
	"github.com/afreakk/i3statusbear/internal/protocol"
)

func ApplyModuleConfigToMessage(module config.Module, message *protocol.Message) {
	if module.Align != "" {
		message.Align = module.Align
	}
	if module.Background != "" {
		message.Background = module.Background
	}
	if module.Border != "" {
		message.Border = module.Border
	}
	if module.Color != "" {
		message.Color = module.Color
	}
	if module.MinWidth != 0 {
		message.MinWidth = module.MinWidth
	}
	if module.Separator != false {
		message.Separator = module.Separator
	}
	if module.SeparatorWidth != 0 {
		message.SeparatorWidth = module.SeparatorWidth
	}
	if module.Urgent != false {
		message.Urgent = module.Urgent
	}
	if module.Markup != "" {
		message.Markup = module.Markup
	}

	if module.Name != "" {
		message.Name = module.Name
	}
	if module.Instance != "" {
		message.Instance = module.Instance
	}
}
func RenderBar(module config.Module, filled int64, total int64) string {
	filledBarsCount := filled * module.BarWidth / total
	remainingCount := Max(module.BarWidth-filledBarsCount, 0)
	filledBars := strings.Repeat(module.BarFilled, int(filledBarsCount))
	emptyBars := strings.Repeat(module.BarEmpty, int(remainingCount))
	return fmt.Sprintf(module.Sprintf, filledBars, emptyBars)
}
func Max(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}
