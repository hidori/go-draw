package draw

import "errors"

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

func (c *Canvas) Width() int {
	return c.width
}

func (c *Canvas) Height() int {
	return c.height
}

func (c *Canvas) calcIndex(x int, y int) (int, error) {
	if x < 0 || c.width <= x || y < 0 || c.height <= y {
		return 0, errors.New("range over")
	}

	return c.width*y + x, nil
}

func (c *Canvas) SetPixel(x int, y int, color byte) error {
	index, err := c.calcIndex(x, y)
	if err != nil {
		return err
	}

	c.pixels[index] = color

	return nil
}

func (c *Canvas) GetPixel(x int, y int) (byte, error) {
	index, err := c.calcIndex(x, y)
	if err != nil {
		return 0, err
	}

	return c.pixels[index], nil
}
