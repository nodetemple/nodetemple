package main

import (
	"github.com/nodetemple/nodetemple/version"
)

var (
	cmdVersion = &Command{
		Name:			"version",
		Description:	"Print the version and exit",
		Summary:		"Print the version and exit",
		Usage:			"",
		Run:			runVersion,
	}
)

func runVersion(args []string) (exit int) {
	stdout("%s version %s", cliName, version.Version)
	return
}
