package main

import "fmt"

func main() {
	var a, b float32
	var op string
	fmt.Scan(&a, &b, &op)

	switch op {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		if b == 0 {
			fmt.Println("Делить на ноль нельзя!")
		} else {
			fmt.Println(a / b)
		}
	case "%":
		fmt.Println(int(a) % int(b))
	}
}
