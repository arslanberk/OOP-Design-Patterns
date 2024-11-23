package main

import (
	"fmt"
)

// Product interface
type Product interface{
	Operation() string
}
// Concrete products
type ConcreteProduct1 struct{}
func (p *ConcreteProduct1) Operation() string{
	return "{Result of the ConcreteProduct1}"
}
type ConcreteProduct2 struct{}
func (p *ConcreteProduct2) Operation() string{
	return "{Result of the ConcreteProduct2}"
}
// Creator
type Creator interface{
	FactoryMethod() Product
}
func SomeOperation(c Creator) string{
	product := c.FactoryMethod()
	return "Creator: The same creator's code has just worked with " + product.Operation()
}
// Concrete Creator
type ConcreteCreator1 struct{}
func (c *ConcreteCreator1) FactoryMethod() Product{
	return &ConcreteProduct1{}
}
func newConcreteCreator1() *ConcreteCreator1{
	return &ConcreteCreator1{}
}
type ConcreteCreator2 struct{}
func (c *ConcreteCreator2) FactoryMethod() Product{
	return &ConcreteProduct2{}
}
// Client Code
func ClientCode(creator Creator){
	fmt.Println("Client: I'm not aware of the creator's class, but it still works.")
	fmt.Println(SomeOperation(creator))
}
// Main function
func main() {
	fmt.Println("App: Launched with the ConcreteCreator1.")
    ClientCode(newConcreteCreator1())
	fmt.Println("")
	fmt.Println("App: Launched with the ConcreteCreator2.")
    ClientCode(&ConcreteCreator2{})
}