package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type Stack struct {
	tokens []string
}

func (s *Stack) IsEmpty() bool {
	return len(s.tokens) == 0
}

func (s *Stack) Push(token string) {
	s.tokens = append(s.tokens, token)
}

func (s *Stack) Pop() string {
	if s.IsEmpty() {
		panic("stack empty, cannot pop")
	}
	ret := s.tokens[len(s.tokens)-1]
	s.tokens = s.tokens[:len(s.tokens)-1]
	return ret
}

func (s *Stack) Top() string {
	if s.IsEmpty() {
		panic("stack empty, cannot get top")
	}
	return s.tokens[len(s.tokens)-1]
}

func parserStrToToken(s string) (ret []string) {
	for _, ch := range ([]byte)(s) {
		fmt.Println(ch)
	}
	return
}

func main() {
	exp := flag.String("e", "", "输入表达式, 如: 1+23*(45-6)")
	// Todo: 词法解析
	tokens := parserStrToToken(*exp)

	tokens = []string{
		// "1+2*((3+4)*5+6)+7"
		"1", "+", "2", "*", "(", "(", "3", "+", "4", ")", "*", "5", "+", "6", ")", "+", "7",
	}
	// tokens = []string{
	// 	// ((1+2)+3*4)-5
	// 	"(", "(", "1", "+", "2", ")", "+", "3", "*", "4", ")", "-", "5",
	// }

	numStack := Stack{tokens: make([]string, 0)}
	opStack := Stack{tokens: make([]string, 0)}

	// Todo: 优化逻辑
	//		 除法实现，非整数
	for _, t := range tokens {
		// fmt.Printf("OP: %v\tNUM: %v\n", opStack, numStack)
		switch t {
		case "+", "-":
			if !opStack.IsEmpty() {
				if op := opStack.Top(); op == "+" || op == "-" || op == "*" || op == "/" {
					v1, _ := strconv.ParseInt(numStack.Pop(), 10, 32)
					v2, _ := strconv.ParseInt(numStack.Pop(), 10, 32)
					var v3 int64
					switch op {
					case "+":
						v3 = v2 + v1
					case "-":
						v3 = v2 - v1
					case "*":
						v3 = v2 * v1
					case "/":
						v3 = v2 / v1
					}
					numStack.Push(strconv.FormatInt(v3, 10))
					opStack.Pop()
				}
			}
			opStack.Push(t)
		case "*", "/":
			if !opStack.IsEmpty() {
				if op := opStack.Top(); op == "*" || op == "/" {
					v1, _ := strconv.ParseInt(numStack.Pop(), 10, 32)
					v2, _ := strconv.ParseInt(numStack.Pop(), 10, 32)
					var v3 int64
					if op == "*" {
						v3 = v2 * v1
					} else {
						v3 = v2 / v1
					}
					numStack.Push(strconv.FormatInt(v3, 10))
					opStack.Pop()
				}
			}
			opStack.Push(t)
		case "(":
			opStack.Push(t)
		case ")":
			op := opStack.Pop()
			for ; op != "("; op = opStack.Pop() {
				v1, _ := strconv.ParseInt(numStack.Pop(), 10, 32)
				v2, _ := strconv.ParseInt(numStack.Pop(), 10, 32)
				var v3 int64
				switch op {
				case "+":
					v3 = v2 + v1
				case "-":
					v3 = v2 - v1
				case "*":
					v3 = v2 * v1
				case "/":
					v3 = v2 / v1
				}
				numStack.Push(strconv.FormatInt(v3, 10))
			}
		default:
			numStack.Push(t)
		}
	}

	for !opStack.IsEmpty() {
		op := opStack.Pop()
		v1, _ := strconv.ParseInt(numStack.Pop(), 10, 32)
		v2, _ := strconv.ParseInt(numStack.Pop(), 10, 32)
		var v3 int64
		switch op {
		case "+":
			v3 = v2 + v1
		case "-":
			v3 = v2 - v1
		case "*":
			v3 = v2 * v1
		case "/":
			v3 = v2 / v1
		}
		numStack.Push(strconv.FormatInt(v3, 10))
	}

	fmt.Printf("%s = %s\n", strings.Join(tokens, ""), numStack.Top())
}
