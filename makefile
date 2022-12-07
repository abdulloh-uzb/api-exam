CURRENT_DIR=$(shell pwd)

run:
	go run cmd/main.go

proto-gen:
	./scripts/gen-proto.sh	${CURRENT_DIR}
