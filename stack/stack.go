package stack

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sp int
var stack [1000]int

func Pop() int {
	if sp > 0 {
		sp--
		return stack[sp]
	} else {
		fmt.Fprint(os.Stderr, "Невозможно выполнить POP для пустого стека.\n")
		return 0
	}
}

func Push(a int) {
	stack[sp] = a
	sp++
}

func Empty() bool {
	return sp == 0
}

func ProcessRPN(rpn string) int {
	tokens := strings.Fields(rpn)

	for _, t := range tokens {
		switch t {
		case "\n":
			break
		case " ":
			break
		case "=":
			fmt.Printf("Result = %d\n", Pop())
			break
		case "+":
			Push(Pop() + Pop())
		case "-":
			a := Pop()
			b := Pop()
			Push(b - a)
		case "*":
			Push(Pop() * Pop())
		case "/":
			a := Pop()
			b := Pop()
			Push(b / a)
		default:
			x, err := strconv.Atoi(t)
			if err != nil {
				fmt.Fprint(os.Stderr, "Can't read integer\n")
				os.Exit(-1)
			} else {
				Push(x)
			}
			break
		}
	}

	return Pop()
}

type Stack struct {
	data     []int
	last     int
	dataSize int
}

func newStack(initialSize int) *Stack {
	if initialSize <= 0 {
		initialSize = 100
	}
	return &Stack{
		data:     make([]int, initialSize),
		last:     0,
		dataSize: initialSize,
	}
}

func (s *Stack) deleteStack() {
	s.data = nil
	s.last = 0
	s.dataSize = 0
}

func (s *Stack) push(a int) {
	if s.last == s.dataSize {
		s.dataSize = (s.dataSize*3 + 1) / 2
		newData := make([]int, s.dataSize)
		copy(newData, s.data)
		s.data = newData
	}
	s.data[s.last] = a
	s.last++
}

func (s *Stack) pop() (int, error) {
	if s.last > 0 {
		s.last--
		return s.data[s.last], nil
	} else {
		return 0, errors.New("attempt to pop from an empty stack")
	}
}
