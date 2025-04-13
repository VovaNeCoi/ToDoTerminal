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
type FileDataStruct struct {
	taskNum   int
	shortName string
	comment   string
}

// Добавляем созданные данные в файл. Каждое новое добавление должно вызывать функцию Add
func Add(num int, shName, comm string) {
	dataForSerialize := &FileDataStruct{taskNum: num, shortName: shName, comment: comm}

	// Сериализовываем данные
	dataForFile, err := json.MarshalIndent(dataForSerialize, " . ", " * ")
	if err != nil {
		fmt.Println("В сериализации ошибка", err)
		return
	}

	// Создаю обычный json файла
	f, err := os.Create("OurToDoList.json")
	if err != nil {
		fmt.Println("Создание файла", err)
		return
	}

	// Нужно реализовать правильную обработку ошибок через panic и вывод ошибки в виде перменной err
	err = os.WriteFile(f.Name(), dataForFile, 0644)
	if err != nil {
		fmt.Println("Запись в файл", err)
		return
	}
}
