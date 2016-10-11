package cmd

import (
	"bitbucket.org/rkranz/tmp/modules/files"
	"bitbucket.org/rkranz/tmp/modules/charset"
	"bitbucket.org/rkranz/tmp/modules/settings"
)

func Run() error {
	list, err := files.SearchFiles()
	if err != nil {
		return err
	}

	for _, path := range list {
		a, b, e := charset.CharSet(path);
		if e != nil {
			return e
		}

		if b.Len() != 0 {
			bkpDst := path + settings.BackupName
			files.Write(path,   b.Bytes())
			files.Write(bkpDst, a.Bytes())
		}
	}
	return nil
}