package tracker

import "image"

// Color represents a color tracker
type Color struct {
	color string
}

// Track return a image.Rectangle with tracked area
func (t *Color) Track() (image.Rectangle, error) {
	return image.Rect(0, 0, 0, 0), nil
}
