.POSIX:

include config.mk

all: star

star:
	$(GO) build $(GOFLAGS)

clean:
	$(RM) star

install: all
	mkdir -p $(DESTDIR)$(PREFIX)/bin
	cp -f star $(DESTDIR)$(PREFIX)/bin
	chmod 755 $(DESTDIR)$(PREFIX)/bin/star

uninstall:
	$(RM) $(DESTDIR)$(PREFIX)/bin/star

.DEFAULT_GOAL := all

.PHONY: all star clean install uninstall
