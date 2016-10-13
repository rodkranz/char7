package settings

import (
	"log"
	"os/user"
)

var (
	HomeDir    string
	ExtFile    []string
	FileName   string
	Backup     bool
	BackupName string
	MapCharset string
	Dir        string

)

func WorkDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal( err )
	}
	return usr.HomeDir
}

func init() {
	HomeDir = WorkDir()
	FileName = ""
	ExtFile = []string{".php", ".html", ".htm", ".js", ".asp", ".tpl", ".txt"}
	Backup = true
	BackupName = ".c7"
	MapCharset = ".c7map"
	Dir = "./"
}
