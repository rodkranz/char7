package main

import (
	"os"
	"fmt"
	"runtime"

	"github.com/urfave/cli"

	"github.com/rodkranz/char7/cmd"
	_ "github.com/rodkranz/char7/modules/settings"
	_ "github.com/rodkranz/char7/modules/mapping"
)

const VER string = "v1.0.0"

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {

	app := cli.NewApp()
	app.Name = "CharSet"
	app.Usage = "CharSet convert"
	app.Author = "Kranz"
	app.Email = "kranz@null.net"
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
