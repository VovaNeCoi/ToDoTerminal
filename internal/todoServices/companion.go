package todoServices

import (
	"bufio"
	"fmt"
	"os"
	models "projMod/internal/models"
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
		// выводит список
		var outputdata []models.FileDataStruct
		outputdata, err := ReadFromFile("")
		if err != nil {
			// TODO in this case you continue as nothing fucking error happened...
			fmt.Println(err)
		}
		for i := 0; i < len(outputdata); i++ {
			fmt.Printf("Number: %d\nTag: %sComment: %s",
				outputdata[i].TaskNum, outputdata[i].ShortName, outputdata[i].Comment)
		}
		// TODO use loop instead of zaloop
		Companion(num, shName, comm)

		// позволяет ввести строку: номер, тег, описание
	case "2":
		fmt.Print("Enter task number: ")
		fmt.Scanln(&num)

		fmt.Print("Enter short description: ")
		var err error
		shName, err = bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			// TODO same problem in err case
			fmt.Println(err)
		}

		fmt.Print("Enter detailed description: ")
		// TODO reuse err variable
		var errr error
		comm, errr = bufio.NewReader(os.Stdin).ReadString('\n')
		if errr != nil {
			// TODO oh God, third time???
			fmt.Println(errr)
		}

		AddTask(num, shName, comm)
		fmt.Println("Ok", num, shName, comm)
		Companion(num, shName, comm)

		// удаляет строку по номеру
	case "3":
		fmt.Print("Enter task number: ")
		fmt.Scanln(&num)
		err := DeleteTask(num)
		if err != nil {
			// TODO догадаешьсяы
			fmt.Println(err)
		}
		fmt.Println("Ok2", num, shName, comm)
		Companion(num, shName, comm)

		// выходит из списка
	case "4":
		break
	default:
		fmt.Println("Unknown command")
		return
	}
}
