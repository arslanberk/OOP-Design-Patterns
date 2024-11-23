package main

import "fmt"

// Element
type Shape interface {
    getType() string
    accept(Visitor)
}

// Concrete element
type Square struct {
    side int
}

func (s *Square) accept(v Visitor) {
    v.visitForSquare(s)
}

func (s *Square) getType() string {
    return "Square"
}

// Concrete element
type Circle struct {
    radius int
}

func (c *Circle) accept(v Visitor) {
    v.visitForCircle(c)
}

func (c *Circle) getType() string {
    return "Circle"
}

// Concrete element
type Rectangle struct {
    l int
    b int
}

func (t *Rectangle) accept(v Visitor) {
    v.visitForrectangle(t)
}

func (t *Rectangle) getType() string {
    return "rectangle"
}

// Visitor
type Visitor interface {
    visitForSquare(*Square)
    visitForCircle(*Circle)
    visitForrectangle(*Rectangle)
}

// Concrete visitor
type AreaCalculator struct {
    area int
}

func (a *AreaCalculator) visitForSquare(s *Square) {
    // Calculate area for square.
    // Then assign in to the area instance variable.
    fmt.Println("Calculating area for square")
}

func (a *AreaCalculator) visitForCircle(s *Circle) {
    fmt.Println("Calculating area for circle")
}
func (a *AreaCalculator) visitForrectangle(s *Rectangle) {
    fmt.Println("Calculating area for rectangle")
}

// Concrete visitor
type MiddleCoordinates struct {
    x int
    y int
}

func (a *MiddleCoordinates) visitForSquare(s *Square) {
    // Calculate middle point coordinates for square.
    // Then assign in to the x and y instance variable.
    fmt.Println("Calculating middle point coordinates for square")
}

func (a *MiddleCoordinates) visitForCircle(c *Circle) {
    fmt.Println("Calculating middle point coordinates for circle")
}
func (a *MiddleCoordinates) visitForrectangle(t *Rectangle) {
    fmt.Println("Calculating middle point coordinates for rectangle")
}

// Client code
func main() {
    square := &Square{side: 2}
    circle := &Circle{radius: 3}
    rectangle := &Rectangle{l: 2, b: 3}

    areaCalculator := &AreaCalculator{}

    square.accept(areaCalculator)
    circle.accept(areaCalculator)
    rectangle.accept(areaCalculator)

    fmt.Println()
    middleCoordinates := &MiddleCoordinates{}
    square.accept(middleCoordinates)
    circle.accept(middleCoordinates)
    rectangle.accept(middleCoordinates)
}