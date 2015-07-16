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

package util

import (
	"log"
	"os"
)

var (
	outLogger = log.New(os.Stdout, "", 0)
	errLogger = log.New(os.Stderr, "Error: ", 0)
)

func Out(format string, a ...interface{}) {
	outLogger.Printf(format, a...)
}

func Err(format string, a ...interface{}) {
	errLogger.Fatalf(format, a...)
}
