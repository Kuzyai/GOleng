package main

import "fmt"

func main() {
	var grad, hour, minutes, indexHour, indexMinutes, i uint16
	indexHour = 360 / 12
	indexMinutes = 12 * 60 / 360
	fmt.Scan(&grad)
	if grad <= 360 {
		hour = grad / indexHour
		i = grad - hour*indexHour
		minutes = i * indexMinutes
		fmt.Println("It is", hour, "hours", minutes, "minutes.")
	}
}
