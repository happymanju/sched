package main

import (
	"os"

	"github.com/happymanju/sched/cmd"
)

func main() {
	os.Exit(cmd.Run(os.Args[0:]))
}
