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
	"os"
	"fmt"
	"strings"
	"text/tabwriter"

	"github.com/nodetemple/nodetemple/common"
	flag "github.com/ogier/pflag"
)

const (
	OK = iota
	ERROR_API
	ERROR_USAGE
	ERROR_NO_COMMAND

	cliName        = "nodectl"
	cliDescription = "CLI for an orchestration of CoreOS and Kubernetes cluster"
	cliVersion     = common.Version
)

type StringFlag struct {
	value    *string
	required bool
}

type Command struct {
	Name        string
	Usage       string
	Summary     string
	Description string
	Flags       flag.FlagSet
	Run         handlerFunc
	Subcommands []*Command
}

type handlerFunc func([]string) int

var (
	out           *tabwriter.Writer
	globalFlagSet *flag.FlagSet
	commands      []*Command

	globalFlags struct {
		Providers     string
		Debug         bool
		Version       bool
		Help          bool
	}
)

func init() {
	out = new(tabwriter.Writer)
	out.Init(os.Stdout, 0, 8, 1, '\t', 0)

	globalFlagSet = flag.NewFlagSet(cliName, flag.ExitOnError)
	globalFlagSet.Usage = func() {
		stderr("Error: incorrect flag usage")
		stderr("Run '%s help' for usage information", cliName)
	}

	globalFlagSet.StringVarP(&globalFlags.Providers, "providers", "p", "", "A comma-separated list of IaaS providers ("+strings.Join(common.AvailableProviders, ",")+") and API keys, format: 'provider:api-key,...'")

	globalFlagSet.BoolVar(&globalFlags.Debug, "debug", false, "Output debugging info to stderr")
	globalFlagSet.BoolVarP(&globalFlags.Version, "version", "v", false, "Print version information and exit")
	globalFlagSet.BoolVarP(&globalFlags.Help, "help", "h", false, "Print usage information and exit")

	commands = []*Command{
		cmdDemo,
		cmdHelp,
	}
}

func main() {
	globalFlagSet.Parse(os.Args[1:])
	var args = globalFlagSet.Args()
	getFlagsFromEnv(cliName, globalFlagSet)

	if globalFlags.Version {
		stdout("%s version %s", cliName, cliVersion)
		os.Exit(OK)
	}

	if globalFlags.Help {
		printGlobalUsage()
		os.Exit(OK)
	}

	if len(args) < 1 {
		args = append(args, "help")
	}

	cmd, name := findCommand("", args, commands)

	if cmd == nil {
		stderr("Error: unknown command '%s'", name)
		stderr("Run '%s help' for usage information", cliName)
		os.Exit(ERROR_NO_COMMAND)
	}

	if cmd.Run == nil {
		printCommandUsage(cmd)
		os.Exit(ERROR_USAGE)
	} else {
		exit := handle(cmd.Run)(&cmd.Flags)
		if exit == ERROR_USAGE {
			printCommandUsage(cmd)
		}
		os.Exit(exit)
	}
}

func handle(fn handlerFunc) func(f *flag.FlagSet) int {
	return func(f *flag.FlagSet) (exit int) {
		exit = fn(f.Args())
		return
	}
}

func findCommand(search string, args []string, commands []*Command) (cmd *Command, name string) {
	if len(args) < 1 {
		return
	}
	if search == "" {
		search = args[0]
	} else {
		search = fmt.Sprintf("%s %s", search, args[0])
	}
	name = search
	for _, c := range commands {
		if c.Name == search {
			cmd = c
			if errHelp := c.Flags.Parse(args[1:]); errHelp != nil {
				printCommandUsage(cmd)
				os.Exit(ERROR_USAGE)
			}
			if len(cmd.Subcommands) != 0 {
				subArgs := cmd.Flags.Args()
				var subCmd *Command
				subCmd, name = findCommand(search, subArgs, cmd.Subcommands)
				if subCmd != nil {
					cmd = subCmd
				}
			}
			break
		}
	}
	return
}
