package main

import "fmt"

func main() {
	var persentMonth1, persentMonth2, sum1, sum2 float64
	var cost1 float64 = 720000
	var cost2 float64 = 1047400
	var persent1 float64 = 0.3499
	var persent2 float64 = 0.1499
	for i := 1; i <= 60; i++ {
		persentMonth1 = (cost1 * persent1) / 12
		persentMonth2 = (cost2 * persent2) / 12
		cost1 += persentMonth1
		cost2 += persentMonth2
		cost1 -= 25548
		cost2 -= 24912
		sum1 += persentMonth1
		sum2 += persentMonth2
		fmt.Println(i, fmt.Sprintf("%.2f", cost1), fmt.Sprintf("%.2f", cost2))
	}

	fmt.Println(" ")
	fmt.Println(fmt.Sprintf("%.2f", sum1+720000), fmt.Sprintf("%.2f", sum2+1047400))
}
