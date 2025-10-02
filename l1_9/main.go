package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ch1 := make(chan int)
	ch2 := make(chan int)
	// 1. Генератор чисел (пишет в ch1)
	go func() {
		for _, v := range arr {
			ch1 <- v
		}
		close(ch1) // когда массив закончился, закрываем канал
	}()

	// 2. Обработчик чисел (читаем из ch1, пишем в ch2)
	go func() {
		for x := range ch1 {
			ch2 <- x * 2
		}
		close(ch2) // когда ch1 закрылся и цикл завершился, закрываем ch2
	}()

	// 3. Чтение из второго канала в main
	for result := range ch2 {
		fmt.Println(result)
	}
}
