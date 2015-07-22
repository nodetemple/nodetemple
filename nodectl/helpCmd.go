/*
Copyright 2015 Nodetemple <hostmaster@nodetemple.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"strings"
	"text/template"

	flag "github.com/ogier/pflag"
)

var (
	cmdHelp = &Command{
		Name:        "help",
		Usage:       "<command>",
		Summary:     "Show a list of commands or help for one command",
		Description: "Show a list of commands or detailed help for one command",
		Run:         runHelp,
	}

	globalUsageTemplate  *template.Template
	commandUsageTemplate *template.Template
	templFuncs           = template.FuncMap{
		"descToLines": func(s string) []string {
			return strings.Split(strings.Trim(s, "\n\t "), "\n")
		},
		"printFlag": func(shorthand, name, defvalue, usage string) string {
			format := "--%s=%s\t%s"
			/*if _, ok := flag.Value.(*stringValue); ok {
				format = "--%s=%q\t%s"
			}*/
			if len(shorthand) > 0 {
				format = "    -%s, " + format
			} else {
				format = "     %s   " + format
			}
			return fmt.Sprintf(format, shorthand, name, defvalue, usage)
		},
	}
)

func init() {
	globalUsageTemplate = template.Must(template.New("global_usage").Funcs(templFuncs).Parse(`
NAME:
{{printf "\t%s - %s" .Executable .Description}}

USAGE:
{{printf "\t%s" .Executable}} [global flags] <command> [command flags] [arguments...]

VERSION:
{{printf "\t%s" .Version}}

COMMANDS:{{range .Commands}}
{{printf "\t%s\t%s" .Name .Summary}}{{end}}

GLOBAL FLAGS:{{range .Flags}}
{{printFlag .Shorthand .Name .DefValue .Usage}}{{end}}

Global flags can also be configured via upper-case environment variables prefixed with '{{.ExeEnvPrefix}}_'
For example: '--some-flag' => '{{.ExeEnvPrefix}}_SOME_FLAG'

Run '{{.Executable}} help <command>' for more details on a specific command
`[1:]))
	commandUsageTemplate = template.Must(template.New("command_usage").Funcs(templFuncs).Parse(`
NAME:
{{printf "\t%s - %s" .Cmd.Name .Cmd.Summary}}

USAGE:
{{printf "\t%s %s %s" .Executable .Cmd.Name .Cmd.Usage}}

{{if .Cmd.Description}}DESCRIPTION:{{range $line := descToLines .Cmd.Description}}
{{printf "\t%s" $line}}{{end}}

{{end}}{{if .Cmd.Subcommands}}COMMANDS:{{range .Cmd.Subcommands}}
{{printf "\t%s\t%s" .Name .Summary}}{{end}}

{{end}}{{if .CmdFlags}}FLAGS:{{range .CmdFlags}}
{{printFlag .Shorthand .Name .DefValue .Usage}}{{end}}

{{end}}For help on global flags run '{{.Executable}} help'
`[1:]))
}

func runHelp(args []string) int {
	if len(args) < 1 {
		printGlobalUsage()
		return OK
	}

	var cmd *Command

	for _, c := range commands {
		if c.Name == args[0] {
			cmd = c
			break
		}
	}

	if cmd == nil {
		stderr("Error: unknown command '%s'\n", args[0])
		stderr("Run '%s help' for usage information\n", cliName)
		return ERROR_NO_COMMAND
	}

	printCommandUsage(cmd)
	return OK
}

func printGlobalUsage() {
	globalUsageTemplate.Execute(out, struct {
		Executable  string
		ExeEnvPrefix string
		Commands    []*Command
		Flags       []*flag.Flag
		Description string
		Version     string
	}{
		cliName,
		strings.ToUpper(cliName),
		commands,
		getAllFlags(),
		cliDescription,
		cliVersion,
	})
}

func printCommandUsage(cmd *Command) {
	commandUsageTemplate.Execute(out, struct {
		Executable   string
		Cmd          *Command
		CmdFlags     []*flag.Flag
	}{
		cliName,
		cmd,
		getFlags(&cmd.Flags),
	})
}

func getAllFlags() (flags []*flag.Flag) {
	return getFlags(globalFlagSet)
}

func getFlags(flagset *flag.FlagSet) (flags []*flag.Flag) {
	flags = make([]*flag.Flag, 0)
	flagset.VisitAll(func(f *flag.Flag) {
		flags = append(flags, f)
	})
	return
}
