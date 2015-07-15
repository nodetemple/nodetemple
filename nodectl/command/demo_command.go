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
		Flags: []cli.Flag{
			cli.StringFlag{Name: "demo-flag, d", Value: "", Usage: "demo flag usage"},
			cli.BoolFlag{Name: "demo-bool", Value: true, Usage: "demo bool usage"},
		},
		Action: func(c *cli.Context) {
			println("result:", c.Args().First(), c.GlobalString("provider"), c.String("demo-flag"), c.String("demo-bool"))
		},
	}
}
