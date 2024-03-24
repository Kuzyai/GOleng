package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	var sum int
	s := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	for s.Scan() {
		num, err := strconv.Atoi(s.Text())
		if err != nil {
			break
		}
		sum += num
	}
	_, err := w.WriteString(strconv.Itoa(sum))
	if err != nil {
		panic(err)
	}
	w.Flush()
}
