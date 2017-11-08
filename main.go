// +build go1.9
package main

import (
	"fmt"
	"os"
	"runtime"

	"gopkg.in/urfave/cli.v2"

	"github.com/rodkranz/char7/cmd"

	_ "github.com/rodkranz/char7/modules/mapping"
	_ "github.com/rodkranz/char7/modules/settings"
)

// AppName is the application name
const AppName string = "char7"

// AppVer is the current version of application
const AppVer string = "v1.0.0"

func init() {
	// Allow to use all core of computer
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	app := cli.App{
		Name:        AppName,
		Usage:       "CharSet convert",
		Description: "Replace Utf8 encoding to Ascii code",
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "Rodrigo Kranz",
				Email: "kranz@null.net",
			},
		},
		Version: AppVer,
		Commands: []*cli.Command{
			cmd.CmdCharSet,
			cmd.CmdRecovery,
		},
	}

	app.Flags = append(app.Flags, []cli.Flag{}...)
	err := app.Run(os.Args)

	if err != nil {
		fmt.Fprintf(app.Writer, "Error: %s\n", err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}
