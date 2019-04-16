package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Module struct {
	Cron    string `json:"cron,omitempty"`
	Sprintf string `json:"sprintf,omitempty"`

	FilePath string `json:"file_path,omitempty"`

	DateTimeFormat string `json:"datetime_format,omitempty"`

	BarWidth  int64  `json:"bar_width"`
	BarFilled string `json:"bar_filled"`
	BarEmpty  string `json:"bar_empty"`

	CommandName string   `json:"command_name,omitempty"`
	CommandArgs []string `json:"command_args,omitempty"`

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
	Modules               []Module `json:"modules"`
	MinimumRenderInterval string   `json:"minimum_render_interval"`
	WMClient              string   `json:"wmclient,omitempty"`
}

func GetConfigFromPath(configFilePath string) Config {
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
	return config
}
