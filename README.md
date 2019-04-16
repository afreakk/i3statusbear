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
    status_command go run ~/go/src/github.com/afreakk/i3statusbear/main.go ~/go/src/github.com/afreakk/i3statusbear/exampleConfigs/mainbar.json
	output $primaryScreen
	tray_output none
    colors {
        statusline #ffffff
        background #323232
        inactive_workspace #32323200 #32323200 #5c5c5c
    }
}

bar {
    position top
	workspace_buttons no
    status_command go run ~/go/src/github.com/afreakk/i3statusbear/main.go ~/go/src/github.com/afreakk/i3statusbear/exampleConfigs/offscreenbartop.json
	output $offScreen
	tray_output none
    colors {
        statusline #ffffff
        background #323232
        inactive_workspace #32323200 #32323200 #5c5c5c
    }
}

bar {
    position bottom
    status_command go run ~/go/src/github.com/afreakk/i3statusbear/main.go ~/go/src/github.com/afreakk/i3statusbear/exampleConfigs/offscreenbarbottom.json
	output $offScreen
    colors {
        statusline #ffffff
        background #323232
        inactive_workspace #32323200 #32323200 #5c5c5c
    }
}

```
