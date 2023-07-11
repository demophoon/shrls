version := $(shell git describe --dirty=-changes)

.PHONY: proto
proto:
	buf generate

.PHONY: ui
ui: proto
	go generate ui/static.go

.PHONY: bin
bin: ui
	go build -o shrls -ldflags="-X gitlab.cascadia.demophoon.com/demophoon/go-shrls/pkg/cli.version=${version}-dev" cmd/shrls/main.go

.PHONY: dist
dist: ui
	CGO_ENABLED=0 GOOS=linux go build -o shrls -a --installsuffix cgo -ldflags="-s -w -X gitlab.cascadia.demophoon.com/demophoon/go-shrls/pkg/cli.version=${version}" cmd/shrls/main.go

.PHONY: docker
docker:
	docker build . -t shrls:${version}
