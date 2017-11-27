package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
	"strconv"
)

// inMandelSet returns the number of iteration it takes to test for divergence // at any given point (x,y) on the complex plane.
// If we reach max_iterations without proving divergence,
// then we assume the point converges.

func inMandelSet(x, y float64, max_iters int) int {
	var (
		k int        = 0
		z complex128 = 0
		c complex128 = complex(x, y)
	)
	for k = 0; k < max_iters; k++ {
		z = cmplx.Pow(z, 4) + c
		if cmplx.Abs(z) > 2 {
			return k
		}
	}
	return k
}

// drawDot adds a colored pixel to the input image.
// Does stuff to make nice colors for R,G,B

func drawDot(img *image.RGBA, q, r int, valueX, valueY float64, m int) {

	var myRed, myGreen, myBlue uint8
	i := inMandelSet(valueX, valueY, m)

	myGreen = uint8((-i)*255/m - 255)
	myRed = uint8((-i)*255/m - 255)

	// myRed = uint8( i * 255 / m )
	myBlue = uint8(i * 255 / m)

	img.Set(q, r, color.RGBA{myRed, myGreen, myBlue, 255})
	return
}

// Main Thing

func buildFractal(width, height, max_iters int, xo, yo, Re float64, filename string) {

	w := float64(width)
	q := float64(height)
	deltaX := math.Abs(2*Re) / w
	deltaY := math.Abs(2*Re) / q

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Draw a red dot at each point inside the set.
	for q := 0; q < width; q++ {
		for r := 0; r < height; r++ {

			valueX := float64(q)*deltaX + xo - Re
			valueY := float64(r)*deltaY + yo - Re

			drawDot(img, q, r, valueX, valueY, max_iters)
		}
		// fmt.Println(width-q-1)
	}

	// Save to a special filename
	// The filename is going to have the input values built into it.
	// This should help separate the files after making them.

	// xoStr := strconv.FormatFloat(xo, 'E', -1, 64)
	// yoStr := strconv.FormatFloat(yo, 'E', -1, 64)
	// ReStr := strconv.FormatFloat(Re, 'E', -1, 64)
	// itStr := strconv.Itoa(max_iters)

	// filename := "Mandelbrot Set "+"x=("+xoStr+") y=("+yoStr+") r=("+ReStr+") iterations=("+itStr+")"
	// fmt.Println("Image has been created! \n Filename: "+filename)

	filepath := "coolPictures/" + filename + ".png"

	// This part actually creates the file itself.

	f, _ := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)
}

// Adding Parralellism for multicore processors

// Main Function that starts automatically.

func main() {
	// Declare Initial Variables.
	// These are the main things that affect the fractals.
	const (
		width     int = 200
		height    int = 200
		max_iters int = 150

		xo float64 = -1.11
		yo float64 = 0
		Re float64 = 1
	)

	imin := 1
	imax := 100
	woah := Re
	varyIters := max_iters

	for i := imin; i < imax; i++ {
		buildFractal(width, height, varyIters, xo, yo, woah, strconv.Itoa(i))

		woah = woah / 1.05
		varyIters = varyIters + 1
		fmt.Println(i, varyIters, woah)
	}
	fmt.Println("COMPLETED !")

}
