package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	ch := make(chan int) // создание канала всегда через make иначе он не инициализируется
	// Считываем флаг -workers по умолчанию создаем только одного
	countWorkers := flag.Int("workers", 1, "Number of workers")
	flag.Parse() // Парсим флаг

	// контекст для отмены всех горутин
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// перехватываем сигналы прерывания и посылаем cancel()
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigCh
		fmt.Println("\nReceived interrupt — shutting down...")
		cancel()
		// можно stop-нуть signal.Notify, но defer cancel уже есть
	}()

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
	// производитель: отправляет данные, но уважает ctx отмену
producer:
	for i := 0; i < 10000000; i++ {
		select {
		case <-ctx.Done():
			// получили отмену — прекращаем отправлять
			break producer
		case ch <- i:
			// отправлено
		}
	}
	// Закрываем канал
	close(ch)
	wg.Wait()
}
