package main

import (
	"flag"
	"fmt"
	"strings"
	"os"
	"text/tabwriter"
)

const (
	cliName        = "nodectl"
	cliDescription = "nodectl is a command-line interface for an orchestration of CoreOS and Kubernetes cluster"
)

var (
	out           *tabwriter.Writer
	commands      []*Command
	globalFlagset = flag.NewFlagSet("nodectl", flag.ExitOnError)
	globalFlags   = struct {
		Help        bool
		Version     bool
		SSHUserName string
	}{}
)

type Command struct {
	Name        string
	Description string
	Summary     string
	Usage       string
	Flags       flag.FlagSet
	Run         func(args []string) int
}

func init() {
	globalFlagset.BoolVar(&globalFlags.Help, "help", false, "Print usage information and exit")
	globalFlagset.BoolVar(&globalFlags.Help, "h", false, "Print usage information and exit")
	globalFlagset.BoolVar(&globalFlags.Version, "version", false, "Print the version and exit")
	globalFlagset.StringVar(&globalFlags.SSHUserName, "ssh-username", "core", "Username to use when connecting to CoreOS instance")
	
	out = new(tabwriter.Writer)
	out.Init(os.Stdout, 0, 8, 1, '\t', 0)
	
	commands = []*Command {
		cmdHelp,
		cmdVersion,
	}
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

func getFlags(flagset *flag.FlagSet) (flags []*flag.Flag) {
	flags = make([]*flag.Flag, 0)
	flagset.VisitAll(func(f *flag.Flag) {
		flags = append(flags, f)
	})
	return
}

func getAllFlags() (flags []*flag.Flag) {
	return getFlags(globalFlagset)
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
