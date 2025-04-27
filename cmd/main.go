// Вызывает работу программы и импортированных пакетов и функций
package main

import (
	functions "projMod/internal/todoServices"
)

func main() {
	var num int
	var shName, comm string
	functions.Companion(num, shName, comm)
}
