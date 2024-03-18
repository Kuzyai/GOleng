package main

import "fmt"

func main() {
	var x, y, n, m, i1, i2 uint16
	var counter1, counter2 byte
	if x <= 10000 && y <= 10000 {
		for fmt.Scan(&x); ; fmt.Scan(&x) {
			var c byte = 0
			i1 = x
			i2 = x
			for i1 > 0 {
				n = i1 % 10
				i1 /= 10
				counter1++
				for i2 > 0 {
					m = i2 % 10
					i2 /= 10
					counter2++
					if n == m && counter1 != counter2 {
						c = 1
					}
				}
				counter2 = 0
				i2 = x
			}
			counter1 = 0
			if c != 1 {
				break
			}
		}
	}
}
