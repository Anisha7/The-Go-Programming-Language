// Mandelbrot emits a PNG image of the Mandelbrot fractal.

package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{
	color.RGBA{46, 49, 49, 2},
	color.RGBA{0, 230, 64, 1},
	color.RGBA{191, 85, 236, 1},
	color.RGBA{249, 105, 14, 1},
	color.RGBA{25, 181, 254, 1},
}

const (
	blackIndex  = 0 // first color in palette
	greenIndex  = 1 // next color in palette
	purpleIndex = 2
	orangeIndex = 3
	blueIndex   = 4
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img) // Note: ignoring errors
}

// checks if repeatedly squaring and adding the number that the point represents
// eventually "escapes" the circle of radius 2
// If so, the point is shaded by the number of iterations it took to escape
// else: it remains black
func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			// getting a random number
			s1 := rand.NewSource(time.Now().UnixNano())
			r1 := rand.New(s1)
			var colorIndex = r1.Intn(4)
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
