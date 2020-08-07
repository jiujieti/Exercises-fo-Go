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
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -10, -10, +10, +10
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, newton(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func newton(z complex128) color.Color {

	const max = 30
	const contrast = 8
	const threshold = 1e-6

	// var v complex128
	iterations := 0
	for distance(z, 1) >= threshold && distance(z, -1) >= threshold && distance(z, -1i) >= threshold && distance(z, 1i) >= threshold && iterations < max {
		z = (3*z + 1/(z*z*z)) / 4
		iterations++
	}
	// fmt.Fprintf(os.Stderr, "z=%v iterations=%d\n", z, iterations)
	if iterations == max {
		return color.Black
	}
	a := uint8(iterations * contrast)
	if distance(z, 1) < threshold {
		return color.RGBA{0xff, 0x00, 0x00, a}
	} else if distance(z, -1) < threshold {
		return color.RGBA{0xff, 0xff, 0x00, a}
	} else if distance(z, 1i) < threshold {
		return color.RGBA{0x00, 0xff, 0x00, a}
	} else {
		return color.RGBA{0x00, 0x00, 0xff, a}
	}
}

func distance(z complex128, r complex128) float64 {
	return cmplx.Abs(z - r)
}
