package main

import (
	"fmt"
	"sync"
)

/*
	Написать программу, которая конкурентно рассчитает значения квадратов чисел,
	взятых из массива [2,4,6,8,10], и выведет результаты в stdout.
*/

func main() {
	arr := [5]int{2, 4, 6, 8, 10}
	wg := new(sync.WaitGroup)

	for i := 0; i < len(arr); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			arr[i] = arr[i] * arr[i]
		}()
	}
	wg.Wait()
	fmt.Println(arr)
}
