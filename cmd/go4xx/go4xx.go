package main

import (
	"github.com/TGPrado/go4xx/cmd/inputReader"
	"github.com/TGPrado/go4xx/internal/launcher"
)

func main() {
	flags := launcher.ParseFlags()
	fileData := make(chan string)
	inputreader.ReadInputAsync(flags.FileName, fileData)
}
