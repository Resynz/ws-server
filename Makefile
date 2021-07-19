SHELL=/bin/bash

EXE = ws-server

all: $(EXE)

ws-server:
	@echo "building $@ ..."
	$(MAKE) -s -f make.inc s=static

clean:
	rm -f $(EXE)

