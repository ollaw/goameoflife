package main

import (
	goameOfLife "gitlab.com/ollaww/goameoflife/src/game"
)

func main() {
	game := goameOfLife.CreateInstance()
	game.Start()
}
