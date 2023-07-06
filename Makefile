.PHONY: gen/server
gen/server:
	buf generate

.PHONY: ui
ui:
	go generate ui/static.go

.PHONY: bin
bin: ui
	go build -o shrls cmd/shrls/main.go

.PHONY: dist
dist: ui
	go build -o shrls -ldflags="-s -w" cmd/shrls/main.go
