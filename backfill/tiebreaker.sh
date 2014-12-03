#!/bin/sh
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 1, "TieBreaker": { "Week": { "Week": 1, "Year": 2014 }, "Value": 40} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 1, "TieBreaker": { "Week": { "Week": 2, "Year": 2014 }, "Value": 44} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 1, "TieBreaker": { "Week": { "Week": 3, "Year": 2014 }, "Value": 45} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 1, "TieBreaker": { "Week": { "Week": 4, "Year": 2014 }, "Value": 46} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 1, "TieBreaker": { "Week": { "Week": 5, "Year": 2014 }, "Value": 46} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 1, "TieBreaker": { "Week": { "Week": 6, "Year": 2014 }, "Value": 44} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 1, "TieBreaker": { "Week": { "Week": 7, "Year": 2014 }, "Value": 46} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 1, "TieBreaker": { "Week": { "Week": 8, "Year": 2014 }, "Value": 45} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 1, "TieBreaker": { "Week": { "Week": 9, "Year": 2014 }, "Value": 50} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 1, "TieBreaker": { "Week": { "Week": 10, "Year": 2014 }, "Value": 49} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 1, "TieBreaker": { "Week": { "Week": 11, "Year": 2014 }, "Value": 47} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 1, "TieBreaker": { "Week": { "Week": 12, "Year": 2014 }, "Value": 51} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 2, "TieBreaker": { "Week": { "Week": 1, "Year": 2014 }, "Value": 44} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 2, "TieBreaker": { "Week": { "Week": 2, "Year": 2014 }, "Value": 50} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 2, "TieBreaker": { "Week": { "Week": 3, "Year": 2014 }, "Value": 50} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 2, "TieBreaker": { "Week": { "Week": 4, "Year": 2014 }, "Value": 45} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 2, "TieBreaker": { "Week": { "Week": 5, "Year": 2014 }, "Value": 56} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 2, "TieBreaker": { "Week": { "Week": 6, "Year": 2014 }, "Value": 44} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 2, "TieBreaker": { "Week": { "Week": 7, "Year": 2014 }, "Value": 40} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 2, "TieBreaker": { "Week": { "Week": 8, "Year": 2014 }, "Value": 54} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 2, "TieBreaker": { "Week": { "Week": 9, "Year": 2014 }, "Value": 48} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 2, "TieBreaker": { "Week": { "Week": 10, "Year": 2014 }, "Value": 49} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 2, "TieBreaker": { "Week": { "Week": 11, "Year": 2014 }, "Value": 47} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 2, "TieBreaker": { "Week": { "Week": 12, "Year": 2014 }, "Value": 60} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 3, "TieBreaker": { "Week": { "Week": 1, "Year": 2014 }, "Value": 47} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 3, "TieBreaker": { "Week": { "Week": 2, "Year": 2014 }, "Value": 41} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 3, "TieBreaker": { "Week": { "Week": 3, "Year": 2014 }, "Value": 44} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 3, "TieBreaker": { "Week": { "Week": 4, "Year": 2014 }, "Value": 48} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 3, "TieBreaker": { "Week": { "Week": 5, "Year": 2014 }, "Value": 57} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 3, "TieBreaker": { "Week": { "Week": 6, "Year": 2014 }, "Value": 46} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 3, "TieBreaker": { "Week": { "Week": 7, "Year": 2014 }, "Value": 48} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 3, "TieBreaker": { "Week": { "Week": 8, "Year": 2014 }, "Value": 49} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 3, "TieBreaker": { "Week": { "Week": 9, "Year": 2014 }, "Value": 43} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 3, "TieBreaker": { "Week": { "Week": 10, "Year": 2014 }, "Value": 43} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 3, "TieBreaker": { "Week": { "Week": 11, "Year": 2014 }, "Value": 45} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 3, "TieBreaker": { "Week": { "Week": 12, "Year": 2014 }, "Value": 52} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 4, "TieBreaker": { "Week": { "Week": 1, "Year": 2014 }, "Value": 47} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 4, "TieBreaker": { "Week": { "Week": 2, "Year": 2014 }, "Value": 55} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 4, "TieBreaker": { "Week": { "Week": 3, "Year": 2014 }, "Value": 42} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 4, "TieBreaker": { "Week": { "Week": 4, "Year": 2014 }, "Value": 40} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 4, "TieBreaker": { "Week": { "Week": 5, "Year": 2014 }, "Value": 53} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 4, "TieBreaker": { "Week": { "Week": 6, "Year": 2014 }, "Value": 41.5} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 4, "TieBreaker": { "Week": { "Week": 7, "Year": 2014 }, "Value": 42} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 4, "TieBreaker": { "Week": { "Week": 8, "Year": 2014 }, "Value": 44.5} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 4, "TieBreaker": { "Week": { "Week": 9, "Year": 2014 }, "Value": 46.5} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 4, "TieBreaker": { "Week": { "Week": 10, "Year": 2014 }, "Value": 49} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 4, "TieBreaker": { "Week": { "Week": 11, "Year": 2014 }, "Value": 53.5} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 4, "TieBreaker": { "Week": { "Week": 12, "Year": 2014 }, "Value": 48.5} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 5, "TieBreaker": { "Week": { "Week": 1, "Year": 2014 }, "Value": 35} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 5, "TieBreaker": { "Week": { "Week": 2, "Year": 2014 }, "Value": 42} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 5, "TieBreaker": { "Week": { "Week": 3, "Year": 2014 }, "Value": 40} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 5, "TieBreaker": { "Week": { "Week": 4, "Year": 2014 }, "Value": 39} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 5, "TieBreaker": { "Week": { "Week": 5, "Year": 2014 }, "Value": 37} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 5, "TieBreaker": { "Week": { "Week": 6, "Year": 2014 }, "Value": 40} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 5, "TieBreaker": { "Week": { "Week": 7, "Year": 2014 }, "Value": 48} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 5, "TieBreaker": { "Week": { "Week": 8, "Year": 2014 }, "Value": 47} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 5, "TieBreaker": { "Week": { "Week": 9, "Year": 2014 }, "Value": 47} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 5, "TieBreaker": { "Week": { "Week": 10, "Year": 2014 }, "Value": 52} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 5, "TieBreaker": { "Week": { "Week": 11, "Year": 2014 }, "Value": 52} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 5, "TieBreaker": { "Week": { "Week": 12, "Year": 2014 }, "Value": 56} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 6, "TieBreaker": { "Week": { "Week": 1, "Year": 2014 }, "Value": 43} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 6, "TieBreaker": { "Week": { "Week": 2, "Year": 2014 }, "Value": 42} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 6, "TieBreaker": { "Week": { "Week": 3, "Year": 2014 }, "Value": 37} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 6, "TieBreaker": { "Week": { "Week": 4, "Year": 2014 }, "Value": 39} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 6, "TieBreaker": { "Week": { "Week": 5, "Year": 2014 }, "Value": 56} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 6, "TieBreaker": { "Week": { "Week": 6, "Year": 2014 }, "Value": 47} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 6, "TieBreaker": { "Week": { "Week": 7, "Year": 2014 }, "Value": 39} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 6, "TieBreaker": { "Week": { "Week": 8, "Year": 2014 }, "Value": 46} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 6, "TieBreaker": { "Week": { "Week": 9, "Year": 2014 }, "Value": 50} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 6, "TieBreaker": { "Week": { "Week": 10, "Year": 2014 }, "Value": 41} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 6, "TieBreaker": { "Week": { "Week": 11, "Year": 2014 }, "Value": 51} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 6, "TieBreaker": { "Week": { "Week": 12, "Year": 2014 }, "Value": 55} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 7, "TieBreaker": { "Week": { "Week": 1, "Year": 2014 }, "Value": 48} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 7, "TieBreaker": { "Week": { "Week": 2, "Year": 2014 }, "Value": 45} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 7, "TieBreaker": { "Week": { "Week": 3, "Year": 2014 }, "Value": 44} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 7, "TieBreaker": { "Week": { "Week": 4, "Year": 2014 }, "Value": 41} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 7, "TieBreaker": { "Week": { "Week": 5, "Year": 2014 }, "Value": 46} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 7, "TieBreaker": { "Week": { "Week": 6, "Year": 2014 }, "Value": 48} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 7, "TieBreaker": { "Week": { "Week": 7, "Year": 2014 }, "Value": 43} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 7, "TieBreaker": { "Week": { "Week": 8, "Year": 2014 }, "Value": 48} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 7, "TieBreaker": { "Week": { "Week": 9, "Year": 2014 }, "Value": 48} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 7, "TieBreaker": { "Week": { "Week": 10, "Year": 2014 }, "Value": 51} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 7, "TieBreaker": { "Week": { "Week": 11, "Year": 2014 }, "Value": 50} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 7, "TieBreaker": { "Week": { "Week": 12, "Year": 2014 }, "Value": 50} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 8, "TieBreaker": { "Week": { "Week": 1, "Year": 2014 }, "Value": 39} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 8, "TieBreaker": { "Week": { "Week": 2, "Year": 2014 }, "Value": 47} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 8, "TieBreaker": { "Week": { "Week": 3, "Year": 2014 }, "Value": 41} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 8, "TieBreaker": { "Week": { "Week": 4, "Year": 2014 }, "Value": 42} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 8, "TieBreaker": { "Week": { "Week": 5, "Year": 2014 }, "Value": 48} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 8, "TieBreaker": { "Week": { "Week": 6, "Year": 2014 }, "Value": 49} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 8, "TieBreaker": { "Week": { "Week": 7, "Year": 2014 }, "Value": 41} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 8, "TieBreaker": { "Week": { "Week": 8, "Year": 2014 }, "Value": 49} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 8, "TieBreaker": { "Week": { "Week": 9, "Year": 2014 }, "Value": 49} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 8, "TieBreaker": { "Week": { "Week": 10, "Year": 2014 }, "Value": 42} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 8, "TieBreaker": { "Week": { "Week": 11, "Year": 2014 }, "Value": 46} } ] }'
curl -XPOST -H'Content-Type:application/json' 127.1:10000/rpc -d'{ "method": "TieBreaker.Submit", "id": 3, "params": [ { "UserId": 8, "TieBreaker": { "Week": { "Week": 12, "Year": 2014 }, "Value": 48} } ] }'