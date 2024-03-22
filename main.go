package main

import "fmt"

func main() {
	cityPopulation := map[string]int{
		"Ryazan":     528,
		"Novouralsk": 78,
		"Tomsk":      556,
		"Kemerovo":   557,
	} //город: население города в тысячах человек,
	groupCity := map[int][]string{
		10:   {"Kuznetsk", "Bor", "Novouralsk"},             // города с населением 10-99 тыс. человек
		100:  {"Ryazan", "Tomsk", "Kemerovo"},               // города с населением 100-999 тыс. человек
		1000: {"Moscow", "Saint Petersburg", "Novosibirsk"}, // города с населением 1000 тыс. человек и более
	}
	for key, _ := range cityPopulation {
		for key1, value1 := range groupCity {
			for _, i := range value1 {
				if key == i && key1 != 100 {
					delete(cityPopulation, key)
				}
			}
		}
	}
	fmt.Println(cityPopulation)
}
