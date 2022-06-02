package golang_united_school_homework

import (
	"errors"
	"fmt"
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if b.shapesCapacity <= len(b.shapes) {
		return errors.New("exceeded shapes capacity")
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i < 0 || i >= len(b.shapes) {
		return nil, errors.New("out of the range")
	}
	return b.shapes[i], nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if i < 0 || i >= len(b.shapes) {
		return nil, errors.New("out of the range")
	}
	outShape := b.shapes[i]
	b.shapes = append(b.shapes[0:i], b.shapes[i+1:]...)
	fmt.Println(b.shapes)
	return outShape, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if i < 0 || i >= len(b.shapes) {
		return nil, errors.New("out of the range")
	}
	outShape := b.shapes[i]
	b.shapes[i] = shape
	return outShape, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var sumPerimeter float64
	for _, shape := range b.shapes {
		sumPerimeter += shape.CalcPerimeter()
	}
	return sumPerimeter
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var sumArea float64
	for _, shape := range b.shapes {
		sumArea += shape.CalcArea()
	}
	return sumArea
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	circleCount := 0
	copyShapes := make([]Shape, len(b.shapes))
	copy(copyShapes, b.shapes)
	for i, shape := range copyShapes {
		switch shape.(type) {
		case *Circle:
			_, err := b.ExtractByIndex(i + circleCount)
			if err != nil {
				return fmt.Errorf("remove all circles erroe: %w", err)
			}
			circleCount--
		}
	}

	if circleCount == 0 {
		return errors.New("circles are not exist in the list")
	}
	return nil
}
