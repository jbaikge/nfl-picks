#!/bin/sh
ID=UpdateWeek.`date +%Y-%m-%dT%T%:z`
curl -XPOST -H'Content-Type: application/json' http://picks.keyboardfu.com/rpc -d '{ "id":"'"$ID"'", "method": "Game.UpdateCurrentWeek", "params": [] }'
