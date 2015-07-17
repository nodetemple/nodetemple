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

package command

import (
	"github.com/codegangsta/cli"
	"github.com/nodetemple/nodetemple/nodectl/util"
)

func DemoCmd() cli.Command {
	return cli.Command{
		Name:  "demo",
		Usage: "A simple `hello world` demo",
		Description: "A simple `hello world` demo with output of flags, args, etc.",
		/*Flags: []cli.Flag{
			cli.StringFlag{Name: "demo-flag, d", Value: "", Usage: "Demo flag usage"},
			cli.BoolFlag{Name: "demo-bool", Usage: "Demo bool usage"},
		},
		Action: demoCmdFunc,*/
		Subcommands: []cli.Command{
			{
				Name:  "add",
				Usage: "Add a new template",
				Action: func(c *cli.Context) {
					Action: demoCmdFunc
				},
			},
			{
				Name:  "remove",
				Usage: "Remove an existing template",
				Action: func(c *cli.Context) {
					Action: demoCmdFunc
				},
			},
		},
	}
}

func demoCmdFunc(c *cli.Context) {
	/*if c.String("demo-flag") == "" {
		util.Err("missing '--demo-flag'")
	}*/

	util.Out("Result:\n\targs: %v\n\tproviders: %v\n\tdemo-flag: %v\n\tdemo-bool: %v", c.Args().Get(0), c.GlobalStringSlice("providers"), c.String("demo-flag"), c.String("demo-bool"))
}
