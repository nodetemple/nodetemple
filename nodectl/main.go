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
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/nodetemple/nodetemple/common"
)

const (
	cliName        = "nodectl"
	cliDescription = "nodectl is a command-line interface for an orchestration of CoreOS and Kubernetes cluster"
)

var (
	tabOut        *tabwriter.Writer
	globalFlags   = struct {
		Debug       bool
		Help        bool
		Provider    string
	}{}

	cmdExitCode int

	cmdNodectl = &cobra.Command{
		Use:   fmt.Sprintf("%s [command]", cliName),
		Short: cliDescription,
	}
)

func init() {
	cmdNodectl.PersistentFlags().BoolVarP(&cmdNodectl.helpFlagVal, "help", "h", false, "Help for "+cmdNodectl.Name())
	cmdNodectl.PersistentFlags().BoolVar(&globalFlags.Debug, "debug", envBool("debug", false), "Print out more debug information")
	cmdNodectl.PersistentFlags().StringVarP(&globalFlags.Provider, "provider", "p", envString("provider", common.DefaultProvider), "Provider to use when managing a cluster")

	tabOut = new(tabwriter.Writer)
	tabOut.Init(os.Stdout, 0, 8, 1, '\t', 0)

	cobra.EnablePrefixMatching = true
}

func main() {
	cmdNodectl.SetUsageFunc(usageFunc)
	cmdNodectl.SetUsageTemplate(`{{.UsageString}}`)

	cmdNodectl.SetHelpCommand(cmdHelp)
	cmdNodectl.SetHelpTemplate(`{{.UsageString}}`)

	cmdNodectl.Execute()
	os.Exit(cmdExitCode)
}

func stderr(format string, a ...interface{}) {
	out := fmt.Sprintf(format, a...)
	fmt.Fprintln(os.Stderr, strings.TrimSuffix(out, "\n"))
}

func stdout(format string, a ...interface{}) {
	out := fmt.Sprintf(format, a...)
	fmt.Fprintln(os.Stdout, strings.TrimSuffix(out, "\n"))
}

func runWrapper(cf func(cmd *cobra.Command, args []string) (exit int)) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		cmdExitCode = cf(cmd, args)
	}
}

func envString(key, def string) string {
	envKey := strings.ToUpper(cliName + "_" + strings.Replace(key, "-", "_", -1))

	if env := os.Getenv(envKey); env != "" {
		return env
	}
	return def
}

func envBool(key string, def bool) bool {
	envKey := strings.ToUpper(cliName + "_" + strings.Replace(key, "-", "_", -1))

	if env := os.Getenv(envKey); env != "" {
		val, err := strconv.ParseBool(env)
		if err != nil {
			stderr("invalid value %q for %q (default: %t): %v", env, key, def, err)
			return def
		}
		return val
	}
	return def
}

func envInt(key string, def int) int {
	envKey := strings.ToUpper(cliName + "_" + strings.Replace(key, "-", "_", -1))

	if env := os.Getenv(envKey); env != "" {
		val, err := strconv.Atoi(env)
		if err != nil {
			stderr("invalid value %q for %q (default: %q): %v", env, key, def, err)
			return def
		}
		return val
	}
	return def
}
