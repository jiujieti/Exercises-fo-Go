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
		xmin, ymin, xmax, ymax = -0.55, -0.55, -0.54, -0.54
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	xp := big.NewFloat(real(z))
	yp := big.NewFloat(imag(z))
	r, i := &big.Float{}, &big.Float{}
	for n := uint8(0); n < iterations; n++ {
		t1, t2, t3 := &big.Float{}, &big.Float{}, &big.Float{}
		t1.Sub(t1.Mul(r, r), t2.Mul(i, i))
		t1.Add(t1, xp)
		t2.Mul(r, i)
		t3.Add(t2, t2)
		t3.Add(t3, yp)
		r.Copy(t1)
		i.Copy(t3)
		t1.Add(t1.Mul(r, r), t2.Mul(i, i))
		if t1.Cmp(big.NewFloat(float64(4))) == 1 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
