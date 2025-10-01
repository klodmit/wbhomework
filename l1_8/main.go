package main

import "fmt"

func main() {
	// Объявление переменных
	var (
		a int64 = 0
		i int64 = 0
		b int64 = 0
	)
	fmt.Println("Введите число")
	_, err := fmt.Scan(&a)
	// Банальная обработка ошибок
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%064b", a)
	fmt.Println("\nВведите бит для изменения")
	// Банальная обработка ошибок
	_, err = fmt.Scan(&i)
	if err != nil {
		fmt.Println(err)
		return
	}
	i-- // Уменьшаем на 1 от значения которое будет введено так как отсчет начинается с 0
	fmt.Println("Укажите на что вы хотите поменять бит на 1 или 0")
	_, err = fmt.Scan(&b)
	if err != nil {
		fmt.Println(err)
		return
	}
	if b == 0 || b == 1 {
		a ^= 1 << i
		fmt.Println(a)
		fmt.Printf("%064b", a)
	} else {
		fmt.Println("ERROR")
	}

}
