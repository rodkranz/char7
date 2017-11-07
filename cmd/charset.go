package cmd

import (
	"fmt"
	"strings"

	"gopkg.in/urfave/cli.v2"

	"github.com/rodkranz/char7/modules/charset"
	"github.com/rodkranz/char7/modules/files"
	"github.com/rodkranz/char7/modules/settings"
)

// CmdCharSet is the main command for application
// define information about the program as parameters
var CmdCharSet = &cli.Command{
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

	var total int = 0
	for _, path := range list {
		bkpPath := path + settings.BackupName

		if !ctx.IsSet("backup") {
			fmt.Fprintf(ctx.App.Writer, "Copyng %s to %s...   ", path, bkpPath)
			if e := files.Copy(path, bkpPath); e != nil {
				fmt.Fprintln(ctx.App.Writer, "[FAIL]")
				continue
			}
			fmt.Fprintln(ctx.App.Writer, "[SUCCESS]")
		}

		fmt.Fprintf(ctx.App.Writer, "Finding chars to convert from %s..   ", path)
		if e := charset.CharSet(path); e != nil {
			fmt.Fprintln(ctx.App.Writer, "[FAIL]")
			continue
		}
		fmt.Fprintln(ctx.App.Writer, "[SUCCESS]")

		if !charset.HasChange {
			total++
			if !ctx.IsSet("backup") {
				fmt.Fprintf(ctx.App.Writer, "Deleting useless file %s..   ", path)
				if err := files.Delete(bkpPath); err != nil {
					fmt.Fprintln(ctx.App.Writer, "[FAIL]")
					continue
				}
				fmt.Fprintln(ctx.App.Writer, "[SUCCESS]")
			}
		}
	}

	fmt.Fprintf(ctx.App.Writer, "Found %d files, changed %d files.\n", len(list), total)
	return nil
}
