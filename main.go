package main

import cmd "github.com/meyermarcel/icm/cmd/icm"

var version = "dev"

func main() {
	cmd.Execute(version)
}
