package main

import (
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
	return 0
}

func main() {
	//scanner := bufio.NewScanner(os.Stdin)
	//
	//for scanner.Scan() {
	//	text := strings.TrimSpace(scanner.Text())
	//
	//	if len(text) == 0 {
	//		continue
	//	}
	//	ProcessRPN(text)
	//}
	var num int
	fmt.Scan(&num)
	firstNum := num / 1000
	lastNum := num % 1000
	a := firstNum % 10
	b := firstNum % 100 / 10
	c := firstNum % 1000 / 100
	sum := a + b + c
	a1 := lastNum % 10
	b1 := lastNum % 100 / 10
	c1 := lastNum % 1000 / 100
	sum1 := a1 + b1 + c1
	fmt.Println(sum)
	fmt.Println(sum1)
	if sum == sum1 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
	//if b != a && c != a && b != c {
	//	fmt.Println("YES")
	//} else {
	//	fmt.Println("NO")
	//}
	//testCases := []string{
	//	"()",
	//	"()[]{}",
	//	"(]",
	//	"([)]",
	//	"{[]}",
	//	")",
	//}
	//
	//for _, testCase := range testCases {
	//	fmt.Printf("Input: \"%s\": %v\n", testCase, isValid(testCase))
	//}
}

func isValid(s string) bool {
	var stack []rune

	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 {
				return false
			}
			var bracket rune
			if char == ')' {
				bracket = '('
			} else if char == '}' {
				bracket = '{'
			} else if char == ']' {
				bracket = '['
			}
			if stack[len(stack)-1] == bracket {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if index, ok := numMap[complement]; ok {
			return []int{index, i}
		}
		numMap[num] = i
	}

	return nil
}

func exp() {
	var number int
	fmt.Scan(&number)
	a := number / 100
	number -= a * 100
	b := number / 10
	number -= b * 10
	c := number
	if b != a && c != a && b != c {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

	//var a int
	//fmt.Scan(&a)
	//switch {
	//case a/10000 >= 1:
	//	fmt.Println(a / 10000)
	//case a/1000 >= 1:
	//	fmt.Println(a / 1000)
	//case a/100 >= 1:
	//	fmt.Println(a / 100)
	//case a/10 >= 1:
	//	fmt.Println(a / 10)
	//default:
	//	fmt.Println(a)
	//}
}
