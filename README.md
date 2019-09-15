# StatusBear
## Go code for talking with [i3 bar](https://i3wm.org/i3bar/) using [i3 protocol](https://i3wm.org/docs/i3bar-protocol.html) and creating usefull and fast bars.
## Also works with [Sway](https://swaywm.org/)!
### Set sway or i3 as 2nd argument, or set wmclient attribute on config.json. (argument will override attribute)

![Example](https://github.com/afreakk/i3statusbear/blob/master/imgs/bars.png "Example bars")

## Modules
- activewindow
- command (run bash scripts etc)
- cpu (usage)
- datetime
- memory (usage)
- pulseaudio (volume bar)
- readfile (read arbitrary file)

## Example usage
```
bar {
	id main_bar
	position bottom
	status_command exec ~/go/bin/i3statusbear ~/go/src/github.com/afreakk/i3statusbear/exampleConfigs/mainbar.json i3 2>> ~/log/statusbear.main_bar.log
	output $primaryScreen
	tray_output none
	bindsym button3 exec "rofi -modi combi -show combi -combi-modi run,drun,window"
	separator_symbol " "
	colors {
		background #282828
		inactive_workspace	#282828 #282828	#685d52
		focused_workspace	#282828 #d79921	#282828
		active_workspace	#282828 #685d52	#282828
		urgent_workspace	#282828 #fb4934	#282828
	}
}

bar {
	id off_bar_top
	position top
	status_command exec ~/go/bin/i3statusbear ~/go/src/github.com/afreakk/i3statusbear/exampleConfigs/offscreenbartop.json i3 2>> ~/log/statusbear.off_bar_top.log
	output $offScreen
	tray_output none
	bindsym button4 exec ~/bin/setSinkVolumeDefault.sh +5%
	bindsym button5 exec ~/bin/setSinkVolumeDefault.sh -5%
	separator_symbol " "
	colors {
		background #282828
		inactive_workspace	#282828 #282828	#685d52
		focused_workspace	#282828 #d79921	#282828
		active_workspace	#282828 #685d52	#282828
		urgent_workspace	#282828 #fb4934	#282828
	}
	workspace_buttons no
}

bar {
	id off_bar_bottom
	position bottom
	status_command exec ~/go/bin/i3statusbear ~/go/src/github.com/afreakk/i3statusbear/exampleConfigs/offscreenbarbottom.json i3 2>> ~/log/statusbear.off_bar_bottom.log
	output $offScreen
	tray_output $offScreen
	bindsym button3 exec "rofi -modi combi -show combi -combi-modi run,drun,window"
	separator_symbol " "
	colors {
		background #282828
		inactive_workspace	#282828 #282828	#685d52
		focused_workspace	#282828 #d79921	#282828
		active_workspace	#282828 #685d52	#282828
		urgent_workspace	#282828 #fb4934	#282828
	}
}

```
