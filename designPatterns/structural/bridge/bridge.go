package main

import "fmt"

// IDrawShape Interface
type IDrawShape interface {
	drawShape(x [5]float32, y [5]float32)
}

// DrawShape struct
type DrawShape struct{}

// DrawShape struct has the drawShape method with float x and y coordinates
func (draShape DrawShape) drawShape(x [5]float32, y [5]float32) {
	fmt.Println("Drawing Shape")
}

// IContour interface
type IContour interface {
	drawContour(x[5] float32, y[5] float32)
	resizeByFactor(factor int)
}

// DrawContour struct
type DrawContour struct {
	x[5] float32
	y[5] float32
	shape DrawShape
	factor int
}

// drawCountour method of DrawCounter struct
func (contour DrawContour) drawContour(x[5] float32, y[5] float32) {
	fmt.Println("Drawing Contour..")
	contour.shape.drawShape(contour.x, contour.y)
}

// resizeByFactor method for the DrawContour struct to implement the IContour interface
func(contour DrawContour) resizeByFactor(factor int) {
	contour.factor = factor
}

// let it all begin!
func main() {
	var x = [5]float32{4, 6, 2, 1, 9}
	var y = [5]float32{0, 2, 5, 6, 8}

	// interface for contract, assign the datatype to the interface
	var contour IContour = DrawContour{
		x,
		y,
		DrawShape{},
		2,
	}

	contour.drawContour(x, y)
	contour.resizeByFactor(2)
}
