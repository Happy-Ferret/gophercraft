package glogger

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

const (
	DEBUG = "*"
	WARN  = "!"
	FATAL = ">"
	OK    = "✓"
)

var Colors = map[string]color.Attribute{
	DEBUG: color.FgGreen,
	WARN:  color.FgYellow,
	FATAL: color.FgRed,
	OK:    color.FgWhite,
}

func prints(ty string, data string) {
	colr := Colors[ty]
	color.Set(colr)
	str := fmt.Sprintf("[ %s ] ", ty) + data
	fmt.Print(str)
	color.Unset()
}

func Fatal(args ...interface{}) {
	prints(FATAL, fmt.Sprintln(args...))
	os.Exit(0)
}

func Println(args ...interface{}) {
	prints(OK, fmt.Sprintln(args...))
}

func Warnln(args ...interface{}) {
	prints(WARN, fmt.Sprintln(args...))
}

func Debugln(args ...interface{}) {
	prints(DEBUG, fmt.Sprintln(args...))
}

func Debugf(fomt string, args ...interface{}) {
	prints(DEBUG, fmt.Sprintf(fomt, args...))
}
