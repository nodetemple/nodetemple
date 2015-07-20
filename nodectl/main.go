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
	"flag"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/nodetemple/nodetemple/common"
)

const (
	OK = iota
	ERROR_API
	ERROR_USAGE
	ERROR_NO_COMMAND

	cliName        = "nodectl"
	cliDescription = "CLI for an orchestration of CoreOS and Kubernetes cluster"
)

type StringFlag struct {
	value    *string
	required bool
}

func (f *StringFlag) Set(value string) error {
	f.value = &value
	return nil
}

func (f *StringFlag) Get() *string {
	return f.value
}

func (f *StringFlag) String() string {
	if f.value != nil {
		return *f.value
	}
	return ""
}

type Command struct {
	Name        string       // Name of the Command and the string to use to invoke it
	Summary     string       // One-sentence summary of what the Command does
	Usage       string       // Usage options/arguments
	Description string       // Detailed description of command
	Flags       flag.FlagSet // Set of flags associated with this command
	Run         handlerFunc  // Run a command with the given arguments
	Subcommands []*Command   // Subcommands for this command.
}

var (
	out           *tabwriter.Writer
	globalFlagSet *flag.FlagSet
	commands      []*Command

	globalFlags struct {
		Debug         bool
		Version       bool
		Help          bool
		Server        string
		Key           string
	}
)

func init() {
	out = new(tabwriter.Writer)
	out.Init(os.Stdout, 0, 8, 1, '\t', 0)

	server := "http://localhost:8000" // default server
	if serverEnv := os.Getenv("UPDATECTL_SERVER"); serverEnv != "" {
		server = serverEnv
	}

	globalFlagSet = flag.NewFlagSet(cliName, flag.ExitOnError)
	globalFlagSet.BoolVar(&globalFlags.Debug, "debug", false, "Output debugging info to stderr")
	globalFlagSet.BoolVar(&globalFlags.Version, "version", false, "Print version information and exit.")
	globalFlagSet.BoolVar(&globalFlags.Help, "help", false, "Print usage information and exit.")
	globalFlagSet.StringVar(&globalFlags.Server, "server", server, "Update server to connect to")
	globalFlagSet.StringVar(&globalFlags.Key, "key", os.Getenv("NODECTL_KEY"), "API Key")

	commands = []*Command{
		//cmdApp,
		cmdHelp,
	}
}

type handlerFunc func([]string, *tabwriter.Writer) int

func handle(fn handlerFunc) func(f *flag.FlagSet) int {
	return func(f *flag.FlagSet) (exit int) {
		key := globalFlags.Key

		exit = OK
		return
	}
}

func printVersion(out *tabwriter.Writer) {
	fmt.Fprintf(out, "%s version %s\n", cliName, common.Version)
	out.Flush()
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

// determine which Command should be run
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

func main() {
	globalFlagSet.Parse(os.Args[1:])
	var args = globalFlagSet.Args()

	if globalFlags.Version {
		printVersion(out)
		os.Exit(OK)
	}

	if globalFlags.Help {
		printGlobalUsage()
		os.Exit(OK)
	}

	// no command specified - trigger help
	if len(args) < 1 {
		args = append(args, "help")
	}

	// trim the right most slash because all other uses of globalFlags.Server
	// append the / already
	globalFlags.Server = strings.TrimRight(globalFlags.Server, "/")

	cmd, name := findCommand("", args, commands)

	if cmd == nil {
		fmt.Printf("%v: unknown subcommand: %q\n", cliName, name)
		fmt.Printf("Run '%v help' for usage.\n", cliName)
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
