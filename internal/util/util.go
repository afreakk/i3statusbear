package util

import (
	Config "github.com/afreakk/i3statusbear/internal/config"
	Protocol "github.com/afreakk/i3statusbear/internal/protocol"
)

func ApplyModuleConfigToMessage(module Config.Module, message *Protocol.Message) {
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
