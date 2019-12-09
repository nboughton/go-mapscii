package mapscii

import "strings"

var (
	hexTmpl = `##\__________/##
##/##########\##
#/############\#
/##############\
\##############/
#\############/#
##\__________/##`

	offset = 5
)

type hexCell struct {
	tmpl     [][]string
	widthTop int
	widthMid int
	height   int
	crdsRow  int // Set the row/col to align the coords to
	crdsCol  int
}

func newHexCell(row, col int) *hexCell {
	c := &hexCell{
		tmpl:    ParseCellTmpl(hexTmpl),
		crdsRow: 1,
		crdsCol: 3,
	}

	c.height = len(c.tmpl)
	c.widthTop = (len(c.tmpl[0]) / 3) * 2
	c.widthMid = len(c.tmpl[0])
	c.setCrds(row, col, false, false)

	return c
}

func (c *hexCell) setCrds(row, col int, rowAlpha, colAlpha bool) {
	for i, sub := range CoordText(row, col, rowAlpha, colAlpha) {
		c.tmpl[c.crdsRow][c.crdsCol+i] = string(sub)
	}
}

// HexMap
type HexMap [][]string

// NewHexMap generates a hex template so that tmpl can be superimposed on it
func NewHexMap(height, width int) HexMap {
	var (
		cl    = newHexCell(0, 0) // Use a cell as a reference
		wDiff = (cl.widthMid - cl.widthTop) / 2
		w     = (width * (cl.widthMid - wDiff)) + wDiff
		h     = (height * (cl.height - 1)) + cl.height/2 + 1 // Shared borders reduce total height
		m     = make(HexMap, h)
	)

	// Create blank template space
	for r := 0; r < h; r++ {
		m[r] = make([]string, w)
		for c := range m[r] {
			m[r][c] = " "
		}
	}

	for r := 0; r < height; r++ {
		row := r * (cl.height - 1)

		for c := 0; c < width; c++ {
			col := c * (cl.widthMid - 3)
			if c%2 != 0 {
				row = r*(cl.height-1) + cl.height/2
			}

			m.emptyCell(row, col, r, c)
			if c%2 != 0 {
				row = row - cl.height/2
			}
		}
	}

	return m
}

// emptyCell writes a blank cell to the Map matrix
func (m HexMap) emptyCell(row, col, rLabel, cLabel int) HexMap {
	r, c, cl := row, col, newHexCell(rLabel, cLabel)

	for cellRow := 0; cellRow < cl.height; cellRow++ {
		for _, char := range cl.tmpl[cellRow] {
			if r >= len(m) || c >= len(m[r]) { // Bounds check matrices references
				return m
			}

			m[r][c] = char
			c++
		}
		r++
		c = col
	}

	return m
}

// SetTxt sets the text of a given hex
func (m HexMap) SetTxt(row, col int, lines [4]string, color colourFunc) {
	// Define row and column based on the same calculations used to place blank
	// cells
	var (
		cl    = newHexCell(0, 0)
		wDiff = (cl.widthMid - cl.widthTop) / 2
		r     = row*(cl.height-1) + cl.crdsRow
		c     = col*(cl.widthMid-wDiff) + cl.crdsCol
	)

	if col%2 != 0 {
		r += cl.height / 2
	}

	for i, line := range lines {
		m.print(r+i+1, c+offset-(len(line)/2), line, color)
	}
}

func (m HexMap) print(startRow, startCol int, text string, colour colourFunc) {
	for row, col, i := startRow, startCol, 0; i < len(text); col, i = col+1, i+1 {
		if col < 0 {
			col = 0
		}

		if col < len(m[row]) {
			m[row][col] = colour(string(text[i]))
		} else {
			m[row] = append(m[row], colour(string(text[i])))
		}
	}
}

func (m HexMap) String() string {
	s := ""

	for _, line := range m {
		s += strings.Join(line, "") + "\n"
	}

	return s
}
