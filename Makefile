version := $(shell git describe --dirty=-changes)
build := $(shell git rev-parse HEAD)
build_date := $(shell date +"%Y-%m-%dT%H:%M:%S%z")

.PHONY: proto
proto:
	buf generate

.PHONY: ui
ui: proto
	cd ui && npm i
	go generate ui/static.go

.PHONY: bin
bin: ui
	go build -o shrls -ldflags="\
		-X github.com/demophoon/shrls/pkg/version.Version=${version}-dev \
		-X github.com/demophoon/shrls/pkg/version.Build=${build} \
		-X github.com/demophoon/shrls/pkg/version.BuildDate=${build_date} \
	" cmd/shrls/main.go

.PHONY: dist
dist: ui
	CGO_ENABLED=0 GOOS=linux go build -o shrls -ldflags="\
		-s -w \
		-X github.com/demophoon/shrls/pkg/version.Version=${version} \
		-X github.com/demophoon/shrls/pkg/version.Build=${build} \
		-X github.com/demophoon/shrls/pkg/version.BuildDate=${build_date} \
	" cmd/shrls/main.go

.PHONY: docker
docker:
	docker build . \
		--label "org.opencontainers.image.source=https://github.com/demophoon/shrls" \
		--label "org.opencontainers.image.description=Simple and small url shortener" \
		--label "org.opencontainers.image.licenses=Apache-2.0" \
		--label "org.opencontainers.image.version=${version}" \
		--label "org.opencontainers.image.ref.name=${build}" \
		--label "org.opencontainers.image.created=${build_date}" \
		-t ghcr.io/demophoon/shrls:${version}

.PHONY: publish
publish:
	docker push ghcr.io/demophoon/shrls:${version}
