GO=go

# Change variables according to os
ifeq ($(OS), Windows_NT)
	BIN=.\bin
	EXE=.exe
	CREATE_BIN=if not exist "$(BIN)" mkdir $(BIN)
	CLEAN=Rmdir /S/Q .\bin
else
	BIN=./bin
	CREATE_BIN=mkdir -p $(BIN)
	CLEAN=rm -rf $(BIN)
endif

build:
	$(CREATE_BIN)
	$(GO) build -o $(BIN)/liproduce$(EXE) ./cmd/liproduce 
	$(GO) build -o $(BIN)/migrations$(EXE) ./cmd/migrations 

clean:
	$(CLEAN)	

gen:
	go run github.com/99designs/gqlgen

start:
	$(BIN)/liproduce$(EXE)

dev: build start