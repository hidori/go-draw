package draw

type Canvas struct {
	width  int
	height int
	pixels []byte
}

func NewCanvas(width int, height int) *Canvas {
	return &Canvas{
		width:  width,
		height: height,
		pixels: make([]byte, width*height),
	}
}

func (c *Canvas) SetPixel(x int, y int, color byte) error {
	c.pixels[c.width*y+x] = color
}

func (c *Canvas) GetPixel(x int, y int) (byte, error) {
	return c.pixels[c.width*y+x], nil
}

func (c *Canvas) calcIndex(x int, y int) (int, error) {
	
}
