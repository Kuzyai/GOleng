package main

import "fmt"

type Text struct {
	On          bool
	Ammo, Power int
}

func (t *Text) Shoot() bool {
	if t.On == false {
		return false
	} else {
		if t.Ammo > 0 {
			t.Ammo--
			return true
		} else {
			return false
		}
	}
}
func (t *Text) RideBike() bool {
	if t.On == false {
		return false
	} else {
		if t.Power > 0 {
			t.Power--
			return true
		} else {
			return false
		}
	}
}
func main() {
	testStruct := new(Text)
	fmt.Scan(&testStruct.On, &testStruct.Ammo, &testStruct.Power)
	fmt.Println(testStruct.Shoot())
	fmt.Println(testStruct.RideBike())
}
