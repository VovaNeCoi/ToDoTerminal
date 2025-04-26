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
}
