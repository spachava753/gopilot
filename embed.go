package main

import "embed"

//go:embed terraform
var tfModules embed.FS
