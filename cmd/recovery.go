package cmd

import (
	"fmt"
	"github.com/urfave/cli"
	"github.com/rodkranz/char7/modules/files"
	"github.com/rodkranz/char7/modules/settings"
)

var CmdRecovery = cli.Command{
	Name:        "recovery",
	Usage:       "recovery backup",
	Description: `restore backup files and overwride the file wripped.`,
	Action:      runRecovery,
	Flags: []cli.Flag{
		stringFlag("dir, d", settings.Dir, "Define the folder that will looking for backup's file."),
		stringFlag("backupName, bn", settings.BackupName, "Search files name with extension backup c7."),
		boolFlag("remove, r", "Delete the backup file when finish."),
	},
}


func runRecovery(ctx *cli.Context) error {
	if ctx.IsSet("dir") {
		settings.Dir = ctx.String("dir")
	}

	if ctx.IsSet("backupName") {
		settings.BackupName = ctx.String("backupName")
	}

	optFilter := &files.Filter{
		Exts:     []string{settings.BackupName},
		Dir:      settings.Dir,
	}

	fmt.Printf("%v",optFilter)

	paths, err := files.SearchFiles(optFilter)
	if err != nil {
		return err
	}

	for _, backupPath := range paths {
		if (len(backupPath) - len(settings.BackupName)) < 0 {
			continue
		}

		restoreName := backupPath[:len(backupPath)-len(settings.BackupName)]
		if err := files.Copy(backupPath, restoreName); err != nil {
			return err
		}

		if ctx.IsSet("remove") {
			if err := files.Delete(backupPath); err != nil {
				return err
			}
		}
	}

	return nil
}
