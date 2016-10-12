package cmd

import (
	"time"

	"github.com/urfave/cli"

	"github.com/rodkranz/char7/modules/settings"
)

func stringFlag(name, value, usage string) cli.StringFlag {
	return cli.StringFlag{
		Name:  name,
		Value: value,
		Usage: usage,
	}
}

func boolFlag(name, usage string) cli.BoolFlag {
	return cli.BoolFlag{
		Name:  name,
		Usage: usage,
	}
}

func intFlag(name string, value int, usage string) cli.IntFlag {
	return cli.IntFlag{
		Name:  name,
		Value: value,
		Usage: usage,
	}
}

func durationFlag(name string, value time.Duration, usage string) cli.DurationFlag {
	return cli.DurationFlag{
		Name:  name,
		Value: value,
		Usage: usage,
	}
}

func parseFlags(ctx *cli.Context) {

	if ctx.IsSet("dir") {
		settings.Dir = ctx.String("dir")
	}

	if ctx.IsSet("backupName") {
		settings.BackupName = ctx.String("backupName")
	}

	if ctx.IsSet("map") {
		settings.MapCharset = ctx.String("map")
	}

	if ctx.IsSet("file") {
		settings.FileName = ctx.String("file")
	}
}
