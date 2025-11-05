package main

import (
	"github.com/dvgamerr-app/go-bitkub/cmd"
)

func main() {
	cmd.Version = "dev"
	cmd.Commit = "none"
	cmd.Date = "unknown"
	cmd.Execute()
}
