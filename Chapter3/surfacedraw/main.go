// Lissajous generates GIF animations of random Lissajous figures

package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"net/http"
	"os"
)

var palette = []color.Color{color.White, color.Black}

const (
	red        = "#ff0000"
	blue       = "#0000ff"
	blackIndex = 1
	whiteIndex = 0
	// new
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axis (= 30 degrees)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // angle = 30 degrees
var max = math.MaxFloat64

func main() {
	surface(os.Stdout)
	// server work, connecting to web server
	handler := func(w http.ResponseWriter, r *http.Request) {
		surface(w)
	}
	http.HandleFunc("/", handler)
}

func surface(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 600   // image canvas covers [-size...+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms unites

	)

	// freq := rand.Float64() * 3.0        // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes} // struct of type gif.GIF
	phase := 0.0                        // phase difference

	for i := 0; i < nframes; i++ { // produces a single frame of animation
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for i := 0; i < cells; i++ {
			for j := 0; j < cells; j++ {
				// evaluate and check for valid polygon
				ax, ay, err := corner(i+1, j)
				if err != nil {
					continue
				}
				img.SetColorIndex(int(ax), int(ay), blackIndex)
				bx, by, err := corner(i, j)
				if err != nil {
					continue
				}
				img.SetColorIndex(int(bx), int(by), blackIndex)
				cx, cy, err := corner(i, j+1)
				if err != nil {
					continue
				}
				img.SetColorIndex(int(cx), int(cy), blackIndex)
				dx, dy, err := corner(i+1, j+1)
				if err != nil {
					continue
				}
				img.SetColorIndex(int(dx), int(dy), blackIndex)
			}
		}

		phase += 0.1
		// use dot notation to access struct fields
		// anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim) // Note: ignoring encoding errors
}

func corner(i, j int) (float64, float64, error) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z
	z := f(x, y)
	if z > max {
		return 0, 0, fmt.Errorf("Invalid Id")
	}

	// Project (x,y,z) isometrically onto 2-D SCG canvas (sx, sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, nil
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
