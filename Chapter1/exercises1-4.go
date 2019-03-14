// Lissajous generates GIF animations of random Lissajous figures
// 5. Change the Lissajous program's color palette to green on black, for added authenticity. 
// color.RGBA{0xRR, 0xGG, 0xBB, 0xff}
// green: rgba(0, 230, 64, 1), black: rgba(46, 49, 49, 1)
// 6. Modify the Lissajous program to produce images in multiple colors by adding more values 
// to palette and then displaying them by changing the third argument of SetColorIndex in 
// some interesting way.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
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
	blackIndex = 0 // first color in palette
	greenIndex = 1 // next color in palette
	purpleIndex = 2
	orangeIndex = 3
	blueIndex = 4
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles 	= 	5 		// number of complete x oscillator revolutions
		res 	= 	0.001 	// angular resolution
		size 	= 	100		// image canvas covers [-size...+size]
		nframes =	64		// number of animation frames
		delay	=	8		// delay between frames in 10ms unites
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{ LoopCount : nframes } // struct of type gif.GIF
	phase := 0.0 // phase difference

	var indices = []uint8 {greenIndex, purpleIndex, orangeIndex, blackIndex}

	for i := 0; i < nframes; i++ { // produces a single frame of animation
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		// getting a random number
		s1 := rand.NewSource(time.Now().UnixNano())
    	r1 := rand.New(s1)
		var colorIndex = r1.Intn(4)

		for t := 0.0; t < cycles*2*math.Pi; t += res { 
			// x oscillator cycles to set color to corresponding (x,y) to black
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), indices[colorIndex])
		}

		phase += 0.1
		// use dot notation to access struct fields
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim) // Note: ignoring encoding errors
}