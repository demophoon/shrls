package ui

import "embed"

// content holds our static web server content.
//
//go:generate sh -c "npm run build-prod"
//go:generate sh -c "cp ../server/gen/shrls.swagger.json ./dist/"
//go:embed dist/*
var Content embed.FS
