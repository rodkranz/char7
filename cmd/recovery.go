package cmd

import (
	"os"

	"gopkg.in/urfave/cli.v2"

	"github.com/rodkranz/char7/modules/files"
	"github.com/rodkranz/char7/modules/logger"
	"github.com/rodkranz/char7/modules/setting"
)

// Recovery is command for recovery any data transformed for application
var Recovery = &cli.Command{
	Name:        "recovery",
	Usage:       "recovery backup",
	Description: `restore backup files and overwride the file wripped.`,
	Action:      runRecovery,
	Flags: []cli.Flag{
		stringFlag("dir, d", setting.Dir, "Define the folder that will looking for backup's file."),
		stringFlag("backupName, bn", setting.BackupName, "Search files name with extension backup c7."),
		boolFlag("remove, r", "Delete the backup file when finish."),
	},
}

func runRecovery(ctx *cli.Context) error {

	// Parse common information
	parseFlags(ctx)

	// filter to find files
	optFilter := &files.Filter{
		Exts: []string{setting.BackupName},
		Dir:  setting.Dir,
	}

	paths, err := files.SearchFiles(optFilter)
	if err != nil {
		return err
	}

	total := 0
	for _, backupPath := range paths {
		if (len(backupPath) - len(setting.BackupName)) < 0 {
			continue
		}

		logger.Restore(backupPath)
		restoreName := backupPath[:len(backupPath)-len(setting.BackupName)]
		if err := files.Copy(backupPath, restoreName); err != nil {
			logger.Fail()
			continue
		}
		logger.Success()
		total++

		if ctx.IsSet("remove") {
			logger.RemoveBackup(backupPath)
			if err := os.Remove(backupPath); err != nil {
				logger.Fail()
				continue
			}
			logger.Success()
		}
	}

	logger.InformationBackup(len(paths), total)
	return nil
}
