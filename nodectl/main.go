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

	"github.com/codegangsta/cli"
	"github.com/nodetemple/nodetemple/common"
)

func main() {
	app := cli.NewApp()
	app.Name = "nodectl"
	app.Usage = "CLI for an orchestration of CoreOS and Kubernetes cluster"
	app.Version = common.Version
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "providers, p", Usage: "A comma-separated list of IaaS providers ("+strings.Join(common.AvailableProviders, ",")+") and API keys, format: 'provider:api-key,...'", EnvVar: envVarConv(app.Name, "providers"),},
	}
	app.Commands = []cli.Command{
		demoCmd,
	}
	app.CommandNotFound = cmdNotFound

	if err := app.Run(os.Args); err != nil {
		stderr(err)
		os.Exit(1)
	}
}

func stdout(format string, a ...interface{}) {
	out := fmt.Sprintf(format, a...)
	fmt.Fprintln(os.Stdout, strings.TrimSuffix(out, "\n"))
}

func stderr(format string, a ...interface{}) {
	out := fmt.Sprintf(format, a...)
	fmt.Fprintln(os.Stderr, strings.TrimSuffix(out, "\n"))
}

func cmdNotFound(c *cli.Context, command string) error {
	fmt.Errorf("unknown command '%v'\nRun '%v help [command]' for usage information", command, c.App.Name)
}
