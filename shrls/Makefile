.PHONY: gen/server
gen/server:
	buf generate

.PHONY: bin
bin:
	go build -o shrls cmd/shrls/main.go

.PHONY: dist
dist:
	go build -o shrls -ldflags="-s -w" cmd/shrls/main.go
