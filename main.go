package main

import (
	"runtime"

	"github.com/rodkranz/char7/cmd"
	_ "github.com/rodkranz/char7/modules/settings"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	cmd.Run()
}
