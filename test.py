#!/usr/bin/env python3
"""Test endpoints for iden3/tx-forwarder
"""

import json
import requests
import provoj

URL = "http://127.0.0.1:11000/api/unstable"

addr = "0xbc88fcc53af747D30F1e17729659001d1129ddd7"

t = provoj.NewTest("tx-forwarder")

data = str.encode("test")
r = requests.post(URL + "/tx", json={'addr': addr, 'dataHex': data.hex()})
t.rStatus("post forward tx", r)
jsonR = r.json()
print(jsonR)

t.printScores()
