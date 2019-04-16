# StatusBear
## Go code for talking with [i3 bar](https://i3wm.org/i3bar/) using [i3 protocol](https://i3wm.org/docs/i3bar-protocol.html) and creating usefull and fast bars.
## Also works with [Sway](https://swaywm.org/)!

![Example](https://github.com/afreakk/i3statusbear/blob/master/imgs/bars.png "Example bars")

## Modules
- activewindow
- command (run bash scripts etc)
- cpu (usage)
- datetime
- memory (usage)
- pulseaudio (volume bar)
- readfile (read arbitrary file)

## Example usage (from my swaywm config)
```
bar {
    position bottom
	id main_bar
    status_command go run ~/go/src/github.com/afreakk/i3statusbear/main.go ~/go/src/github.com/afreakk/i3statusbear/exampleConfigs/mainbar.json
	output $primaryScreen
	tray_output none
	separator_symbol ""
    colors {
        background #000000
        inactive_workspace #000000 #000000 #00FFD5
        focused_workspace #000000 #00FFD5 #FD00E1
		active_workspace #000000 #00FFD540 #000000
		urgent_workspace #000000 #FD00E1 #000000
    }
}

bar {
    position top
	workspace_buttons no
    status_command go run ~/go/src/github.com/afreakk/i3statusbear/main.go ~/go/src/github.com/afreakk/i3statusbear/exampleConfigs/offscreenbartop.json
	output $offScreen
	tray_output none
	id off_bar_top
	bindsym button4 exec ~/bin/setSinkVolumeDefault.sh +5%
	bindsym button5 exec ~/bin/setSinkVolumeDefault.sh -5%
	separator_symbol ""
    colors {
        background #000000
    }
}

bar {
    position bottom
	id off_bar_bottom
    status_command go run ~/go/src/github.com/afreakk/i3statusbear/main.go ~/go/src/github.com/afreakk/i3statusbear/exampleConfigs/offscreenbarbottom.json
	output $offScreen
	separator_symbol ""
    colors {
        background #000000
        inactive_workspace #000000 #000000 #00FFD5
        focused_workspace #000000 #00FFD5 #FD00E1
		active_workspace #000000 #00FFD540 #000000
    }
}

```
