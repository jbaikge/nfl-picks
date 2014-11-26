#!/bin/sh
ID=UpdateLines.`date +%Y-%m-%dT%T%:z`
curl -XPOST -H'Content-Type: application/json' http://picks.keyboardfu.com/rpc -d '{ "id":"'"$ID"'", "method": "Lines.ImportCurrent", "params": [] }'
