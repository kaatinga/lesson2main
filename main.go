package main

import (
	"log"

	lesson2 "github.com/kaatinga/lesson2package"
)

var input int64 // the variable to keep input option

var tasks [4]string // the array of tasks will be here

func main() {
	p := log.Println // the alias for log.Print in order to simplify the code

	tasks[0] = "1. Написать функцию, которая определяет, четное ли число"
	tasks[1] = "2. Написать функцию, которая определяет, делится ли число без остатка на 3"
	tasks[2] = "3. Написать функцию, которая последовательно выводит на экран 100 первых чисел Фибоначчи, начиная от 0"
	tasks[3] = "4. Заполнить массив из 100 элементов различными простыми числами"

	p("Hello. All the homework tasks are resolved in one app")
	for {
		p("The list of tasks:")
		for _, value := range tasks {
			p(value)
		}
		p("Choose the task.")

		input = lesson2.UserInput(1, 4)
		p("You have choosen the option", input, ":", tasks[input-1])
		switch input {
		case 1:
			lesson2.IfEven()
		case 2:
			lesson2.IfDevisionHasRemainder()
		case 3:
			p("The list of the first 100 Fibonacci Numbers:")
			lesson2.Fibonacci(0, 0, 1)
		case 4:
			lesson2.Sieve(541)
		}
	}
}
