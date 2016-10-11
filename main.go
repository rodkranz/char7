package main

import (
	"runtime"

	"bitbucket.org/rkranz/tmp/cmd"
	_ "bitbucket.org/rkranz/tmp/modules/settings"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	cmd.Run()
}
