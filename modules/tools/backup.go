package tools

import (
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/rodkranz/char7/modules/settings"
)

func GenBackupNameFor(path string) (string) {
	count := 1
	bkp := path + settings.BackupName
	for {
		if !Exists(bkp) {
			return bkp
		}

		bkp = fmt.Sprintf("%v%v.%d", path, settings.BackupName, count)
		count++

		if count > 1000 {
			return uuid.NewV4().String()
		}
	}
}
