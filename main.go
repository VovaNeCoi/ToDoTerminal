// Вызывает работу программы и импортированных пакетов и функций
package main

import (
	"fmt"

	"./functionality"
)

func main() {
	fmt.Println("Hello, I am ToDo list")

	functionality.Add(1, "Покакать", "Нужно открыть дверь, зайти в туалет, снять штаны, сделать 'дела', вытереть жопу, надеть штаны")
}
