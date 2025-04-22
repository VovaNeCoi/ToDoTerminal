// Вызывает работу программы и импортированных пакетов и функций
package main

import (
	"fmt"
	"log"
	functions "projMod/internal/todoServices"
)

func companion() {
	var operator string
	fmt.Println("Choose command (write number):")
	fmt.Println("1. Output ToDo list")
	fmt.Println("2. Add line")
	fmt.Println("3. Delete line")
	fmt.Println("4. Exit")
	fmt.Scanln(&operator)
	switch operator {
	case "1":
		functional.ReadFromFile()
	case "2":
		functional.AddTask()
	case "3":
		functional.DeleteTask()
	case "4":
		break
	default:
		fmt.Println("Unknown command")
		return
	}

func main() {

	fmt.Println("Hello, I am ToDo list")
	fmt.Println("How I can help you?")

	var TaskNum int
	var ShortName, comm string
	
	var err error
	fmt.Println("Hello, I am ToDo list")

	err = functions.AddTask(1, "Покакать", "Нужно открыть дверь, зайти в туалет, снять штаны, сделать 'дела', вытереть жопу, надеть штаны")
	if err != nil {
		log.Printf("Ошибка добавления: %s", err)
	}

	err = functions.DeleteTask(1)
	if err != nil {
		log.Printf("Ошибка удаления: %v", err)
	}
}
