package main

import "fmt"

func main() {
	arr := [...]float32{-30, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := make(map[int][]float32)

	for _, temp := range arr {
		key := int(temp) / 10 * 10
		groups[key] = append(groups[key], temp)
	}

	fmt.Println(groups)
}
