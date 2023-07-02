package static

import "embed"

// content holds our static web server content.
//go:embed dist/*
var Content embed.FS
