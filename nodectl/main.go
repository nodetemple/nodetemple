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

	"github.com/nodetemple/cli"
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
		cli.BoolFlag{Name: "debug", Usage: "Print out more debug information to stderr"},
		cli.StringFlag{Name: "provider, p", Value: common.DefaultProvider, Usage: "Provider to use when managing a cluster", EnvVar: util.EnvVarConv(app.Name, "provider"),},
		cli.StringFlag{Name: "provider-key, k", Value: "", Usage: "Provider's API key to use when managing a cluster", EnvVar: util.EnvVarConv(app.Name, "provider-key"),},
	}
	app.Commands = []cli.Command{
		command.DemoCmd(),
	}
	app.CommandNotFound = func(c *cli.Context, command string) {
		util.Err("unknown command '%v'\nRun '%v help [command]' for usage information", command, c.App.Name)
	}

	app.Run(os.Args)
}
