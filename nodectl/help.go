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
	"strings"
	"text/template"

	"github.com/nodetemple/nodetemple/version"
)

var (
	commandUsageTemplate *template.Template
	templFuncs           = template.FuncMap{
		"descToLines": func(s string) []string {
			return strings.Split(strings.Trim(s, "\n\t "), "\n")
		},
		"cmdName": func(cmd *cobra.Command, startCmd *cobra.Command) string {
			parts := []string{cmd.Name()}
			for cmd.HasParent() && cmd.Parent().Name() != startCmd.Name() {
				cmd = cmd.Parent()
				parts = append([]string{cmd.Name()}, parts...)
			}
			return strings.Join(parts, " ")
		},
	}
)

func init() {
	commandUsage := `
{{ $cmd := .Cmd }}\
{{ $cmdname := cmdName .Cmd .Cmd.Root }}\
NAME:
{{ if not .Cmd.HasParent }}\
{{printf "\t%s - %s" .Cmd.Name .Cmd.Short}}
{{else}}\
{{printf "\t%s - %s" $cmdname .Cmd.Short}}
{{end}}\

USAGE:
{{printf "\t%s" .Cmd.UseLine}}
{{ if not .Cmd.HasParent }}\

VERSION:
{{printf "\t%s" .Version}}
{{end}}\
{{if .Cmd.HasSubCommands}}\

COMMANDS:
{{range .SubCommands}}\
{{ $cmdname := cmdName . $cmd }}\
{{ if .Runnable }}\
{{printf "\t%s\t%s" $cmdname .Short}}
{{end}}\
{{end}}\
{{end}}\
{{ if .Cmd.Long }}\

DESCRIPTION:
{{range $line := descToLines .Cmd.Long}}{{printf "\t%s" $line}}
{{end}}\
{{end}}\
{{if .Cmd.HasLocalFlags}}\

OPTIONS:
{{.Cmd.LocalFlags.FlagUsages}}\
{{end}}\
{{if .Cmd.HasInheritedFlags}}\

GLOBAL OPTIONS:
{{.Cmd.InheritedFlags.FlagUsages}}

Global options can also be configured via upper-case environment variables prefixed with "{{.EnvFlag}}_".
For example: "--some-flag" => "{{.EnvFlag}}_SOME_FLAG".\
{{end}}
`[1:]

	commandUsageTemplate = template.Must(template.New("command_usage").Funcs(templFuncs).Parse(strings.Replace(commandUsage, "\\\n", "", -1)))
}

func getSubCommands(cmd *cobra.Command) []*cobra.Command {
	subCommands := []*cobra.Command{}
	for _, subCmd := range cmd.Commands() {
		subCommands = append(subCommands, subCmd)
		subCommands = append(subCommands, getSubCommands(subCmd)...)
	}
	return subCommands
}

func usageFunc(cmd *cobra.Command) error {
	subCommands := getSubCommands(cmd)
	commandUsageTemplate.Execute(tabOut, struct {
		Cmd         *cobra.Command
		SubCommands []*cobra.Command
		EnvFlag     string
		Version     string
	}{
		cmd,
		subCommands,
		strings.ToUpper(cliName),
		version.Version,
	})
	tabOut.Flush()
	return nil
}
