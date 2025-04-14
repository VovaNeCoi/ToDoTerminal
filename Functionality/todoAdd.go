// Функция добавления новых данных в файл. Функция получает номер задачи int,
// краткое название string и комментарий с пояснением string. Дальше можно накинуть проверку
// существования таски
package functionality

import (
	"encoding/json"
	"fmt"
	"os"
)

// Создаем структуру для хранения данных в определенном виде.
// Добавить пометки для полей
type FileDataStruct struct {
	TaskNum   int
	ShortName string
	Comment   string
}

// Добавляем созданные данные в файл. Каждое новое добавление должно вызывать функцию Add
func Add(num int, shName, comm string) {
	dataForSerialize := &FileDataStruct{TaskNum: num, ShortName: shName, Comment: comm}

	// Сериализовываем данные. Они хранятся в типе байт
	dataForFile, err := json.Marshal(dataForSerialize)
	if err != nil {
		fmt.Println("В сериализации ошибка", err)
		return
	}

	// Название файла
	fileName := "OurToDoList.json"
	// Создаю обычный json файл
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
