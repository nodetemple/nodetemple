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

	"github.com/codegangsta/cli"
	"github.com/nodetemple/nodetemple/common"
	"github.com/nodetemple/nodetemple/nodectl/util"
	"github.com/nodetemple/nodetemple/nodectl/command"
)

func main() {
	app := cli.NewApp()
	app.Name = "nodectl"
	app.Usage = "CLI for an orchestration of CoreOS and Kubernetes cluster"
	app.Version = common.Version
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "providers, p", Usage: "A comma-separated list of IaaS providers ("+strings.Join(common.AvailableProviders, ",")+") and API keys, format: 'provider:api-key,...'", EnvVar: util.EnvVarConv(app.Name, "providers"),},
		cli.BoolFlag{Name: "debug", Usage: "Print out more debug information to stderr"},
	}
	app.Before = func(c *cli.Context) {
		if c.String("providers") == "" && !c.Bool("help") && !c.Bool("version") {
			util.Err("set at least one provider with a valid API key")
		}
	}
	app.Commands = []cli.Command{
		command.DemoCmd(),
	}
	app.CommandNotFound = func(c *cli.Context, command string) {
		util.Err("unknown command '%v'\nRun '%v help [command]' for usage information", command, c.App.Name)
	}

	app.RunAndExitOnError()
}
