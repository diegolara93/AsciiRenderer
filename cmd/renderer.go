package cmd

import (
	"github.com/gdamore/tcell/v2"
	"math"
)

// Bresenhams line algorithm implemented from: https://zingl.github.io/bresenham.html
func DrawLine(s tcell.Screen, x0, y0, x1, y1 int, style tcell.Style) {
	dx := int(math.Abs(float64(x1 - x0)))
	dy := -int(math.Abs(float64(y1 - y0)))
	sx := 0
	sy := 0
	err := dx + dy
	if x0 < x1 {
		sx = 1
	} else {
		sx = -1
	}
	if y0 < y1 {
		sy = 1
	} else {
		sy = -1
	}
	for {
		s.SetContent(x0, y0, '@', nil, style)
		e2 := 2 * err
		if e2 >= dy {
			if x0 == x1 {
				break
			} else {
				err += dy
				x0 += sx
			}
		}
		if e2 <= dx {
			if y0 == y1 {
				break
			} else {
				err += dx
				y0 += sy
			}
		}
	}
}

/*
					x0, y0				x0 + width, y0
	 				-----------------
					|				|
					|				|
					|				|
					|				|
					|				|
					-----------------
				x0, y0+height		x0+width, y0+height
*/
func DrawSquare(s tcell.Screen, x0, y0, width, height int) {
	DrawLine(s, x0, y0, x0+width, y0, tcell.StyleDefault)               // Draws top line
	DrawLine(s, x0+width, y0, x0+width, y0+height, tcell.StyleDefault)  // Draws right line
	DrawLine(s, x0+width, y0+height, x0, y0+height, tcell.StyleDefault) // Draws bottom line
	DrawLine(s, x0, y0+height, x0, y0, tcell.StyleDefault)              // Draws left line
}

/*
x0 y0 indicate the tip of the triangle, for example
*/
func DrawTriangle(s tcell.Screen, x0, y0, width, height int) {

}

func DrawText(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string) {
	row := y1
	col := x1
	for _, r := range []rune(text) {
		s.SetContent(col, row, r, nil, style)
		col++
		if col >= x2 {
			row++
			col = x1
		}
		if row > y2 {
			break
		}
	}
}

func Render(s tcell.Screen) {

}
