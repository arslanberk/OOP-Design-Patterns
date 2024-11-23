package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Expression interface with Interpret method
type Expression interface {
	Interpret() int
}

// Number struct for terminal expressions
type Number struct {
	value int
}

func (n *Number) Interpret() int {
	return n.value
}

// Plus struct for non-terminal expressions
type Plus struct {
	left, right Expression
}

func (p *Plus) Interpret() int {
	return p.left.Interpret() + p.right.Interpret()
}

// Minus struct for non-terminal expressions
type Minus struct {
	left, right Expression
}

func (m *Minus) Interpret() int {
	return m.left.Interpret() - m.right.Interpret()
}

// Parser to build the AST
func parse(expression string) Expression {
	tokens := strings.Fields(expression)
	stack := []Expression{}

	for _, token := range tokens {
		switch token {
		case "+":
			right := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			left := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, &Plus{left: left, right: right})
		case "-":
			right := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			left := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, &Minus{left: left, right: right})
		default:
			value, _ := strconv.Atoi(token)
			stack = append(stack, &Number{value: value})
		}
	}

	return stack[0]
}

func main() {
	expression := "5 3 + 2 -"
	ast := parse(expression)
	result := ast.Interpret()
	fmt.Printf("Result of '%s' is %d\n", expression, result)
}