package protocol

import (
	"encoding/json"
	"fmt"
	"os"
	"syscall"
)

type Output struct {
	encodeToStdout *json.Encoder
	Messages       []*Message
}

func (o *Output) Init() {
	// init stdout writer
	o.encodeToStdout = json.NewEncoder(os.Stdout)

	// then we need to specify protocol we use
	protocol := Protocol{
		Version: 1,
		// dont know about these params, found it somewhere, and it works..
		StopSignal:  int(syscall.Signal(10)),
		ContSignal:  int(syscall.Signal(12)),
		ClickEvents: false,
	}
	// and send it to Stdout
	o.encodeToStdout.Encode(protocol)
	// start our array of arrays.. (we never end it though)
	// kindof hacky, is there a better way ?
	fmt.Print("[")

	o.Messages = []*Message{}
}

func (o Output) PrintMsgs() {
	o.encodeToStdout.Encode(o.Messages)
	// And then separator between messages in our infinite array that never ends
	fmt.Print(",")
}

// === START: helper types for decode and encode to i3bar protocol ===
type Protocol struct {
	Version     int  `json:"version"`
	StopSignal  int  `json:"stop_signal"`
	ContSignal  int  `json:"cont_signal"`
	ClickEvents bool `json:"click_events"`
}

type Click struct {
	Name       string   `json:"name"`
	Instance   string   `json:"instance"`
	Button     int      `json:"button"`
	X          int      `json:"x"`
	Y          int      `json:"y"`
	Relative_x int      `json:relative_x`
	Relative_y int      `json:relative_y`
	Width      int      `json:width`
	Height     int      `json:height`
	Modifiers  []string `json:modifiers`
}

type Message struct {
	FullText       string `json:"full_text"`
	ShortText      string `json:"short_text,omitempty"`
	Color          string `json:"color,omitempty"`
	Background     string `json:"background,omitempty"`
	Border         string `json:"border,omitempty"`
	MinWidth       int    `json:"min_width,omitempty"`
	Align          string `json:"align,omitempty"`
	Name           string `json:"name,omitempty"`
	Instance       string `json:"instance,omitempty"`
	Urgent         bool   `json:"urgent,omitempty"`
	Separator      bool   `json:"separator,omitempty"`
	SeparatorWidth int    `json:"separator_block_width,omitempty"`
	Markup         string `json:"markup,omitempty"`
}

func HandleInput() {
	decodeStdin := json.NewDecoder(os.Stdin)
	// Read openbracket
	_, err := decodeStdin.Token()
	if err != nil {
		panic(err)
	}
	for decodeStdin.More() {
		click := &Click{}
		err := decodeStdin.Decode(click)
		if err != nil {
			panic(err)
		}
	}
}
