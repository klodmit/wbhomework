package main

import "fmt"

// Human Структура человек имеет поле name (имя)
type Human struct {
	name string
}

// Name возвращает имя
func (h Human) Name() string {
	return h.name
}

// Action структура которая наследует поля и методы Human
type Action struct {
	Human
}

// SayName собственный метод структуры Action
func (a Action) SayName() string {
	return "Hello, my name is " + a.name // Напрмяую обращаюсь к полю name в структуре Human
}

func main() {
	// Заполняем структуру
	action := Action{Human{"John"}}
	fmt.Println(action.SayName())
	fmt.Println(action.Name())
}
