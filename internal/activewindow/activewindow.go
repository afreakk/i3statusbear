package activewindow

import (
	"fmt"
	"html"
	"log"

	"github.com/afreakk/go-i3"

	Config "github.com/afreakk/i3statusbear/internal/config"
	Protocol "github.com/afreakk/i3statusbear/internal/protocol"
	Util "github.com/afreakk/i3statusbear/internal/util"
)

func subscribeToI3Event(event i3.EventType, cb func()) {
	recv := i3.Subscribe(event)
	for recv.Next() {
		cb()
	}
	log.Fatal(recv.Close())
}

func ActiveWindow(output *Protocol.Output, module Config.Module) {
	formatString := func() string {
		tree, err := i3.GetTree()
		if err != nil {
			panic(err)
		}
		return fmt.Sprintf(
			module.Sprintf,
			// html.EscapeString should be replaced with somthing like
			// https://webreflection.github.io/gjs-documentation/GLib-2.0/GLib.markup_escape_text.html
			// but whatever, htmi escape works for now
			html.EscapeString(
				tree.Root.FindFocused(func(node *i3.Node) bool {
					return node.Focused == true
				}).Name))
	}
	wndMsg := &Protocol.Message{
		FullText: formatString(),
	}
	Util.ApplyModuleConfigToMessage(module, wndMsg)
	output.Messages = append(output.Messages, wndMsg)
	go subscribeToI3Event(i3.WindowEventType, func() {
		wndMsg.FullText = formatString()
		output.PrintMsgs()
	})
}
