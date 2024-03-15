package main

import "fmt"

type Point struct {
	X, Y int
}

func main() {
	pointsMap := map[string]Point{}
	otherPointsMap := make(map[string]Point)
	//var oneMorePointsMap map[string]Point
	pointsMap["a"] = Point{X: 1, Y: 2}
	otherPointsMap["a"] = Point{1, 2}
	fmt.Println(pointsMap)
	fmt.Println(otherPointsMap)
	fmt.Println(pointsMap["a"])
	fmt.Println(otherPointsMap["a"])
}
