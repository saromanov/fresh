package pkg

import (
	"github.com/fatih/color"
)

// Info prints text in blue
func Info(s string) {
	color.Blue(s)
}

// Warning rprint text in red
func Warning(s string) {
	color.Red(s)
}
