package main

const (
	Version = "0.0.1"
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
	stdout("%s version %s", cliName, Version)
	return
}
