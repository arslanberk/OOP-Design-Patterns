package main

import (
	"fmt"
	"strings"
)

// Product
type Product1 struct{
	parts []string
}
func (p Product1) listParts(){
	fmt.Println("Product parts: "+strings.Join(p.parts, ", "))
}

// Builder interface
type Builder interface{
	ProducePartA()
	ProducePartB()
	ProducePartC()
	GetProduct() Product1
}
func getBuilder() Builder{
	return newConcreteBuilder1()
}

// Concrete Builder
type ConcreteBuilder1 struct{ 
	product Product1 
}
func newConcreteBuilder1() *ConcreteBuilder1{
	return &ConcreteBuilder1{}
}
func (b *ConcreteBuilder1) ProducePartA(){
	b.product.parts = append(b.product.parts, "PartA1")
}
func (b *ConcreteBuilder1) ProducePartB(){
	b.product.parts = append(b.product.parts, "PartB1")
}
func (b *ConcreteBuilder1) ProducePartC(){
	b.product.parts = append(b.product.parts, "PartC1")
}
func (b *ConcreteBuilder1) GetProduct() Product1{
	result := Product1{
		parts: b.product.parts,
	}
	b.product = Product1{}
	return result
}

// Director
type Director struct{ 
	builder Builder 
}
func newDirector(b Builder) *Director{
	return &Director{
		builder: b,
	}
}
func (d *Director) SetBuilder(b Builder) {
	d.builder = b
}
func (d *Director) BuildMinimalViableProduct() {
	d.builder.ProducePartA()
}
func (d *Director) BuildFullFeaturedProduct() {
	d.builder.ProducePartA()
	d.builder.ProducePartB()
	d.builder.ProducePartC()
}

// Client Code
func ClientCode(director Director){
	builder := getBuilder()
	director.SetBuilder(builder)

	fmt.Println("Standard basic product:")
	director.BuildMinimalViableProduct()
	builder.GetProduct().listParts()
	fmt.Println("")

	fmt.Println("Standard full featured product:")
	director.BuildFullFeaturedProduct()
	builder.GetProduct().listParts()
	fmt.Println("")

	fmt.Println("Custom product:")
	builder.ProducePartA()
	builder.ProducePartC()
	builder.GetProduct().listParts()
}

// Main function
func main() {
    ClientCode(Director{})
}