version := $(shell git describe --dirty=+changes)

.PHONY: gen/server
gen/server:
	buf generate

.PHONY: ui
ui:
	go generate ui/static.go

.PHONY: bin
bin: ui
	echo ${version}
	go build -o shrls -ldflags="-X gitlab.cascadia.demophoon.com/demophoon/go-shrls/pkg/cli.version=${version}-dev" cmd/shrls/main.go

.PHONY: dist
dist: ui
	go build -o shrls -ldflags="-s -w -X gitlab.cascadia.demophoon.com/demophoon/go-shrls/pkg/cli.version=${version}" cmd/shrls/main.go
