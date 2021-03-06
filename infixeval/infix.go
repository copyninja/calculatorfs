package infix

import (
	"errors"
	// "fmt"
	"strconv"
	"strings"
	"text/scanner"
)

// Operators represented as rune's
const (
	LPAREN = iota
	RPAREN
	ADD
	SUB
	MUL
	DIV
	MOD
)

type infixEval struct {
	operator Stack
	operand  Stack
}

func translateOperator(text string) rune {
	if text == "(" {
		return LPAREN
	}

	if text == ")" {
		return RPAREN
	}

	if text == "+" {
		return ADD
	}

	if text == "-" {
		return SUB
	}

	if text == "*" {
		return MUL
	}

	if text == "/" {
		return DIV
	}

	if text == "%" {
		return MOD
	}

	return -1
}

func precedence(op rune) int {
	if op == MUL || op == DIV || op == MOD {
		return 1
	}

	return 0
}

func evalInt(val1, val2 int64, op rune) (int64, error) {
	switch op {
	case ADD:
		return val1 + val2, nil
	case SUB:
		return val1 - val2, nil
	case MUL:
		return val1 * val2, nil
	case DIV:
		if val2 == 0 {
			return 0, errors.New("Attempt to divide by 0")
		}

		return val1 / val2, nil
	case MOD:
		if val2 == 0 {
			return 0, errors.New("Attempt to divide by 0")
		}
		return val1 % val2, nil
	}

	return 0, nil
}

// Eval executes operation on the given operands and returns result.
func eval(val1, val2 interface{}, op rune) (interface{}, error) {
	a, ok1 := val1.(float64)
	b, ok2 := val2.(float64)

	if !ok1 && !ok2 {
		return evalInt(val1.(int64), val2.(int64), op)
	} else if ok1 && ok2 {
		if op == MOD {
			return 0, errors.New("Modulus operator is not supported on float")
		}
	} else {

		if ok1 && !ok2 {
			b = float64(val2.(int64))
		} else {
			a = float64(val1.(int64))
		}
	}

	switch op {
	case ADD:
		return a + b, nil
	case SUB:
		return a - b, nil
	case MUL:
		return a * b, nil
	case DIV:
		return a / b, nil
	}

	return 0, nil
}

func (infix *infixEval) evaluateTop() (interface{}, error) {
	val2 := infix.operand.Pop()
	val1 := infix.operand.Pop()
	op := infix.operator.Pop()

	return eval(val1, val2, op.(rune))
}

func (infix *infixEval) handleInput(tok, symbol rune, text string) error {
	switch symbol {
	case LPAREN:
		infix.operator.Push(symbol)
	case ADD, SUB, MUL, DIV, MOD:
		for !infix.operator.IsEmpty() && (precedence(symbol) < precedence(infix.operator.Peek().(rune))) {
			result, err := infix.evaluateTop()
			if err != nil {
				return err
			}
			infix.operand.Push(result)
		}
		infix.operator.Push(symbol)
	case RPAREN:
		for infix.operator.Peek().(rune) != LPAREN {
			result, err := infix.evaluateTop()
			if err != nil {
				return err
			}
			infix.operand.Push(result)
			if infix.operator.IsEmpty() {
				return errors.New("Invalid/unbalanced expression")
			}
		}
		// pop the (
		infix.operator.Pop()
	default:
		if tok == scanner.Int {
			value, _ := strconv.ParseInt(text, 10, 64)
			infix.operand.Push(value)
		} else if tok == scanner.Float {
			value, _ := strconv.ParseFloat(text, 64)
			infix.operand.Push(value)
		} else {
			return errors.New("Invalid tokens in the expression")
		}
	}

	return nil
}

// Process parses the infix expression and returns the result
func (infix *infixEval) process(expression string) error {
	input := strings.NewReader(expression)
	var s scanner.Scanner
	s.Filename = "<stdin>"
	s.Init(input)
	var tok rune

	for {
		if tok = s.Scan(); tok == scanner.EOF {
			break
		}

		text := s.TokenText()
		symbol := translateOperator(text)

		if err := infix.handleInput(tok, symbol, text); err != nil {
			return err
		}
	}

	for !infix.operator.IsEmpty() {
		if infix.operand.Count() < 2 {
			return errors.New("Invalid/unbalanced expression")
		}

		result, err := infix.evaluateTop()
		if err != nil {
			return err
		}
		infix.operand.Push(result)
	}

	return nil
}

// Evaluate the given infix expression, returns result on success and error on
// parsing failure
func Evaluate(expression string) (interface{}, error) {
	i := infixEval{}
	if err := i.process(expression); err != nil {
		return nil, err
	}

	if i.operand.IsEmpty() || i.operand.Count() > 1 {
		return nil, errors.New("Invalid/unbalanced expression")
	}

	return i.operand.Pop(), nil
}
