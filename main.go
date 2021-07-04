package main

import (
	"kite/cmd"
	"kite/config"
)

func main() {
	config.Init()
	cmd.Execute()
}
