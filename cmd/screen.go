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

// Uhm prob not gonna need this
