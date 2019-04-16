#!/bin/bash
df $1 -h| tail -n 1| awk '{ print $4 "/" $2 }'
