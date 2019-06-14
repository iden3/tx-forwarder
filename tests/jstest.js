const axios = require('axios');
const ethUtil = require('ethereumjs-util');
const bs58 = require('bs58');

const url = 'http://127.0.0.1:11000/api/unstable';

const proof = {
    a: ["0x14d6994cfd7b91d300be287ca6bbd80592ad8bdd097b3f05e15addc04115a20a", "0x2780097a84c727a3fba03f30c7931137fbe04efdabdabf726d7e067238ff39e8"],
    b: [["0x0f2a93a2a0a55ef4cf3caa9f31c9225689bc5a7348f20c7ca1bb0c46ffd3df0a", "0x14b363739ed8c649f2fd001680cedd70f415027c7e1f9e79e07426a640e0a9da"],
        ["0x18bbb49de0fd32781a6e220ce4a41f23e0f1455b77e59605619926665bf0c941", "0x17bc9c9855472d25aaf2b66772986a8561361af9705f9ed04ee47ccfa8a8e17a"]],
    c: ["0x1ae66d42ebd922b92e4ca1cd8b62b1ec896b0a629e9b98e63072b482145c683b", "0x07e44704aef1180a80ac66e3544545ce29bd9d9440c769783cf9c6bf7f96ec32"],
    inputs:[ "0x2f6e27cd0d29ca4b9b6e6db61ff1b61da3549b330ab2a20ec0d9631e290f62f6",
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
  };

axios.post(url + '/tx/zkpverifier', proof)
  .then(function (res) {
    // handle success
    console.log(res.data);
  })
  .catch(function (error) {
    // handle error
    console.log(error.res.data);
  });

// disableid

const privateKeyHex = '0x0102030405060708091011121314151617181920212223242526272829303132';
const privateKey = Buffer.from(privateKeyHex.substr(2), 'hex');
const kopHex = '0x966764905ac3e864c4bad1641659eda209b551b4cd78b08073db328b270a7f11';
const kdisHex = '0xc40966dd2c5af51ef1f431dc6b937ae1cab07be6';
const kreenHex = '0xe0fbce58cfaa72812103f003adce3f284fe5fc7c';

// Id genesis generated from above keys
const id = '112utFr3U7nqgq1g41qcfQhbyrvoTXuifV6umzer2F';
const idBytes = bs58.decode(id);
const idString = '0x' + idBytes.toString('hex');
const root = idBytes.slice(2, 29);
const rootString = '0x' + root.toString('hex');

// Proof claim generated from claimKethDis 
const proofHex = '0x0001000000000000000000000000000000000000000000000000000000000001'
                  +'106211b6e71b0758fb754ee02927806ad700469465785e00715a6aa45330b12b';

// Sign message random with Kdisable
const msg = Buffer.from('This is a test message');
const msgHash = ethUtil.hashPersonalMessage(msg);
const msgHashHex = ethUtil.bufferToHex(msgHash);
const sig = ethUtil.ecsign(msgHash, privateKey);
const sigHex = "0x" + Buffer.concat([sig.r,sig.s,ethUtil.toBuffer(sig.v)]).toString('hex');

const disableData = {
  mtp: proofHex,
  id: idString,
  ethaddress: kdisHex,
  msghash: msgHashHex,
  rsv: sigHex,
};
axios.post(url + '/tx/disableid', disableData)
  .then(function (res) {
    // handle success
    console.log(res.data);
  })
  .catch(function (error) {
    // handle error
    console.log(error.response.data);
  });
