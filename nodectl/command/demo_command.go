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
)

func DemoCommand() cli.Command{
	return cli.Command{
		Name:  "demo",
		Usage: "a simple `hello world` demo",
		/*Action: func(c *cli.Context) {
			println("Hello: ", c.Args().First())
		},*/
		Subcommands: []cli.Command{
			{
				Name:  "add",
				Usage: "add a new template",
				Flags: []cli.Flag{
					//cli.IntFlag{Name: "ttl", Value: 0, Usage: "key time-to-live"},
					cli.StringFlag{Name: "demo-flag, d", Value: "", Usage: "demo flag usage"},
				},
				Action: func(c *cli.Context) {
					println("new task template:", c.Args().First(), c.GlobalString("provider"), c.String("demo-flag"))
				},
			},
			{
				Name:  "remove",
				Usage: "remove an existing template",
				Flags: []cli.Flag{
					//cli.IntFlag{Name: "ttl", Value: 0, Usage: "key time-to-live"},
					cli.StringFlag{Name: "demo-flag, d", Value: "", Usage: "demo flag usage"},
				},
				Action: func(c *cli.Context) {
					println("removed task template:", c.Args().First(), c.GlobalString("provider"), c.String("demo-flag"))
				},
			},
		},
	}
}
