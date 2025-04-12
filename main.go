// Вызывает работу программы и импортированных пакетов и функций
package main

import "fmt"

func main() {
	arr := "helo435ohMygod./m,*79384)()(_+=`)"

	for i := range arr {
		if arr[i] >= '0' && arr[i] <= '9' {
			fmt.Print(string(arr[i]))
		}
	}

	fmt.Println()
	fmt.Println("Hello, I am ToDo list")
<<<<<<< HEAD

	//
	fmt.Print("Something")
=======
	fmt.Println("/")
>>>>>>> 7790a2219f1bd6d3a7f37e399373fc4a4db489f2
}
