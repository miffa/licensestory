GOCMD=go
GOBUILD=$(GOCMD) build
GOINSTALL=$(GOCMD) install
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

BINARY_NAME=milicense
BIN_PATH=bin

.PHONY: build clean run all

all: clean build install

hello:
	@echo "Hello"

build:
	@[ ! -d $(BIN_PATH) ] && mkdir $(BIN_PATH) || echo bin  dir ok
	@echo build milicense  
	$(GOBUILD) \
                        -ldflags "-w -s"\
                        -o $(BIN_PATH)/$(BINARY_NAME) \
                        main.go 


clean:
	@echo cleaning 
	rm -rf $(BIN_PATH)/$(BINARY_NAME) && $(GOCLEAN)

test:
	$(GOTEST)  -v ./...

install:
	@echo installing  
	$(GOINSTALL) -ldflags "-w -s" 
