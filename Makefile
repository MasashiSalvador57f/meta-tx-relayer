build_contracts:
	cd contracts; truffle build;

migrate_reset_local:
	cd contracts; truffle migrate --reset --network local

extract_abi:
	/bin/bash scripts/extract_abi.sh

abigen_all:
	docker run -v ${WORKDIR}:/home ethereum/client-go:alltools-stable abigen --abi /home/contracts/build/contracts/SimpleStorage.json.abi --pkg contract --type SimpleStorage --out /home/simple_storage.go
	mv simple_storage.go api/contracts/
	docker run -v ${WORKDIR}:/home ethereum/client-go:alltools-stable abigen --abi /home/contracts/build/contracts/TxRelay.json.abi --pkg contract --type TxRelay --out /home/tx_relay.go
	mv tx_relay.go api/contracts/
	docker run -v ${WORKDIR}:/home ethereum/client-go:alltools-stable abigen --abi /home/contracts/build/contracts/SignValidator.json.abi --pkg contract --type SignValidator --out /home/sign_validator.go
	mv sign_validator.go api/contracts/

.PHONY: reset_all
reset_all: migrate_reset_local extract_abi abigen_all