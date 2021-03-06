package protocol

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"syscall"
	"time"

	"github.com/afreakk/i3statusbear/internal/config"
)

type Output struct {
	encodeToStdout       *json.Encoder
	Messages             []*Message
	renderTimer          *time.Timer
	renderTimerIsRunning bool
	renderInterval       time.Duration
	mux                  sync.Mutex
	jsonSeparator        []byte
}

func (o *Output) Init(cfg config.Config) (err error) {
	o.jsonSeparator = []byte(",")
	// init stdout writer
	o.encodeToStdout = json.NewEncoder(os.Stdout)
	o.encodeToStdout.SetEscapeHTML(false)

	// then we need to specify protocol we use
	// and send it to Stdout
	err = o.encodeToStdout.Encode(Protocol{
		Version: 1,
		// dont know about these params, found it somewhere, and it works..
		StopSignal:  int(syscall.Signal(10)),
		ContSignal:  int(syscall.Signal(12)),
		ClickEvents: false,
	})
	if err != nil {
		return
	}
	// start our array of arrays.. (we never end it though)
	// kindof hacky, is there a better way ?
	_, err = os.Stdout.Write([]byte("["))
	if err != nil {
		return
	}

	o.Messages = []*Message{}

	if o.renderInterval, err = time.ParseDuration(cfg.MinimumRenderInterval); err != nil {
		return
	}
	o.renderTimer = time.AfterFunc(o.renderInterval, o.ActuallyPrintMsgs)
	return
}

func (o *Output) PrintMsgs() {
	o.mux.Lock()
	if !o.renderTimerIsRunning {
		o.renderTimer.Reset(o.renderInterval)
		o.renderTimerIsRunning = true
	}
	o.mux.Unlock()
}

func printToStderrIfErr(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func (o *Output) ActuallyPrintMsgs() {
	o.renderTimerIsRunning = false
	printToStderrIfErr(o.encodeToStdout.Encode(o.Messages))
	// And then separator between messages in our infinite array that never ends
	_, err := os.Stdout.Write(o.jsonSeparator)
	printToStderrIfErr(err)
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
	Relative_x int      `json:"relative_x"`
	Relative_y int      `json:"relative_y"`
	Width      int      `json:"width"`
	Height     int      `json:"height"`
	Modifiers  []string `json:"modifiers"`
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
	/*
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
	*/
}
