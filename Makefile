GOC=go build
BIN=bin
SRCS=$(shell find . -type f -name '*.go')
MIGRATIONS_MAIN=./cmd/migrations/main.go
LIPRODUCE_MAIN=./cmd/liproduce/main.go
LIPRODUCE_EXE=$(BIN)/liproduce
MIGRATIONS_EXE=$(BIN)/migrations

all: $(LIPRODUCE_EXE) $(MIGRATIONS_EXE)

$(LIPRODUCE_EXE): $(LIPRODUCE_MAIN) $(SRCS) $(BIN) 
	$(GOC) -o $@ $<

$(MIGRATIONS_EXE): $(MIGRATIONS_MAIN) $(BIN)
	$(GOC) -o $@ $<

$(BIN):
	mkdir $@

clean:
	rm -f $(LIPRODUCE_EXE)
	rm -f $(MIGRATIONS_EXE)

# MIGRATIONS/RUNNING
migrateinit: $(MIGRATIONS_EXE)
	$(MIGRATIONS_EXE) init

migrateup: $(MIGRATIONS_EXE)
	$(MIGRATIONS_EXE) up

migratedown: $(MIGRATIONS_EXE)
	$(MIGRATIONS_EXE) down

start: $(LIPRODUCE_EXE)
	$(LIPRODUCE_EXE)

dev: clean $(LIPRODUCE_EXE)
	$(LIPRODUCE_EXE)