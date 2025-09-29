package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== ДЕМОНСТРАЦИЯ ВСЕХ СПОСОБОВ ОСТАНОВКИ ГОРУТИН ===\n")

	// 1. Выход по условию
	fmt.Println("1. Выход по условию:")
	stopFlag := false
	go func() {
		for !stopFlag {
			fmt.Print("⏰ ")
			time.Sleep(200 * time.Millisecond)
		}
		fmt.Println("\nГорутина остановлена по условию")
	}()
	time.Sleep(1 * time.Second)
	stopFlag = true
	time.Sleep(100 * time.Millisecond)

	// 2. Через канал уведомления
	fmt.Println("\n2. Через канал уведомления:")
	stopChan := make(chan struct{})
	go func() {
		for {
			select {
			case <-stopChan:
				fmt.Println("Горутина остановлена через канал")
				return
			default:
				fmt.Print("🔔 ")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()
	time.Sleep(1 * time.Second)
	close(stopChan)
	time.Sleep(100 * time.Millisecond)

	// 3. Использование контекста
	fmt.Println("\n3. Использование контекста:")

	// Контекст с таймаутом
	ctxTimeout, cancelTimeout := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelTimeout()

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Горутина остановлена по таймауту контекста")
				return
			default:
				fmt.Print("⏱️ ")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}(ctxTimeout)
	time.Sleep(2 * time.Second)

	// Контекст с отменой
	ctxCancel, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Горутина остановлена по отмене контекста")
				return
			default:
				fmt.Print("🎯 ")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}(ctxCancel)
	time.Sleep(1 * time.Second)
	cancel()
	time.Sleep(100 * time.Millisecond)

	// 4. Использование sync.WaitGroup
	fmt.Println("\n4. Использование sync.WaitGroup:")
	var wg sync.WaitGroup
	stopChan2 := make(chan struct{})

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-stopChan2:
				fmt.Println("Горутина завершена с WaitGroup")
				return
			default:
				fmt.Print("👥 ")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()
	time.Sleep(1 * time.Second)
	close(stopChan2)
	wg.Wait()

	// 5. Использование runtime.Goexit()
	fmt.Println("\n5. Использование runtime.Goexit():")
	go func() {
		defer fmt.Println("Горутина завершена через Goexit")

		ticker := time.NewTicker(200 * time.Millisecond)
		defer ticker.Stop()

		counter := 0
		for range ticker.C {
			fmt.Print("🚪 ")
			counter++

			if counter >= 5 {
				fmt.Println("\nВызываем runtime.Goexit()")
				runtime.Goexit()
			}
		}
	}()
	time.Sleep(2 * time.Second)

	// 6. Паника и восстановление
	fmt.Println("\n6. Паника и восстановление:")
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Восстановлены после паники: %v\n", r)
			}
		}()

		counter := 0
		for {
			fmt.Print("💥 ")
			time.Sleep(200 * time.Millisecond)
			counter++

			if counter >= 4 {
				panic("Искусственная паника для остановки")
			}
		}
	}()
	time.Sleep(2 * time.Second)

	// 7. Комбинированный подход с несколькими каналами
	fmt.Println("\n7. Комбинированный подход:")
	jobs := make(chan int, 5)
	stopChan3 := make(chan struct{})
	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			select {
			case job := <-jobs:
				fmt.Printf("📦[%d] ", job)
				time.Sleep(300 * time.Millisecond)
			case <-stopChan3:
				fmt.Println("\nКомбинированная горутина завершена")
				return
			}
		}
	}()

	// Отправляем задания
	for i := 1; i <= 3; i++ {
		jobs <- i
	}
	time.Sleep(1 * time.Second)
	close(stopChan3)
	<-done

	// 8. Использование таймера
	fmt.Println("\n8. Использование таймера:")
	go func() {
		timer := time.NewTimer(1 * time.Second)

		for {
			select {
			case <-timer.C:
				fmt.Println("\nТаймер сработал - останавливаем горутину")
				return
			default:
				fmt.Print("⏲️ ")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()
	time.Sleep(2 * time.Second)

	// 9. Закрытие канала данных
	fmt.Println("\n9. Закрытие канала данных:")
	dataChan := make(chan int)
	go func() {
		for data := range dataChan {
			fmt.Printf("📊[%d] ", data)
			time.Sleep(200 * time.Millisecond)
		}
		fmt.Println("\nКанал данных закрыт - горутина завершена")
	}()

	for i := 1; i <= 3; i++ {
		dataChan <- i
	}
	time.Sleep(500 * time.Millisecond)
	close(dataChan)
	time.Sleep(100 * time.Millisecond)

	fmt.Println("\n=== ВСЕ ДЕМОНСТРАЦИИ ЗАВЕРШЕНЫ ===")

	// Даем время для завершения всех горутин
	time.Sleep(1 * time.Second)
}
