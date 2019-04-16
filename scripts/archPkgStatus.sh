#!/bin/bash
pac=$(checkupdates | wc -l)
aur=$(auracle sync | wc -l)
echo "<span fgcolor='#FD00E1'>$pac</span> ⇅ <span fgcolor='#FD00E1'>$aur</span>"
