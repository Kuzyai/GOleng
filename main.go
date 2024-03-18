package main

import "fmt"

func main() {
	var x, y, n, m, i1, i2, z uint16
	var counter1, counter2, counter3, q byte
	if x <= 10000 && y <= 10000 {
		for fmt.Scan(&x); ; fmt.Scan(&x) {
			var c byte = 0
			i1 = x
			i2 = x
			for i1 > 0 {
				n = i1 % 10
				i1 /= 10
				counter1++
				counter3++
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
			} else {
				counter3 = 0
			}
		}
		for fmt.Scan(&y); ; fmt.Scan(&y) {
			var c byte = 0
			i1 = y
			i2 = y
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
				i2 = y
			}
			counter1 = 0
			if c != 1 {
				break
			}
		}
		i1 = x
		i2 = y
		z = x
		for ; q < counter3; q++ {
			i1 = z
			for i1 > 0 {
				n = i1 % 10
				counter1++
				if i1 < 10 {
					for i2 > 0 {
						m = i2 % 10
						i2 /= 10
						if n == m {
							fmt.Printf("%d ", n)
						}
					}
					i2 = y
				}
				i1 /= 10
			}
			i1 = z
			switch counter1 {
			case 2:
				z = i1 % 10
			case 3:
				z = i1 % 100
			case 4:
				z = i1 % 1000
			default:
				z = 0
			}
			counter1 = 0
		}
	}
}
