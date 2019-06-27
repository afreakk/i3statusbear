#!/bin/bash
USD_PRICE=$(curl https://api.coinmarketcap.com/v1/ticker/$1/ 2> /dev/null | jq -r '.[0] | .price_usd')
#we only want first 5 difits, any more is overkill
#and use sed to remove last character if it is a decimal separator
echo $USD_PRICE | awk -F. '{print $1}'
