package main

import "fmt"

// AbstractFactory interface
type AbstractFactory interface {
	CreateProductA() AbstractProductA
	CreateProductB() AbstractProductB
}

// Concrete Factories
// ConcreteFactory1 implements AbstractFactory
type ConcreteFactory1 struct{}
func(factory ConcreteFactory1) CreateProductA() AbstractProductA {
	return ConcreteProductA1{}
}
func(factory ConcreteFactory1) CreateProductB() AbstractProductB {
	return ConcreteProductB1{}
}

// ConcreteFactory2 implements AbstractFactory
type ConcreteFactory2 struct{}
func(factory ConcreteFactory2) CreateProductA() AbstractProductA {
	return ConcreteProductA2{}
}
func(factory ConcreteFactory2) CreateProductB() AbstractProductB {
	return ConcreteProductB2{}
}

// AbstractProductA interface
type AbstractProductA interface {
    UsefulFunctionA() string
}

// ConcreteProductA1 implements AbstractProductA 
type ConcreteProductA1 struct{}
func (product ConcreteProductA1) UsefulFunctionA() string {
    return "The result of the product A1."
}

// ConcreteProductA2 implements AbstractProductA
type ConcreteProductA2 struct{}
func (product ConcreteProductA2) UsefulFunctionA() string {
    return "The result of the product A2."
}

// AbstractProductB interface
type AbstractProductB interface {
    UsefulFunctionB() string
    AnotherUsefulFunctionB(collaborator AbstractProductA) string
}

// ConcreteProductB1 implements AbstractProductB
type ConcreteProductB1 struct{}
func (product ConcreteProductB1) UsefulFunctionB() string {
    return "The result of the product B1."
}
func (product ConcreteProductB1) AnotherUsefulFunctionB(collaborator AbstractProductA) string {
    return "The result of the B1 collaborating with the (" + collaborator.UsefulFunctionA() +")"
}

// ConcreteProductB2 implements AbstractProductB
type ConcreteProductB2 struct{}
func (product ConcreteProductB2) UsefulFunctionB() string {
    return "The result of the product B2."
}
func (product ConcreteProductB2) AnotherUsefulFunctionB(collaborator AbstractProductA) string {
    return "The result of the B2 collaborating with the (" + collaborator.UsefulFunctionA() +")"
}

// Client code
func ClientCode(factory AbstractFactory){
	var productA = factory.CreateProductA()
	var productB = factory.CreateProductB()

	fmt.Println(productB.UsefulFunctionB())
	fmt.Println(productB.AnotherUsefulFunctionB(productA))
}

// Main function
func main() {
	fmt.Println("Client: Testing client code with the first factory type:")
    ClientCode(ConcreteFactory1{})
	fmt.Println("")
	fmt.Println("Client: Testing the same client code with the second factory type:")
    ClientCode(ConcreteFactory2{})
}