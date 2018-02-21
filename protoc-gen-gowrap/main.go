package main

import (
	"os"

	"github.com/RangelReale/protoc-gowrap"
	"github.com/RangelReale/protoc-gowrap/generator"
)

func main() {
	g := generator.New()
	protoc_gowrap.Main(g, os.Stdin, os.Stdout)
}
