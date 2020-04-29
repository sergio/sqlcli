.PHONY: build
build:
	go build -o sqlcli .

.PHONY: install
install: build
	cp sqlcli /usr/local/bin/sqlcli
