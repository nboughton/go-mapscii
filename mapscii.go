// Package mapscii is a simple libary that overlays cell text over a generated ascii hex map
package mapscii

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

type colourFunc func(string, ...interface{}) string

var (
	offset = 5
)

// Colour funcs
var (
	White   = color.WhiteString
	Red     = color.RedString
	Yellow  = color.YellowString
	Magenta = color.MagentaString
	Green   = color.GreenString
	Blue    = color.BlueString
	Cyan    = color.CyanString

	sqTmpl = `______________
|              |
|              |
|              |
|              |
|              |
|______________|`
)

type Cell interface {
	Width() int
	Height() int
	Text() [][]string
}

// ParseCellTmpl reads a cell template and returns it as a character matrix. Cell templates should use
// '#' instead of white space as gofmt automatically strips whitespace off the ends of lines.
/*  For example, the default hex template is declared as follows:
var hexTmpl = `##\__________/##
##/##########\##
#/############\#
/##############\
\##############/
#\############/#
##\__________/##`
*/
func ParseCellTmpl(s string) [][]string {
	var out [][]string

	for _, row := range strings.Split(s, "\n") {
		r := []string{}
		for _, col := range row {
			if string(col) == "#" {
				col = rune(' ')
			}
			r = append(r, string(col))
		}
		out = append(out, r)
	}

	return out
}

func genCrdText(row, col int) string {
	rStr, cStr := strconv.Itoa(row), strconv.Itoa(col)
	if row < 10 {
		rStr = "0" + rStr
	}
	if col < 10 {
		cStr = "0" + cStr
	}

	return fmt.Sprintf("%s,%s", rStr, cStr)
}

// Colour toggles the colour output on/off (true/false)
func Colour(b bool) {
	color.NoColor = !b
}
