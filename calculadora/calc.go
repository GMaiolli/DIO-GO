package main

import "fmt"

func main() {
	x := sum(1, 3, 6)
	y := mult(6, 7)
	w := sub(5, 9)
	z := div(67)
	fmt.Println(x, y, w, z)
}

func sum(i ...int) int {
	res := 0
	for _, v := range i {
		res += v
	}
	return res
}

func sub(i ...int) int {
	if len(i) == 0 {
		return 0
	}
	
	res := i[0] 
	
	for _, v := range i[1:] {
		res -= v
	}
	return res
}

func mult(i ...int) int {
	res := 1
	for _, v := range i {
		res = res * v
	}
	return res
}

func div(i ...int) int {
	if len(i) == 0 {
		return 0
	}
	res := i[0]
	
	for _, v := range i[1:] {
		if v == 0 {
			panic("divisão por zero!")
		}
		res = res / v
	}
	return res
}