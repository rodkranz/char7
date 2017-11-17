// +build go1.9
package main

import (
	"log"
	"os"
	"runtime"

	"gopkg.in/urfave/cli.v2"

	"github.com/rodkranz/char7/cmd"
	"github.com/rodkranz/char7/modules/setting"
)

// AppName is the application name
const AppName = "char7"

// AppVer is the current version of application
const AppVer = "v1.0.0"

func init() {
	// Allow to use all core of computer
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	app := cli.App{
		Name:        AppName,
		Usage:       "CharSet convert",
		Description: "Replace Utf8 encoding to Ascii code",
		Authors: []*cli.Author{{
			Name:  "Rodrigo Kranz",
			Email: "kranz@null.net",
		}},
		Before:  setting.Bootstrap,
		Version: AppVer,
		Commands: []*cli.Command{
			cmd.CharSet,
			cmd.Recovery,
		},
	}

	app.Flags = append(app.Flags, []cli.Flag{}...)
	err := app.Run(os.Args)

	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}
