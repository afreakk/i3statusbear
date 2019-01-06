package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"syscall"
	"time"

	"github.com/godbus/dbus"
	"github.com/robfig/cron"
	"github.com/sqp/pulseaudio"
)

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
	Position       int    `json:"-"`
	FullText       string `json:"full_text"`
	ShortText      string `json:"short_text,omitempty"`
	Color          string `json:"color,omitempty"`
	Background     string `json:"background,omitempty"`
	Border         string `json:"border,omitempty"`
	MinWidth       string `json:"min_width,omitempty"`
	Align          Align  `json:"align,omitempty"`
	Name           string `json:"name,omitempty"`
	Instance       string `json:"instance,omitempty"`
	Urgent         bool   `json:"urgent,omitempty"`
	Separator      bool   `json:"separator,omitempty"`
	SeparatorWidth int    `json:"separator_block_width,omitempty"`
}

type Align string

const (
	LEFT   Align = "left"
	RIGHT  Align = "right"
	CENTER Align = "center"
)

// === END: helper types for decode and encode to i3bar protocol ===

func handleInput(logger *log.Logger) {
	decodeStdin := json.NewDecoder(os.Stdin)
	// Read openbracket
	_, err := decodeStdin.Token()
	if err != nil {
		logger.Fatal(err)
	}
	for decodeStdin.More() {
		click := &Click{}
		err := decodeStdin.Decode(click)
		if err != nil {
			logger.Fatal(err)
		}
		logger.Println(click)
	}
}

func handleOutput(logger *log.Logger) {
	// init stdout writer
	encodeToStdout := json.NewEncoder(os.Stdout)

	// then we need to specify protocol we use
	protocol := Protocol{
		Version: 1,
		// dont know about these params, found it somewhere, and it works..
		StopSignal:  int(syscall.Signal(10)),
		ContSignal:  int(syscall.Signal(12)),
		ClickEvents: true,
	}
	// and send it to Stdout
	encodeToStdout.Encode(protocol)

	// start our array of arrays.. (we never end it though)
	// kindof hacky, is there a better way ?
	fmt.Print("[")

	messages := []*Message{}

	// render method.. gets called by modules when they need to redraw
	renderMsgs := func() {
		encodeToStdout.Encode(messages)
		fmt.Print(",")
	}

	// cron can be used by many modules
	c := cron.New()

	// === Start: DateTime setup ===
	formatDateTimeMsg := func() string {
		return time.Now().Format("15:04 1/Jan")
	}
	dateTimeMsg := &Message{
		FullText: formatDateTimeMsg(),
		Name:     "dateTimeMsg",
	}
	c.AddFunc("0 * * * * *", func() {
		dateTimeMsg.FullText = formatDateTimeMsg()
		renderMsgs()
	})
	// === End: DateTime setup ===

	// === Start: PulseAudio setup ===
	pulse, e := pulseaudio.New()
	if e != nil {
		logger.Panicln(e)
	}

	var pathToFallbackSink dbus.ObjectPath
	// Here we assume you are using fallbacksink, so we query that
	// altough later in pulse-callback we render whatever device you changed volume on :)
	pulse.Core().Get("FallbackSink", &pathToFallbackSink)
	volumes, e := pulse.Device(pathToFallbackSink).ListUint32("Volume")
	if e != nil {
		logger.Panicln(e)
	}
	baseVolume, e := pulse.Device(pathToFallbackSink).Uint32("BaseVolume")
	if e != nil {
		logger.Panicln(e)
	}
	formatPulseAudioText := func(volumes []uint32, baseVolume uint32) string {
		return fmt.Sprintf("vol %.2f", float32(volumes[0])/float32(baseVolume))
	}
	pulseAudioMsg := &Message{
		FullText: formatPulseAudioText(volumes, baseVolume),
		Name:     "pulseMsg",
	}
	updatePulseMsg := func(volumes []uint32, baseVolume uint32) {
		pulseAudioMsg.FullText = formatPulseAudioText(volumes, baseVolume)
		renderMsgs()
	}

	client := &Client{logger, pulse, updatePulseMsg}
	pulse.Register(client)

	go pulse.Listen()
	// === End: PulseAudio setup ===

	//ordering of modules
	messages = append(messages, pulseAudioMsg)
	messages = append(messages, dateTimeMsg)

	c.Start()
	renderMsgs()
}

func main() {
	// Init log etc
	logFile, _ := os.OpenFile("log/i3statusbear"+time.Now().Format(time.RFC3339)+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	logger := log.New(logFile, "logger: ", log.Lshortfile)

	logger.Println("starting..")
	go handleInput(logger)
	handleOutput(logger)
	//hacky way of blocking forever..
	select {}
}

// Pulse audio subscribe stuff
type Client struct {
	logger         *log.Logger
	pulse          *pulseaudio.Client
	updatePulseMsg func(volumes []uint32, baseVolume uint32)
}

func (cl *Client) DeviceVolumeUpdated(path dbus.ObjectPath, values []uint32) {
	cl.logger.Println("device volume", path, values)
	var baseVolume uint32
	cl.pulse.Device(path).Get("BaseVolume", &baseVolume)
	cl.updatePulseMsg(values, baseVolume)
}
