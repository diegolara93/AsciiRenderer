package cmd

type Pixel struct {
	x     float64
	y     float64
	color string
}
type Screen struct {
	width  int
	height int
	pixels [][]Pixel
}

func NewScreen(width, height int) *Screen {
	return &Screen{width: width, height: height, pixels: make([][]Pixel, height*width)}
}

func (s *Screen) DrawStraightLine(x1, x2 float64) {
		
}
