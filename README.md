# tx-forwarder
Server that forwards the tx to the specified smart contract.


## Usage
In production use `./tx-forwarder`, in development use `go run main.go` in order to make sure to be running last @master version.

### Config
Deploy needed smart contracts, each one of the commands, will return the contract deployed address, that will be needed to be placed in the `config.yaml`.
```
./tx-forwarder deploy mimc7
./tx-forwarder deploy rootcommits
./tx-forwarder deploy whitelist
./tx-forwarder deploy zkpverifier
./tx-forwarder deploy disableid
```

This will print the deployed contract address, then copy&paste in the config file `config.yaml`:
```
server:
        serviceapi: 0.0.0.0:11000
        adminapi: 0.0.0.0:11001
web3:
        url: http://127.0.0.1:8545
keystore:
        path: /var/config/keystore
        address: 0x123456789...
        passwd: /var/config/keystore.password
        keyjsonpath: /var/config/keystore/UTC-etc
contracts:
        samplecontract: 0xasdf
        mimc7contract: 0xasdf
        rootcommitscontract: 0xasdf
        whitelistcontract: 0xasdf
        zkpverifiercontract: 0xasdf
        disableidcontract: 0xasdf
```

### Run
Then, run the server:
```
./tx-forwarder start
```

Also can info can be getted by running:
```
./tx-forwarder info
```


### Contract
Current contract is just a working sample contract.

The contract handlers goes inside `eth/contract/` directory.

Once having the contract file written in Solidity, in order to generate the Go handlers:
```
solc --abi --bin SampleContract.sol -o build
abigen --bin=./build/SampleContract.bin --abi=./build/SampleContract.abi --pkg=SampleContract --out=SampleContract.go

// or directly (use this one)
abigen --sol SampleContract.sol --pkg SampleContract --out=SampleContract.go
```
And place the `sampleContract.go` file in the `eth/contract/sampleContract.go` path.
