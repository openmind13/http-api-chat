ENTRY_POINT = ./cmd/chatapi
BIN_NAME = ./chatapi

build:
	go build -v $(ENTRY_POINT)

run:
	$(BIN_NAME)

build_and_run:
	sudo service postgresql start && go build -v $(ENTRY_POINT) && echo "\n" && $(BIN_NAME)


.DEFAULT_GOAL := build_and_run