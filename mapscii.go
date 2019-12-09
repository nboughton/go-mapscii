// Package mapscii is a simple libary that overlays cell text over a generated ascii hex map
package mapscii

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

type colourFunc func(string, ...interface{}) string

// Colour funcs
var (
	White   = color.WhiteString
	Red     = color.RedString
	Yellow  = color.YellowString
	Magenta = color.MagentaString
	Green   = color.GreenString
	Blue    = color.BlueString
	Cyan    = color.CyanString

	abc = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
)

type Cell interface {
	Row() int
	Col() int
	Width() int
	Height() int
	Tmpl() [][]string
	SetContent([]string)
}

type Map [][]Cell

/* ParseCellTmpl reads a cell template and returns it as a character matrix. Cell templates should use
 '#' instead of white space as gofmt automatically strips whitespace off the ends of lines.
  For example, the default hex template is declared as follows:
var hexTmpl = `##\__________/##
##/##########\##
#/############\#
/##############\
\##############/
#\############/#
##\__________/##` */
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

// Generate coordinate identifer text. If rowAlpha or colAlpha is true
// the row/col identifer returned will be in the form of a letter or letters
// where every number over 26 increments to another cycle of the alphabet.
// I.e 1 = a, 3 = c, 27 = aa
func CoordText(row, col int, rowAlpha, colAlpha bool) string {
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
