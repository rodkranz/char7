package settings

import (
	"log"
	"os/user"
)

var (
	// HomeDir is the user dir
	HomeDir string
	// ExtFile is the extensions allowed to replace
	ExtFile []string
	// FileName specify only one file to change
	FileName string
	// Backup indicate if has backup or not
	Backup bool
	// BackupName indicate backup's name
	BackupName string
	// MapCharset is the file path where has the map config.
	MapCharset string
	// Dir is the directory where the app will looking for files.
	Dir string
)

// WorkDir return the current user's dir
func WorkDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir
}

func init() {
	HomeDir = WorkDir()
	FileName = ""
	ExtFile = []string{".php", ".html", ".htm", ".js", ".asp", ".tpl", ".txt", ".srt"}
	Backup = true
	BackupName = ".c7"
	MapCharset = ".c7map"
	Dir = "./"
}
