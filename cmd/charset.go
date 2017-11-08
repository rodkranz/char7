package cmd

import (
	"strings"
	
	"gopkg.in/urfave/cli.v2"
	
	"github.com/rodkranz/char7/modules/files"
	"github.com/rodkranz/char7/modules/settings"
	"github.com/rodkranz/char7/modules/logger"
	"github.com/rodkranz/char7/modules/charset"
)

// CharSet is the main command for application
// define information about the program as parameters
var CharSet = &cli.Command{
	Name:        "charset",
	Usage:       "Change charset to html cod",
	Description: `change key in map to value`,
	Action:      runCharSet,
	Flags: []cli.Flag{
		stringFlag("ext, e", strings.Join(settings.ExtFile, ","), "Define the extension type to filter files"),
		stringFlag("file, f", settings.FileName, "Define the file name that needs to change."),
		stringFlag("dir, d", settings.Dir, "Define the folder that will looking for files."),
		boolFlag("backup, b", "Disable backup file when change charset"),
		stringFlag("backupName, bn", settings.BackupName, "Define the file name that needs to change."),
	},
}

// runCharSet is the funcation that make the magic
// find files with filter defined by user and resturn the files.
func runCharSet(ctx *cli.Context) error {
	
	// Parse common information
	parseFlags(ctx)
	
	// split extensions
	if ctx.IsSet("ext") {
		settings.ExtFile = strings.Split(ctx.String("ext"), ",")
	}
	
	// filter to find files
	optFilter := &files.Filter{
		FileName: settings.FileName,
		Exts:     settings.ExtFile,
		Dir:      settings.Dir,
	}
	
	// search files using the filter defined by user
	list, err := files.SearchFiles(optFilter)
	if err != nil {
		return err
	}
	
	var total int
	for _, path := range list {
		bkpPath := path + settings.BackupName
		if !ctx.IsSet("backup") {
			logger.Copy(path, bkpPath)
			if e := files.Copy(path, bkpPath); e != nil {
				logger.Fail()
				continue
			}
			logger.Success()
		}
		
		logger.Convert(path)
		if e := charset.CharSet(path); e != nil {
			logger.Fail()
			continue
		}
		logger.Success()
		
		//
		if !charset.HasChange {
			total++
			if !ctx.IsSet("backup") {
				logger.Delete(path)
				if err := files.Delete(bkpPath); err != nil {
					logger.Fail()
					continue
				}
				logger.Success()
			}
		}
	}
	
	logger.Information(len(list), total)
	return nil
}
