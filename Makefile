PREFIX=/usr/local
BINDIR=${PREFIX}/bin
BLDDIR = build
BLDFLAGS=
EXT=
ifeq (${GOOS},windows)
	EXT=.exe
endif

APPS = generator
all: $(APPS)

$(BLDDIR)/generator: $(wildcard apps/generator/*.go)

$(BLDDIR)/%:
	@mkdir -p $(dir $@)
	go build ${BLDFLAGS} -o $@ ./apps/$*

$(APPS): %: $(BLDDIR)/%

clean:
	rm -fr $(BLDDIR)

test:
	go test ./...