package main

import (
	"flag"
	"fmt"
	"strings"
	"text/template"
)

const (
	// used to indicate flag usage should not be printed
	hidden = "hidden"
)

var (
	cmdHelp = &Command{
		Name:			"help",
		Description:	"Show a list of commands or detailed help for one command",
		Summary:		"Show a list of commands or help for one command",
		Usage:			"[COMMAND]",
		Run:			runHelp,
	}

	globalUsageTemplate  *template.Template
	commandUsageTemplate *template.Template
	templFuncs           = template.FuncMap{
		"descToLines": func(s string) []string {
			// trim leading/trailing whitespace and split into slice of lines
			return strings.Split(strings.Trim(s, "\n\t "), "\n")
		},
		"printOption": func(name, defvalue, usage string) string {
			if usage == hidden {
				return ""
			}
			prefix := "--"
			if len(name) == 1 {
				prefix = "-"
			}
			return fmt.Sprintf("\n\t%s%s=%s\t%s", prefix, name, defvalue, usage)
		},
	}
)

func init() {
	globalUsageTemplate = template.Must(template.New("global_usage").Funcs(templFuncs).Parse(`
NAME:
{{printf "\t%s - %s" .Executable .Description}}

USAGE: 
{{printf "\t%s" .Executable}} [global options] <command> [command options] [arguments...]

VERSION:
{{printf "\t%s" .Version}}

COMMANDS:{{range .Commands}}
{{printf "\t%s\t%s" .Name .Summary}}{{end}}

GLOBAL OPTIONS:{{range .Flags}}{{printOption .Name .DefValue .Usage}}{{end}}

Global options can also be configured via upper-case environment variables prefixed with "{{.EnvFlag}}_"
For example, "some-flag" => "{{.EnvFlag}}_SOME_FLAG"

Run "{{.Executable}} help <command>" for more details on a specific command.
`[1:]))
	commandUsageTemplate = template.Must(template.New("command_usage").Funcs(templFuncs).Parse(`
NAME:
{{printf "\t%s - %s" .Cmd.Name .Cmd.Summary}}

USAGE:
{{printf "\t%s %s %s" .Executable .Cmd.Name .Cmd.Usage}}

DESCRIPTION:
{{range $line := descToLines .Cmd.Description}}{{printf "\t%s" $line}}
{{end}}
{{if .CmdFlags}}OPTIONS:{{range .CmdFlags}}
{{printOption .Name .DefValue .Usage}}{{end}}

{{end}}For help on global options run "{{.Executable}} help"
`[1:]))
}

func runHelp(args []string) (exit int) {
	if len(args) < 1 {
		printGlobalUsage()
		return
	}

	var cmd *Command

	for _, c := range commands {
		if c.Name == args[0] {
			cmd = c
			break
		}
	}

	if cmd == nil {
		stderr("Unrecognized command: %s", args[0])
		return 1
	}

	printCommandUsage(cmd)
	return
}

func printGlobalUsage() {
	globalUsageTemplate.Execute(out, struct {
		Executable	string
		EnvFlag		string
		Commands	[]*Command
		Flags		[]*flag.Flag
		Description	string
		Version		string
	}{
		cliName,
		cliEnvFlag,
		commands,
		getAllFlags(),
		cliDescription,
		Version,
	})
	out.Flush()
}

func printCommandUsage(cmd *Command) {
	commandUsageTemplate.Execute(out, struct {
		Executable	string
		Cmd			*Command
		CmdFlags	[]*flag.Flag
	}{
		cliName,
		cmd,
		getFlags(&cmd.Flags),
	})
	out.Flush()
}
