#!/bin/bash
# rm -rf store.go build
./solc --abi store.sol -o build --overwrite
./solc --bin store.sol -o build --overwrite
./abigen --bin=./build/Store.bin --abi=./build/Store.abi --pkg=store --out=store.go
