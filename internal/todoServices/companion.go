package todoServices

import (
	"fmt"
)

// UI для пользователя
func Companion(num int, shName string, comm string) {
	var operator string
	fmt.Println("Choose command (write number):")
	fmt.Println("1. Output ToDo list")
	fmt.Println("2. Add line")
	fmt.Println("3. Delete line")
	fmt.Println("4. Exit")
	fmt.Scanln(&operator)
	switch operator {
	case "1":
		// читает, но не выводит (нужно создать перменную куда будет выводиться чтение)
		ReadFromFile("")
		Companion(num, shName, comm)
	case "2":
		fmt.Print("Enter task number: ")
		fmt.Scanln(&num)
		fmt.Print("Enter short description: ")
		fmt.Scanln(&shName)
		// не позволяет ввести описание, сразу выводит ок || не понимает пробел
		fmt.Print("Enter detailed description: ")
		fmt.Scanln(&comm)
		AddTask(num, shName, comm)
		fmt.Println("Ok", num, shName, comm)
		Companion(num, shName, comm)
	case "3":
		// не работает удаление
		fmt.Print("Enter task number: ")
		fmt.Scanln(&num)
		err := DeleteTask(num)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Ok2", num, shName, comm)
		Companion(num, shName, comm)
	case "4":
		break
	default:
		fmt.Println("Unknown command")
		return
	}
}
