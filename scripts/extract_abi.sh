#!/bin/bash
for json in `ls contracts/build/contracts/*.json`; do cat $json | jq .abi > $json.abi; done;