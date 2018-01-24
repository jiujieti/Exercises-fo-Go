// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"fmt"
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
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
			fmt.Fprintln(os.Stderr, px, py)
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	x := (&big.Rat{}).SetFloat64(real(z))
	y := (&big.Rat{}).SetFloat64(imag(z))
	v, w := &big.Rat{}, &big.Rat{}
	for n := uint8(0); n < iterations; n++ {
		t1, t2, t3 := &big.Rat{}, &big.Rat{}, &big.Rat{}
		t1.Sub(t1.Mul(v, v), t2.Mul(w, w))
		t1.Add(t1, x)
		t3.Set(t1)
		t1.Mul(v, w)
		t1.Add(t1.Add(t1, t1), y)
		w.Set(t1)
		v.Set(t3)
		t1.Add(t1.Mul(v, v), t2.Mul(w, w))
		if t1.Cmp(big.NewRat(4, 1)) == 1 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

//!-
