// Вызывает работу программы и импортированных пакетов и функций
package main

import (
	"fmt"
	"log"
)

func main() {
	var err error
	fmt.Println("Hello, I am ToDo list")

	// err = functions.AddTask(2, "Покакать", "Нужно открыть дверь, зайти в туалет, снять штаны, сделать 'дела', вытереть жопу, надеть штаны")
	// if err != nil {
	// 	log.Printf("Ошибка добавления: %s", err)
	// }

	err = functions.DeleteTask(3)
	if err != nil {
		log.Printf("Ошибка удаления: %v", err)
	}
}
