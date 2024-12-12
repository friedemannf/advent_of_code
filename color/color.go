package color

import (
	"github.com/fatih/color"
)

var Red = color.New(color.FgRed).SprintFunc()
var Green = color.New(color.FgGreen).SprintFunc()
var Blue = color.New(color.FgBlue).SprintFunc()
var Yellow = color.New(color.FgYellow).SprintFunc()
