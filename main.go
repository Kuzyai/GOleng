package main

import (
	"fmt"
	"math"
)

type S struct {
	n, k, p, v, m, w, t float64
}

func main() {
	ex := &S{}
	fmt.Scan(&ex.k, &ex.p, &ex.v, &ex.n)
	ex.M()
	ex.W()
	ex.T()
	fmt.Println(ex.t)
}

func (s *S) M() {
	s.m = s.v * s.p
}

func (s *S) W() {
	s.w = math.Sqrt(s.k / s.m)
}

func (s *S) T() {
	s.t = s.n / s.w
}
