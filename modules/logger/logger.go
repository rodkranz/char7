package logger

import (
	"fmt"
	
	"github.com/fatih/color"
	"path"
)

const (
	file    = "cyan"
	info    = "blue"
	success = "green"
	fail    = "red"
	number  = "yellow"
	
	MaxNum = 30
)

var colors = map[string]*color.Color{
	"blue":   color.New(color.FgBlue),
	"red":    color.New(color.FgRed),
	"yellow": color.New(color.FgYellow),
	"green":  color.New(color.FgGreen),
	"cyan":   color.New(color.FgCyan),
}

func Information(nFiles, nSuccess int) {
	fmt.Printf(
		"%s %s %s %s %s\n",
		colors[info].Sprint("Found"),
		colors[number].Sprint(nFiles),
		colors[info].Sprint("file, changed"),
		colors[number].Sprint(nSuccess),
		colors[info].Sprint("file"),
	)
}

func Delete(fileName string) {
	fmt.Printf(
		"%s [%s]%v ",
		colors[info].Sprint("Deleting useless file"),
		colors[file].Sprint(LimitChar(fileName)),
		colors[info].Sprint("..."),
	)
}

func Convert(fileName string) {
	fmt.Printf(
		"%s [%s]%s ",
		colors[info].Sprint("Finding chars to convert from"),
		colors[file].Sprint(LimitChar(fileName)),
		colors[info].Sprint("..."),
	)
}

func Copy(fileNameA, fileNameB string) {
	// Create a new color object
	fmt.Printf(
		"%s [%s] %s [%s]%s ",
		colors[info].Sprint("Copying"),
		colors[file].Sprint(LimitChar(fileNameA)),
		colors[info].Sprint("to"),
		colors[file].Sprint(LimitChar(fileNameB)),
		colors[info].Sprint("..."),
	)
}

func Success() {
	fmt.Printf(
		"[%s]\n",
		colors[success].Sprint("SUCCESS"),
	)
}

func Fail() {
	fmt.Printf(
		"[%s]\n",
		colors[fail].Sprint("FAIL"),
	)
}

func LimitChar(s string) string {
	s = path.Base(s)
	
	if len(s) <= MaxNum {
		return s
	}
	
	return s[:MaxNum]
}
