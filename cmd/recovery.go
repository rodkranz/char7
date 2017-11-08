package cmd

import (
	"fmt"

	"gopkg.in/urfave/cli.v2"

	"github.com/rodkranz/char7/modules/files"
	"github.com/rodkranz/char7/modules/settings"
)

// Recovery is command for recovery any data transformed for application
var Recovery = &cli.Command{
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

	// Parse common information
	parseFlags(ctx)

	// filter to find files
	optFilter := &files.Filter{
		Exts: []string{settings.BackupName},
		Dir:  settings.Dir,
	}

	paths, err := files.SearchFiles(optFilter)
	if err != nil {
		return err
	}

	total := 0
	for _, backupPath := range paths {
		if (len(backupPath) - len(settings.BackupName)) < 0 {
			continue
		}

		fmt.Fprintf(ctx.App.Writer, "Restouring %s...   ", backupPath)
		restoreName := backupPath[:len(backupPath)-len(settings.BackupName)]
		if err := files.Copy(backupPath, restoreName); err != nil {
			fmt.Fprintln(ctx.App.Writer, "[FAIL]")
			continue
		}
		fmt.Fprintln(ctx.App.Writer, "[SUCCESS]")
		total++

		if ctx.IsSet("remove") {
			fmt.Fprintf(ctx.App.Writer, "Removing backup %s...   ", backupPath)
			if err := files.Delete(backupPath); err != nil {
				fmt.Fprintln(ctx.App.Writer, "[FAIL]")
				continue
			}
			fmt.Fprintln(ctx.App.Writer, "[SUCCESS]")
		}
	}

	fmt.Fprintf(ctx.App.Writer, "Found %d backups, restoured %d files.\n", len(paths), total)
	return nil
}
