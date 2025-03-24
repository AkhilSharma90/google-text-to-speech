package main

import (
	"os"

	"github.com/akhilsharma90/google-text-to-speech/cmd"
)

func main() {
	cli := &cmd.CLI{ErrStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}
