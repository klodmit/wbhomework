package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	// Добавляем флаг для времени работы
	duration := flag.Int("time", 5, "Program duration in seconds")
	countWorkers := flag.Int("workers", 1, "Number of workers")
	flag.Parse()

	ch := make(chan int)

	// Создаем контекст с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(*duration)*time.Second)
	defer cancel()

	// Перехватываем сигналы прерывания
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		select {
		case <-sigCh:
			fmt.Println("\nReceived interrupt — shutting down...")
			cancel()
		case <-ctx.Done():
			// Контекст завершился по таймауту
		}
	}()

	wg := &sync.WaitGroup{}

	// Создаем воркеров
	for i := 0; i < *countWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for {
				select {
				case v, ok := <-ch:
					if !ok {
						return // канал закрыт
					}
					fmt.Printf("worker %d -> %d\n", workerID, v)
				case <-ctx.Done():
					return // контекст завершен
				}
			}
		}(i)
	}

	// Производитель
	go func() {
		i := 0
		for {
			select {
			case <-ctx.Done():
				return
			case ch <- i:
				i++
			}
		}
	}()

	// Ждем завершения контекста (по таймауту или отмене)
	<-ctx.Done()
	fmt.Printf("Time's up after %d seconds\n", *duration)

	// Закрываем канал и ждем завершения воркеров
	close(ch)
	wg.Wait()
}
