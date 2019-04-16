#!/bin/bash
icon=""

if [ "$(pgrep -x redshift)" ]; then
    temp=$(redshift -p 2> /dev/null | grep temp | cut -d ":" -f 2 | tr -dc "[:digit:]")
    if [ -z "$temp" ]; then
        echo "<span fgcolor='#65737E'>$icon</span>"
    elif [ "$temp" -ge 5000 ]; then
        echo "<span fgcolor='#8FA1B3'>$icon</span>"
    elif [ "$temp" -ge 4000 ]; then
        echo "<span fgcolor='#EBCB8B'>$icon</span>"
    else
        echo "<span fgcolor='#D08770'>$icon</span>"
    fi
else
    echo "<span fgcolor='#FFFFFF'>$icon(off)</span>"
fi
