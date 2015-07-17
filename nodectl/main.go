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
	"strings"

	"github.com/codegangsta/cli"
	"github.com/nodetemple/nodetemple/common"
	"github.com/nodetemple/nodetemple/version"
	"github.com/nodetemple/nodetemple/nodectl/util"
	"github.com/nodetemple/nodetemple/nodectl/command"
)

func main() {
	app := cli.NewApp()
	app.Name = "nodectl"
	app.Usage = "CLI for an orchestration of CoreOS and Kubernetes cluster"
	app.Version = version.Version
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "providers, p", Usage: "A comma-separated list of providers ("+strings.Join(common.AvailableProviders, ", ")+") to use when managing a cluster, e.g.: do:54c234d6e7...", EnvVar: util.EnvVarConv(app.Name, "providers"),},
		cli.BoolFlag{Name: "debug", Usage: "Print out more debug information to stderr"},
	}
	/*app.Before = func(c *cli.Context) error {
		if c.String("api-key") != "" {
			APIKey = c.String("api-key")
		}

		if APIKey == "" && !c.Bool("help") && !c.Bool("version") {
			return errors.New("must provide API Key via NODECTL_PROVIDERS environment variable or via CLI argument.")
		}

		switch c.String("format") {
		case "json":
			OutputFormat = c.String("format")
		case "yaml":
			OutputFormat = c.String("format")
		default:
			return fmt.Errorf("invalid output format: %q, available output options: json, yaml.", c.String("format"))
		}

		return nil
	}*/
	app.Commands = []cli.Command{
		command.DemoCmd(),
	}
	app.CommandNotFound = func(c *cli.Context, command string) {
		util.Err("unknown command '%v'\nRun '%v help [command]' for usage information", command, c.App.Name)
	}

	app.Run(os.Args)
}
