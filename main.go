package main

import (
	"advancedGolang/cli"
	"advancedGolang/db"
)

func main() {
	defer db.Close()

	cli.Start()
}
