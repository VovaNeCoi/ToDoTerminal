// Функция добавления новых данных в файл. Функция получает номер задачи int,
// краткое название string и комментарий с пояснением string. Дальше можно накинуть проверку
// существования таски
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Создаем структуру для хранения данных в определенном виде.
// Добавить пометки для полей
type FileDataStruct struct {
	TaskNum   int    `json:"Номер задачи"`
	ShortName string `json:"Название"`
	Comment   string `json:"Комментарий"`
}

// Добавляем созданные данные в файл. Каждое новое добавление должно вызывать функцию Add
func Add(num int, shName, comm string) {
	// Название файла
	fileName := "OurToDoList.json"

	// Нужно читать данные с файла и к ним добавлять новые **Добавить обработку ошибки нулевого файла
	tmp, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Ошибка чтения файла, ", err)
	}

	// Создается слайс с типом данных FileDataStruct
	var dataForSerialize []FileDataStruct
	json.Unmarshal(tmp, &dataForSerialize)

	newData := &FileDataStruct{TaskNum: num, ShortName: shName, Comment: comm}
	dataForSerialize = append(dataForSerialize, *newData)

	fmt.Println(dataForSerialize)

	// Сериализовываем данные. Они хранятся в типе байт
	dataForFile, err := json.MarshalIndent(dataForSerialize, " ", " ")
	if err != nil {
		fmt.Println("В сериализации ошибка", err)
		return
	}

	// Создаю обычный json файл С проверкой существования файла
	if _, err := os.Stat("OurToDoList.json"); os.IsNotExist(err) {
		_, err := os.Create("OurToDoList.json")
		if err != nil {
			fmt.Println("Создание файла", err)
			return
		}
	}

	// Нужно реализовать правильную обработку ошибок через panic и вывод ошибки в виде перменной err
	err = os.WriteFile(fileName, dataForFile, 0644)
	if err != nil {
		fmt.Println("Запись в файл", err)
		return
	}
}

// Для дальнейшего масштабирования. Можно номер задачи задавать не вручную, а автоматически
// Можно добавить шаги для задач, их нумерацию и дать возможность удалять их и добавлять
