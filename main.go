package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"syscall"
	"time"

	"github.com/godbus/dbus"
	"github.com/robfig/cron"
	"github.com/sqp/pulseaudio"
)

type Module struct {
	Color          string `json:"color,omitempty"`
	Background     string `json:"background,omitempty"`
	Border         string `json:"border,omitempty"`
	MinWidth       int    `json:"min_width,omitempty"`
	Align          string `json:"align,omitempty"`
	Urgent         bool   `json:"urgent,omitempty"`
	Separator      bool   `json:"separator,omitempty"`
	SeparatorWidth int    `json:"separator_block_width,omitempty"`
	Markup         string `json:"markup,omitempty"`

	Name     string `json:"name,omitempty"`
	Instance string `json:"instance,omitempty"`
}

type Config struct {
	Modules []Module `json:"modules"`
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

// === END: helper types for decode and encode to i3bar protocol ===

func handleInput() {
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
		fmt.Println(json.Marshal(click))
	}
}

func applyModuleConfigToMessage(module Module, message *Message) {
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

func module_datetime(output *Output, module Module) {

	// cron can be used by many modules
	c := cron.New()
	// === Start: DateTime setup ===
	formatDateTimeMsg := func() string {
		return time.Now().Format("15:04 1/Jan")
	}
	dateTimeMsg := &Message{
		FullText: formatDateTimeMsg(),
	}
	applyModuleConfigToMessage(module, dateTimeMsg)
	c.AddFunc("0 * * * * *", func() {
		dateTimeMsg.FullText = formatDateTimeMsg()
		output.printMsgs()
	})
	c.Start()
	// === End: DateTime setup ===
	output.messages = append(output.messages, dateTimeMsg)
}

func module_pulseaudio(output *Output, module Module) error {
	// === Start: PulseAudio setup ===
	pulse, e := pulseaudio.New()
	if e != nil {
		panic(e)
	}

	var pathToFallbackSink dbus.ObjectPath
	// Here we assume you are using fallbacksink, so we query that
	// altough later in pulse-callback we render whatever device you changed volume on :)
	pulse.Core().Get("FallbackSink", &pathToFallbackSink)
	volumes, e := pulse.Device(pathToFallbackSink).ListUint32("Volume")
	if e != nil {
		return e
	}
	baseVolume, e := pulse.Device(pathToFallbackSink).Uint32("BaseVolume")
	if e != nil {
		return e
	}
	formatPulseAudioText := func(volumes []uint32, baseVolume uint32) string {
		return fmt.Sprintf("vol %.2f", float32(volumes[0])/float32(baseVolume))
	}
	pulseAudioMsg := &Message{
		FullText: formatPulseAudioText(volumes, baseVolume),
	}
	applyModuleConfigToMessage(module, pulseAudioMsg)
	updatePulseMsg := func(volumes []uint32, baseVolume uint32) {
		pulseAudioMsg.FullText = formatPulseAudioText(volumes, baseVolume)
		output.printMsgs()
	}

	client := &Client{pulse, updatePulseMsg}
	pulse.Register(client)

	go pulse.Listen()
	// === End: PulseAudio setup ===
	output.messages = append(output.messages, pulseAudioMsg)
	return nil
}

type Output struct {
	encodeToStdout *json.Encoder
	messages       []*Message
}

func (o *Output) init() {
	// init stdout writer
	o.encodeToStdout = json.NewEncoder(os.Stdout)

	// then we need to specify protocol we use
	protocol := Protocol{
		Version: 1,
		// dont know about these params, found it somewhere, and it works..
		StopSignal:  int(syscall.Signal(10)),
		ContSignal:  int(syscall.Signal(12)),
		ClickEvents: true,
	}
	// and send it to Stdout
	o.encodeToStdout.Encode(protocol)
	// start our array of arrays.. (we never end it though)
	// kindof hacky, is there a better way ?
	fmt.Print("[")

	o.messages = []*Message{}
}

func (o Output) printMsgs() {
	o.encodeToStdout.Encode(o.messages)
	// And then separator between messages in our infinite array that never ends
	fmt.Print(",")
}

func main() {
	// Init log etc
	configFilePath := os.Args[1]
	configFile, error := os.Open(configFilePath)
	if error != nil {
		panic(error)
	}
	defer configFile.Close()
	configTxt, error := ioutil.ReadAll(configFile)
	if error != nil {
		panic(error)
	}
	var config Config
	error = json.Unmarshal(configTxt, &config)
	if error != nil {
		panic(error)
	}

	go handleInput()

	output := Output{}
	output.init()
	for _, module := range config.Modules {
		switch module.Name {
		case "datetime":
			module_datetime(&output, module)
		case "pulseaudio":
			module_pulseaudio(&output, module)
		}
	}
	output.printMsgs()
	//hacky way of blocking forever..
	select {}
}

// Pulse audio subscribe stuff
type Client struct {
	pulse          *pulseaudio.Client
	updatePulseMsg func(volumes []uint32, baseVolume uint32)
}

func (cl *Client) DeviceVolumeUpdated(path dbus.ObjectPath, values []uint32) {
	var baseVolume uint32
	cl.pulse.Device(path).Get("BaseVolume", &baseVolume)
	cl.updatePulseMsg(values, baseVolume)
}
