// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/big"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := big.NewFloat(float64(py)/height*(ymax-ymin) + ymin)
		for px := 0; px < width; px++ {
			x := big.NewFloat(float64(px)/width*(xmax-xmin) + xmin)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(x, y))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(x, y *big.Float) color.Color {
	const iterations = 200
	const contrast = 15

	var v, w *big.Float
	for n := uint8(0); n < iterations; n++ {
		var t1, t2, t3, t4 *big.Float
		t3.Mul(v, v)
		t4.Mul(w, w)
		t4.Sub(t3, t4)
		t1.Add(x, t4)
		t3.Mul(v, w)
		t2.Add(t3.Add(t3, t3), y)
		v = t1
		w = t2
		t4.Abs(t3.Add(t1.Mul(t1, t1), t2.Mul(t2, t2)))
		if t4.Cmp(big.NewFloat(float64(4))) == 1 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

//!-
