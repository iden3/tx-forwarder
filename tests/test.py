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
    "inputs":[ "0x2f6e27cd0d29ca4b9b6e6db61ff1b61da3549b330ab2a20ec0d9631e290f62f6",
            "0x0448a124f20e3d56aedc6157c967a43b9bbad690db11da462d0e85a6ace182b2",
            "0x0442b6fd8d2ee686b8ae1fb6ac7f82dbc82706c2937fa7299e48c82f61a444c0",
            "0x274dbce8d15179969bc0d49fa725bddf9de555e0ba6a693c6adb52fc9ee7a82c",
            "0x05ce98c61b05f47fe2eae9a542bd99f6b2e78246231640b54595febfd51eb853",
            "0x12a111c28a644cfbad11793464dc448ebe3a1c4c965415c9fe060139531704a1",
            "0x03b2c0e02ad9d172c8deed564553972149075dfa07a9cd4563f3df1ee3fc9343",
            "0x03a5c4b0163b38ac909f94b0f841141e15dad9fbd76987aab60105d99211a411",
            "0x247f2262a8ee3a73d452dfcfc20d1dfd04175d8cb018aec42ad8dcf7712ae7ea",
            "0x18aad37ad7ab569a0eec5ea49747d98604a41a485682f2c6bce9dac8ecd77888",
            "0x0000000000000000000000000102030405060708091011121314151617181920"],
        }
print(zkpdata)
r = requests.post(URL + "/tx/zkpverifier", json=zkpdata)
t.rStatus("post forward tx", r)
jsonR = r.json()
print(jsonR)

t.printScores()
