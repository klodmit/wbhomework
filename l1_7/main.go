package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.RWMutex   // Используем RWMutex для разделения блокировок на чтение и запись
	var wg sync.WaitGroup // Создаем группу ожидания
	data := make(map[int]int)

	// Запускаем 1000 горутин для записи
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func(key int) { // Передаем i как аргумент, чтобы каждая горутина получила свое значение
			defer wg.Done() // Сообщаем группе о завершении при выходе из функции

			mu.Lock() // Блокируем на запись
			data[key]++
			fmt.Println("key:", key, "value:", data[key])
			mu.Unlock() // Снимаем блокировку
		}(i) // Явно передаем текущее значение i в горутину
	}

	// Ожидаем завершения всех горутин для записи
	wg.Wait()

	// Теперь можно безопасно читать, так как все горутины-писатели завершились
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func(key int) {
			defer wg.Done()
			mu.RLock() // Блокируем на чтение (можно использовать несколько RLock одновременно)
			defer mu.RUnlock()
			fmt.Println("Результат чтения:", key, "value:", data[key])
		}(i)
	}

	wg.Wait() // Ожидаем завершения горутины для чтения
}
