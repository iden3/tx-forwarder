#!/usr/bin/env python3
"""Test endpoints for iden3/tx-forwarder
"""

import json
import requests
import provoj

URL = "http://127.0.0.1:11000/api/unstable"

addr = "0xbc88fcc53af747D30F1e17729659001d1129ddd7"

t = provoj.NewTest("tx-forwarder")

# Sample contract call
data = str.encode("test")
r = requests.post(URL + "/tx/sample", json={'addr': addr, 'dataHex': data.hex()})
t.rStatus("post forward tx", r)
jsonR = r.json()
print(jsonR)


# ZKPVerifier contract forward
zkpdata = {
        "a": ["0x14d6994cfd7b91d300be287ca6bbd80592ad8bdd097b3f05e15addc04115a20a", "0x2780097a84c727a3fba03f30c7931137fbe04efdabdabf726d7e067238ff39e8"],
    "b": [["0x0f2a93a2a0a55ef4cf3caa9f31c9225689bc5a7348f20c7ca1bb0c46ffd3df0a", "0x14b363739ed8c649f2fd001680cedd70f415027c7e1f9e79e07426a640e0a9da"],
        ["0x18bbb49de0fd32781a6e220ce4a41f23e0f1455b77e59605619926665bf0c941", "0x17bc9c9855472d25aaf2b66772986a8561361af9705f9ed04ee47ccfa8a8e17a"]],
    "c": ["0x1ae66d42ebd922b92e4ca1cd8b62b1ec896b0a629e9b98e63072b482145c683b", "0x07e44704aef1180a80ac66e3544545ce29bd9d9440c769783cf9c6bf7f96ec32"],
    "inputs":[ "0x0f35b130a033a10d021a636717b9c881e9d552e6e974bad03c18b47dbd2995e7",
          "0x1e6b7ac642d82edfdf550806de7193d126ce6b54fe02dbf326b5270cda87b532",
          "0x274dbce8d15179969bc0d49fa725bddf9de555e0ba6a693c6adb52fc9ee7a82c",
          "0x05ce98c61b05f47fe2eae9a542bd99f6b2e78246231640b54595febfd51eb853",
          "0x1bddd2ee218e2a2eaf8dc4ed626280424a5424197f4e68ff030b21b86a6d8a8f",
          "0x0b1e3150c37f3349941416396465c569ed015a11bf4d2d98d632dac8f6f17f44",
          "0x03a5c4b0163b38ac909f94b0f841141e15dad9fbd76987aab60105d99211a411",
          "0x247f2262a8ee3a73d452dfcfc20d1dfd04175d8cb018aec42ad8dcf7712ae7ea",
          "0x19779274226c66273b914b6847b206e277eedc0a242bf04c9ce1eb8734d9fb0e",
          "0x0000000000000000000000000102030405060708091011121314151617181920"],
        }
print(zkpdata)
r = requests.post(URL + "/tx/zkpverifier", json=zkpdata)
t.rStatus("zkpverifier post forward tx", r)
jsonR = r.json()
print(jsonR)


disabledata = {
        "mtp": "0x0001000000000000000000000000000000000000000000000000000000000001106211b6e71b0758fb754ee02927806ad700469465785e00715a6aa45330b12b",
  "id": "0x00002a1cd65ad84621144f479139efd0d08c69c6163685688b969b280fc184",
  "ethaddress": "0xc40966dd2c5af51ef1f431dc6b937ae1cab07be6",
  "msghash": "0xd1622712b967e14e10ad6be84125b0e271bc9ad999c23d737e45e524e09ecc81",
  "rsv": "0xfffad6502cb890b168c46fd04c56f5d532c3d872595c330d31e8e8258fb6c38d581efa2491785508a19c2739f591c32ac078238e52a7bd03b227be35833d36111c"
}
r = requests.post(URL + "/tx/disableid", json=disabledata)
t.rStatus("disableid post forward tx", r)
jsonR = r.json()
print(jsonR)

t.printScores()
