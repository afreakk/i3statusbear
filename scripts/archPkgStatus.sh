#!/bin/bash
pac=$(checkupdates | wc -l)
aur=$(auracle sync | wc -l)
echo "<span fgcolor='#d79921'>$pac</span> ⇅ <span fgcolor='#d79921'>$aur</span>"
