package settings

var (
	ExtFile []string
	FileName string
	Backup bool
	BackupName string
	MapCharset string
	Folder string
)

func init() {
	FileName = ""
	ExtFile = []string{".php", ".html", ".htm", ".js", ".asp", ".tpl", ".txt"}
	Backup = true
	BackupName = ".c7bkp"
	MapCharset = ".charset"
	Folder = "./"
}

