package activewindow

import (
	"fmt"
	"log"

	"github.com/afreakk/go-i3"

	Config "github.com/afreakk/i3statusbear/internal/config"
	Protocol "github.com/afreakk/i3statusbear/internal/protocol"
	Util "github.com/afreakk/i3statusbear/internal/util"
)

func subscribe(cb func()) {
	recv := i3.Subscribe(i3.WindowEventType)
	for recv.Next() {
		ev := recv.Event().(*i3.WindowEvent)
		cb()
		log.Printf("change: %s", ev.Change)
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
			tree.Root.FindFocused(func(node *i3.Node) bool {
				return node.Focused == true
			}).Name)
	}
	wndMsg := &Protocol.Message{
		FullText: formatString(),
	}
	Util.ApplyModuleConfigToMessage(module, wndMsg)
	output.Messages = append(output.Messages, wndMsg)
	go subscribe(func() {
		wndMsg.FullText = formatString()
		output.PrintMsgs()
	})
}
