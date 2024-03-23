package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var result float64
	value1, value2, operation := readTask()
	a, err := doFloat(value1)
	if err != nil {
		fmt.Println(err)
		return
	}
	b, err := doFloat(value2)
	if err != nil {
		fmt.Println(err)
		return
	}
	operation, err = doSign(operation)
	if err != nil {
		fmt.Println(err)
		return
	}
	switch operation {
	case "+":
		result = func(v1, v2 float64) float64 { return v1 + v2 }(a, b)
	case "-":
		result = func(v1, v2 float64) float64 { return v1 - v2 }(a, b)
	case "*":
		result = func(v1, v2 float64) float64 { return v1 * v2 }(a, b)
	case "/":
		result = func(v1, v2 float64) float64 { return v1 / v2 }(a, b)
	}
	fmt.Printf("%.4f", result)
}
func readTask() (interface{}, interface{}, interface{}) {
	var value1, value2, operation interface{}
	r := bufio.NewReader(os.Stdin)
	value1, _ = r.ReadString('\n')
	value2, _ = r.ReadString('\n')
	operation, _ = r.ReadString('\n')
	return value1, value2, operation
}
func doSign(i interface{}) (string, error) {
	if v, ok := i.(string); ok {
		v = strings.Trim(v, "\n")
		if v == "+" || v == "-" || v == "*" || v == "/" {
			return v, nil
		}
		return "", fmt.Errorf("неизвестная операция")
	}
	return "", fmt.Errorf("неизвестная операция")
}
func doFloat(i interface{}) (float64, error) {
	var s string
	if v, ok := i.(string); ok {
		s = v
	}
	s = strings.Trim(s, "\n")
	n, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, fmt.Errorf("value = %v: %T", n, n)
	}
	return n, nil
}
