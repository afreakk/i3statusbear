package activewindow

import (
	"fmt"
	"html"

	"github.com/afreakk/go-i3"

	"github.com/afreakk/i3statusbear/internal/config"
	"github.com/afreakk/i3statusbear/internal/protocol"
	"github.com/afreakk/i3statusbear/internal/util"
)

func ActiveWindow(output *protocol.Output, module config.Module) {
	formatString := func() string {
		tree, err := i3.GetTree()
		if err != nil {
			return err.Error()
		}
		return fmt.Sprintf(
			module.Sprintf,
			// html.EscapeString should be replaced with somthing like
			// https://webreflection.github.io/gjs-documentation/GLib-2.0/GLib.markup_escape_text.html
			// but whatever, htmi escape works for now
			html.EscapeString(
				tree.Root.FindFocused(func(node *i3.Node) bool {
					return node.Focused
				}).Name))
	}
	wndMsg := &protocol.Message{
		FullText: formatString(),
	}
	util.ApplyModuleConfigToMessage(module, wndMsg)
	output.Messages = append(output.Messages, wndMsg)
	var lastFullText string
	go func() {
		recv := i3.Subscribe(i3.WindowEventType)
		for recv.Next() {
			wndMsg.FullText = formatString()
			if lastFullText != wndMsg.FullText {
				output.PrintMsgs()
				lastFullText = wndMsg.FullText
			}
		}
		wndMsg.FullText = recv.Close().Error()
		output.PrintMsgs()
	}()
}
