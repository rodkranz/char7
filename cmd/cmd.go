package cmd

import (
	"gopkg.in/urfave/cli.v2"

	"github.com/rodkranz/char7/modules/setting"
)

func stringFlag(name, value, usage string) *cli.StringFlag {
	return &cli.StringFlag{
		Name:  name,
		Value: value,
		Usage: usage,
	}
}

func boolFlag(name, usage string) *cli.BoolFlag {
	return &cli.BoolFlag{
		Name:  name,
		Usage: usage,
	}
}

func parseFlags(ctx *cli.Context) {
	if ctx.IsSet("dir") {
		setting.Dir = ctx.String("dir")
	}

	if ctx.IsSet("backupName") {
		setting.BackupName = ctx.String("backupName")
	}

	if ctx.IsSet("map") {
		setting.MapCharset = ctx.String("map")
	}

	if ctx.IsSet("file") {
		setting.FileName = ctx.String("file")
	}
}
