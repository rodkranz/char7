package main

import (
	"runtime"

	"github.com/rodkranz/char7/cmd"
	"github.com/urfave/cli"

	_ "github.com/rodkranz/char7/modules/settings"
	"os"
	"fmt"
)

const VER string = "v1.0.0"

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {

	app := cli.NewApp()
	app.Name = "CharSet"
	app.Usage = "CharSet convert"
	app.Version = VER
	app.Commands = []cli.Command{
		cmd.CmdCharSet,
		cmd.CmdRecovery,
	}

	app.Flags = append(app.Flags, []cli.Flag{}...)
	err := app.Run(os.Args)

	if err != nil {
		fmt.Fprintf(app.Writer, "Error: %s\n", err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}