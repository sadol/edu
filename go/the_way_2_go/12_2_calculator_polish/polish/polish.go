// using polish notation in order to make simple calculators
// ASSUMPTION: ALL OPERATORS ARE BINARY!!!
package polish

import (
	"../genstack"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Operand genstack.Element
type Operator genstack.Element
type Token genstack.Element

const MAXVALUE = 999999 // max value of the operand

type tokenType int

const ( // token types
	OPERATOR tokenType = iota
	OPERAND
	ERR
)

// mapping operator string to operation;
// usage : Operation[Operator](operand1.(type), operand2.(type));
// it's not part of the NPN structure because of possibility of creating
// another structure of the polish notation -> Reverse Polish Notation

// TODO: replace this repetitions with closure of some kind
// summator:
func summator(oprnd1, oprnd2 Operand) (result Operand, erro error) {
	var errMsg = "Unknown type or uncomparable types of operands."
	switch oprnd1.(type) { // DOUBLE!!! type switch: OMG
	case int64:
		switch oprnd2.(type) {
		case int64:
			result = oprnd1.(int64) + oprnd2.(int64)
		default:
			erro = errors.New(errMsg)
		}
	case float64:
		switch oprnd2.(type) {
		case float64:
			result = oprnd1.(float64) + oprnd2.(float64)
		default:
			erro = errors.New(errMsg)
		}
	default:
		erro = errors.New(errMsg)
	}
	return
}

// differentiator
func differentiator(oprnd1, oprnd2 Operand) (result Operand, erro error) {
	var errMsg = "Unknown type or uncomparable types of operands."
	switch oprnd1.(type) { // DOUBLE!!! type switch: OMG
	case int64:
		switch oprnd2.(type) {
		case int64:
			result = oprnd1.(int64) - oprnd2.(int64)
		default:
			erro = errors.New(errMsg)
		}
	case float64:
		switch oprnd2.(type) {
		case float64:
			result = oprnd1.(float64) - oprnd2.(float64)
		default:
			erro = errors.New(errMsg)
		}
	default:
		erro = errors.New(errMsg)
	}
	return
}

// multiplicator
func multiplicator(oprnd1, oprnd2 Operand) (result Operand, erro error) {
	var errMsg = "Unknown type or uncomparable types of operands."
	switch oprnd1.(type) { // DOUBLE!!! type switch: OMG
	case int64:
		switch oprnd2.(type) {
		case int64:
			result = oprnd1.(int64) * oprnd2.(int64)
		default:
			erro = errors.New(errMsg)
		}
	case float64:
		switch oprnd2.(type) {
		case float64:
			result = oprnd1.(float64) * oprnd2.(float64)
		default:
			erro = errors.New(errMsg)
		}
	default:
		erro = errors.New(errMsg)
	}
	return
}

// divider
func divider(oprnd1, oprnd2 Operand) (result Operand, erro error) {
	var errMsg = "Unknown type or uncomparable types of operands."
	switch oprnd1.(type) { // DOUBLE!!! type switch: OMG
	case int64:
		switch oprnd2.(type) {
		case int64:
			result = oprnd1.(int64) / oprnd2.(int64)
		default:
			erro = errors.New(errMsg)
		}
	case float64:
		switch oprnd2.(type) {
		case float64:
			result = oprnd1.(float64) / oprnd2.(float64)
		default:
			erro = errors.New(errMsg)
		}
	default:
		erro = errors.New(errMsg)
	}
	return
}

var Operation = map[Operator]func(Operand, Operand) (Operand, error){
	"+": summator,
	"-": differentiator,
	"*": multiplicator,
	"/": divider,
}

// checks correctness of the token and returns type of the token;
// it's not part of the NPN structure because this is very general function to
// serve other structures as well (for example Reverse Polish Notation)
func CheckToken(token Token) (typ tokenType, erro error) {
	switch token.(type) {
	case float64:
		if token.(float64) < float64(MAXVALUE) {
			typ, erro = OPERAND, nil
		} else {
			typ, erro = ERR, errors.New(fmt.Sprintf("Token value: `%f' overflow `%d'.",
				token.(float64), MAXVALUE))
		}
	case int64:
		if token.(int64) < MAXVALUE {
			typ, erro = OPERAND, nil
		} else {
			typ, erro = ERR, errors.New(fmt.Sprintf("Token value: `%d' overflow `%d'.",
				token.(int64), MAXVALUE))
		}
	case string:
		if _, ok := Operation[token.(string)]; !ok { // go idiom of checking presence
			typ, erro = ERR, errors.New(fmt.Sprintf("Token value: `%s' unknown operation.",
				token.(string)))
		} else {
			typ, erro = OPERATOR, nil
		}
	default:
		typ, erro = ERR, errors.New(fmt.Sprintf("Illegal token: `%s'.", token.(string)))
	}
	return
}

// Polish Notation calculator`s interface (for NPN and RPN alike)
type calculator interface {
	PopulateExpressionVector(string) error
	Calculate() (Operand, error)
}

// Normal Polish Notation (NPN): operator operand1 operand2
type NPN struct {
	calculator "embedded generic calculator interface"
	Expression genstack.Vector "tokenized expression generic vector"
	tokenStack genstack.Stack  "generic stack to store partial results"
}

//helper function for parsing user defined strings into proper base types
func parseString(stringElement string) (retElement genstack.Element) {
	if someInt, err := strconv.ParseInt(stringElement, 10, 64); err == nil {
		retElement = genstack.Element(someInt)
	} else if someFloat, err := strconv.ParseFloat(stringElement, 64); err == nil {
		retElement = genstack.Element(someFloat)
	} else { // default -> unrecognized type treated as string, LET THE CALLER DECIDE ABOUT IT
		retElement = genstack.Element(stringElement)
	}
	return
}

// extracts tokens from user input and puts them into generic vector of tokens
func (npn *NPN) PopulateExpressionVector(userInput string) (err error) {
	var tokens = strings.Fields(userInput)
	var tokenLen = len(tokens)
	if tokenLen < 3 { // need to have at least 3 tokens: one operator and 2 operands
		err = errors.New(fmt.Sprintf("Invalid expression: %s", userInput))
	} else {
		npn.Expression = make(genstack.Vector, tokenLen)
		for index, value := range tokens { // can NOT use copy???
			// WARNING: cast on PROPER BASE TYPE before recasting on
			//		    `genstack.Element' !!!
			npn.Expression[index] = genstack.Element(parseString(value))
		}
	}
	return
}

// calculates final result of the expression in the form of the polish notation
func (npn *NPN) Calculate() (retOperand Operand, erro error) {
	var exprLen = len(npn.Expression)
	var oprnd1, oprnd2, partialResult genstack.Element
	var token Token
	if exprLen == 0 {
		erro = errors.New("Expression vector is empty.")
		goto terminus
	}
	// utilize all elements of the expression vector
	for index := exprLen - 1; index >= 0; index-- {
		token = npn.Expression[index]
		typ, erro := CheckToken(token)
		switch typ {
		case OPERATOR:
			if oprnd1, erro = npn.tokenStack.Pop(); erro != nil {
				goto terminus
			}
			if oprnd2, erro = npn.tokenStack.Pop(); erro != nil {
				goto terminus
			}
			if partialResult, erro = Operation[token](oprnd1, oprnd2); erro != nil {
				goto terminus
			}
			npn.tokenStack.Push(partialResult)
		case OPERAND:
			npn.tokenStack.Push(token)
		case ERR:
			goto terminus // awful, awful, awful, pass error up
		}
	}
	// check if expression was malformed by looking into stack
	if npn.tokenStack.Len() > 1 { // malformed expression
		erro = errors.New("Expression malformed.")
	} else {
		retOperand = partialResult
	}
terminus:
	return
}

// factory function for NPN struct
func NewNPN(input string) (retNPN *NPN, erro error) {
	retNPN = new(NPN)
	if erro = retNPN.PopulateExpressionVector(input); erro == nil {
		retNPN.tokenStack = *genstack.NewStack()
	} else { // pass error up
		retNPN = nil
	}
	return
}
