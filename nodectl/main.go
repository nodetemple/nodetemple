/* docker run -d --name=ubuntu-go ubuntu /bin/bash -c 'while true; do sleep 1; done'
 * docker exec -it ubuntu-go /bin/bash
 * apt-get update && apt-get install -y golang nano
 * mkdir -p /root/nodectl && cd /root/nodectl && export TERM=xterm
 * rm nodectl.go && nano nodectl.go
 * rm version.go && nano version.go
 * go build .
 * ./nodectl
 * ./nodectl -h
*/

package main

import (
	"flag"
	"fmt"
	"strings"
	"os"
	"text/tabwriter"
)

const (
	cliName			=	"nodectl"
	cliEnvFlag		=	"NODECTL"
	cliDescription	=	"nodectl is a command-line interface for a cluster-wide orchestration of CoreOS servers"
)

var (
	out				*tabwriter.Writer
	globalFlagset = flag.NewFlagSet("nodectl", flag.ExitOnError)
	
	commands []*Command
	
	globalFlags = struct {
		Version		bool
		Help		bool
		
		SSHUserName	string
	}{}
)

func init() {
	globalFlagset.BoolVar(&globalFlags.Help, "help", false, "Print usage information and exit")
	globalFlagset.BoolVar(&globalFlags.Help, "h", false, "Print usage information and exit")
	
	globalFlagset.BoolVar(&globalFlags.Version, "version", false, "Print the version and exit")
	
	globalFlagset.StringVar(&globalFlags.SSHUserName, "ssh-username", "core", "Username to use when connecting to CoreOS instance")
}

type Command struct {
	Name        string       // Name of the Command and the string to use to invoke it
	Summary     string       // One-sentence summary of what the Command does
	Usage       string       // Usage options/arguments
	Description string       // Detailed description of command
	Flags       flag.FlagSet // Set of flags associated with this command

	Run func(args []string) int // Run a command with the given arguments, return exit status
}

func init() {
	out = new(tabwriter.Writer)
	out.Init(os.Stdout, 0, 8, 1, '\t', 0)
	commands = []*Command{
		cmdHelp,
		cmdVersion,
	}
}

func getAllFlags() (flags []*flag.Flag) {
	return getFlags(globalFlagset)
}

func getFlags(flagset *flag.FlagSet) (flags []*flag.Flag) {
	flags = make([]*flag.Flag, 0)
	flagset.VisitAll(func(f *flag.Flag) {
		flags = append(flags, f)
	})
	return
}

func maybeAddNewline(s string) string {
	if !strings.HasSuffix(s, "\n") {
		s = s + "\n"
	}
	return s
}

func stderr(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, maybeAddNewline(format), args...)
}

func stdout(format string, args ...interface{}) {
	fmt.Fprintf(os.Stdout, maybeAddNewline(format), args...)
}

func main() {
    globalFlagset.Parse(os.Args[1:])

	args := globalFlagset.Args()

	getFlagsFromEnv(cliName, globalFlagset)
	
	if globalFlags.Version {
		args = []string{"version"}
	} else if len(args) < 1 || globalFlags.Help {
		args = []string{"help"}
	}
	
	var cmd *Command
	
	for _, c := range commands {
		if c.Name == args[0] {
			cmd = c
			if err := c.Flags.Parse(args[1:]); err != nil {
				stderr("%v", err)
				os.Exit(2)
			}
			break
		}
	}

	if cmd == nil {
		stderr("%v: unknown subcommand: %q", cliName, args[0])
		stderr("Run '%v help' for usage.", cliName)
		os.Exit(2)
	}
	
	visited := make(map[string]bool, 0)
	globalFlagset.Visit(func(f *flag.Flag) { visited[f.Name] = true })
	
	if cmd.Name != "help" && cmd.Name != "version" {
		/*var err error
		cAPI, err = getClient()
		if err != nil {
			stderr("Unable to initialize client: %v", err)
			os.Exit(1)
		}*/
	}

	os.Exit(cmd.Run(cmd.Flags.Args()))
}

func getFlagsFromEnv(prefix string, fs *flag.FlagSet) {
	alreadySet := make(map[string]bool)
	fs.Visit(func(f *flag.Flag) {
		alreadySet[f.Name] = true
	})
	fs.VisitAll(func(f *flag.Flag) {
		if !alreadySet[f.Name] {
			key := strings.ToUpper(prefix + "_" + strings.Replace(f.Name, "-", "_", -1))
			val := os.Getenv(key)
			if val != "" {
				fs.Set(f.Name, val)
			}
		}

	})
}
