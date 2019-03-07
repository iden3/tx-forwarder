# Gas Station
Gas station for ethereum blockchain.

Server that pays for client signed transactions.





### Contract
The contract handlers goes inside `eth/contract/` directory.

Once having the contract file written in Solidity, in order to generate the Go handlers:
```
solc --abi --bin SampleContract.sol -o build
abigen --bin=./build/SampleContract.bin --abi=./build/SampleContract.abi --pkg=SampleContract --out=SampleContract.go
```
And place the `sampleContract.go` file in the `eth/contract/sampleContract.go` path.
