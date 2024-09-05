package main

import (
	"go.uber.org/fx"

	bundlefx "scaffold/bundle"
)

func main() {
	fx.New(bundlefx.Bundle).Run()
}
