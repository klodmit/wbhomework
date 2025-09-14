package main

import (
	"flag"
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int) // создание канала всегда через make иначе он не инициализируется
	// Считываем флаг -workers по умолчанию создаем только одного
	countWorkers := flag.Int("workers", 1, "Number of workers")
	flag.Parse() // Парсим флаг

	// Создаем waitgroup чтоб main не завершился раньше горутин
	wg := &sync.WaitGroup{}

	// Создаем воркеров
	for i := 0; i < *countWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range ch { // читаем, пока канал не закроют
				fmt.Printf("worker %d -> %d\n", i, v)
			}
		}()
	}
	// Отправляем данные в канал
	for i := 0; i < 10000; i++ {
		ch <- i
	}
	// Закрываем канал
	close(ch)
	wg.Wait()
}
