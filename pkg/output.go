package pkg

import (
	"github.com/fatih/color"
)

// Info prints text in blue
func Info(s string) {
	color.Blue(s)
}

// Infof prints text in blue with formatting
func Infof(f string, s ...interface{}) {
	color.Blue(f, s...)
}

// Text returns formatted string
func Text(s string) string {
	return color.GreenString(s)
}

// Warningf rprint text in red
func Warningf(f string, s ...interface{}) {
	color.Red(f, s...)
}
