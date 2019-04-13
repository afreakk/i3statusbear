# StatusBear
# Go code for talking with [i3 bar](https://i3wm.org/i3bar/) using [i3 protocol](https://i3wm.org/docs/i3bar-protocol.html) and creating usefull and fast bars for multiple monitors.  
# WIP.


# Example usage (in swaywm)
```
bar {
    position bottom
    status_command go run ~/go/src/github.com/afreakk/i3statusbear/main.go ~/go/src/github.com/afreakk/i3statusbear/configExample.json
	output $primaryScreen
    colors {
        statusline #ffffff
        background #323232
        inactive_workspace #32323200 #32323200 #5c5c5c
    }
}
```