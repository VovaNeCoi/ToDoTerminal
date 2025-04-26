// Вызывает работу программы и импортированных пакетов и функций
package main

import (
	"fmt"
	"log"
	functions "projMod/internal/todoServices"
)

func main() {

	fmt.Println("Hello, I am ToDo list")
	fmt.Println("How I can help you?")

	var num int
	var shName, comm string

	functions.Companion(num, shName, comm)

	var err error
	fmt.Println("Goodbye, I am ToDo list")

	// err = functions.AddTask(2, "Покакать", "Нужно открыть дверь, зайти в туалет, снять штаны, сделать 'дела', вытереть жопу, надеть штаны")
	// if err != nil {
	// 	log.Printf("Ошибка добавления: %s", err)
	// }

	err = functions.DeleteTask(3)
	if err != nil {
		log.Printf("Ошибка удаления: %v", err)
	}
}
