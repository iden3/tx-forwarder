const axios = require('axios');
const ethUtil = require('ethereumjs-util');
const bs58 = require('bs58');

const url = 'http://127.0.0.1:11000/api/unstable';

// get tx
const txid = '0xe991ed8da5d304224d935da6a69aca7fc4c94bb81366479b5a5ff095d2061004';
axios.get(url + '/tx/' + txid)
  .then(function (res) {
    // handle success
    console.log("tx/:txhash", res.data);
  })
  .catch(function (error) {
    // handle error
    console.log("error", error.res);
  });

// zkpverifier
const proof = {
  a: ["0x0eaab1567625b698a9579be2873ce9bf4dc8b369768ef64d9ef9145e3ab2a5f4", "0x26ce6c3c50f6d3c7ab0d9c25818628755cda2960ec1772e4db9264675f3df48e"],
  b: [["0x08cb8aa60e5f4b9ddfb8a3f27fb250d69f9bec999bcb7f547a541fd98831ad03", "0x21d1eec48a0866d915345768b0da5b8d46c88dc2fcf8c8a9d5a305454e2458c1"],
      ["0x06cd98bb696185fe988746e4ffc35d42b7fc2774d098b70ac606ba911b805853", "0x034a3f8ba380fbf1621b88ad664c9b129e61122d7c729b724a6170b577baecb7"]],
  c: ["0x29a8e58d1e370820ab7c093287333e6ceed8b56708f047b1db37a57fba36ccb0", "0x13698c2758716fcfc97258438cc8adfa66731a9891290914ba36bd41ad7fc627"],
  inputs:[ "0x0f35b130a033a10d021a636717b9c881e9d552e6e974bad03c18b47dbd2995e7",
          "0x1e6b7ac642d82edfdf550806de7193d126ce6b54fe02dbf326b5270cda87b532",
          "0x274dbce8d15179969bc0d49fa725bddf9de555e0ba6a693c6adb52fc9ee7a82c",
          "0x05ce98c61b05f47fe2eae9a542bd99f6b2e78246231640b54595febfd51eb853",
          "0x1bddd2ee218e2a2eaf8dc4ed626280424a5424197f4e68ff030b21b86a6d8a8f",
          "0x0b1e3150c37f3349941416396465c569ed015a11bf4d2d98d632dac8f6f17f44",
          "0x03a5c4b0163b38ac909f94b0f841141e15dad9fbd76987aab60105d99211a411",
          "0x247f2262a8ee3a73d452dfcfc20d1dfd04175d8cb018aec42ad8dcf7712ae7ea",
          "0x19779274226c66273b914b6847b206e277eedc0a242bf04c9ce1eb8734d9fb0e",
          "0x0000000000000000000000000102030405060708091011121314151617181920"],
};

axios.post(url + '/tx/zkpverifier', proof)
  .then(function (res) {
    console.log(res.data);
    let ethTx = res.data.ethTx;
    console.log("getting tx", ethTx)
    // axios.get(url + '/tx/' + ethTx)
    //   .then(function (res) {
    //     console.log("tx/:txhash", res.data);
    //   })
    //   .catch(function (error) {
    //     console.log("error", error.res);
    //   });
  })
  .catch(function (error) {
    console.log("error:", error.res);
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
/*
axios.post(url + '/tx/disableid', disableData)
  .then(function (res) {
    // handle success
    console.log(res.data);
  })
  .catch(function (error) {
    // handle error
    console.log(error.response.data);
  });
  */
