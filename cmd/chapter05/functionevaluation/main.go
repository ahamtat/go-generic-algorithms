package main

import (
	"fmt"
	"github.com/ahamtat/go-generic-algorithms/pkg/stack"
)

var values = map[string]float64{
	"a": 10,
	"b": 5,
	"c": 2,
	"d": 4,
	"e": 3,
}

func precedence(symbol1, symbol2 string) bool {
	// Returns true if symbol1 has a higher precedence
	// than symbol2
	if (symbol1 == "+" || symbol1 == "-") && (symbol2 == "(" || symbol2 == "/") {
		return false
	} else if (symbol1 == "(" && symbol2 != ")") || symbol2 == "(" {
		return false
	} else {
		return true
	}
}

func isPresent(symbol string, operators []string) bool {
	for i := 0; i < len(operators); i++ {
		if symbol == string(operators[i]) {
			return true
		}
	}
	return false
}

func infixpostfix(infix string) (postfix string) {
	operators := []string{"+", "-", "*", "/", "(", ")"}
	postfix = ""
	nodeStack := stack.Nodes[string]{}
	for index := 0; index < len(infix); index++ {
		newSymbol := string(infix[index])
		if newSymbol == " " || newSymbol == "\n" {
			continue
		}

		if newSymbol >= "a" && newSymbol <= "z" {
			postfix += newSymbol
		}

		if isPresent(newSymbol, operators) {
			if !nodeStack.IsEmpty() {
				topSymbol := nodeStack.Top()
				if precedence(topSymbol, newSymbol) {
					if topSymbol != "(" {
						postfix += topSymbol
					}
					_, _ = nodeStack.Pop()
				}
			}

			if newSymbol != ")" {
				nodeStack.Push(newSymbol)
			} else {
				// Pop nodeStack down to first left parenthesis
				for {
					if nodeStack.IsEmpty() {
						break
					}
					char := nodeStack.Top()
					if char != "(" {
						postfix += char
						_, _ = nodeStack.Pop()
					} else {
						_, _ = nodeStack.Pop()
						break
					}
				}
			}
		}
	}

	for {
		if nodeStack.IsEmpty() {
			break
		}

		if nodeStack.Top() != "(" {
			postfix += nodeStack.Top()
			_, _ = nodeStack.Pop()
		}
	}

	return postfix
}

func evaluate(postfix string) float64 {
	operandStack := stack.Nodes[float64]{}

	for index := 0; index < len(postfix); index++ {
		ch := string(postfix[index])
		if ch >= "a" && ch <= "z" {
			operandStack.Push(values[ch])
		} else {
			// ch is an operator
			operand1, _ := operandStack.Pop()
			operand2, _ := operandStack.Pop()
			switch ch {
			case "+":
				operandStack.Push(operand1 + operand2)
			case "-":
				operandStack.Push(operand1 - operand2)
			case "*":
				operandStack.Push(operand1 * operand2)
			case "/":
				operandStack.Push(operand1 / operand2)
			}
		}
	}

	return operandStack.Top()
}

func main() {
	postfix := infixpostfix("a + (b - c) / (d * e)")
	fmt.Println(postfix)

	result := evaluate(postfix)
	fmt.Println("function evaluates to: ", result)
}
