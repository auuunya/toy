package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var s string
	flag.StringVar(&s, "b", "", "binary is 0 or 1")
	flag.Parse()
	if s == "" {
		fmt.Printf("please input binary\n")
		os.Exit(0)
	}
	llen := len(s)
	var bina []int
	for _, r := range s {
		n, ok := isBianry(r)
		if !ok {
			fmt.Printf("input invalid\n")
			os.Exit(0)
		}
		bina = append(bina, n)
	}
	var sum int
	for i, num := range bina {
		tmp := num * int(math.Pow(float64(2), float64(llen-i-1)))
		sum += int(tmp)
	}
	fmt.Printf("binary is: %v, convert to decimal: %v\n", s, sum)
}

func isBianry(s rune) (int, bool) {
	i, _ := strconv.Atoi(string(s))
	return i, i == 0 || i == 1
}
