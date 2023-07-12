version := $(shell git describe --dirty=-changes)

.PHONY: proto
proto:
	buf generate

.PHONY: ui
ui: proto
	cd ui && npm i
	go generate ui/static.go

.PHONY: bin
bin: ui
	go build -o shrls -ldflags="-X gitlab.cascadia.demophoon.com/demophoon/go-shrls/pkg/cli.version=${version}-dev" cmd/shrls/main.go

.PHONY: dist
dist: ui
	CGO_ENABLED=0 GOOS=linux go build -o shrls -a --installsuffix cgo -ldflags="-s -w -X gitlab.cascadia.demophoon.com/demophoon/go-shrls/pkg/cli.version=${version}" cmd/shrls/main.go

.PHONY: docker
docker:
	docker build . \
		--label "org.opencontainers.image.source=https://github.com/demophoon/shrls" \
		--label "org.opencontainers.image.description=Simple and small url shortener" \
		--label "org.opencontainers.image.licenses=Apache-2.0" \
		-t ghcr.io/demophoon/shrls:${version}

.PHONY: publish
publish:
	docker push ghcr.io/demophoon/shrls:${version}
