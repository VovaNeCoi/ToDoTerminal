// Функция добавления новых данных в файл. Функция получает номер задачи,
// краткое название и комментарий с пояснением. Дальше можно накинуть проверку
// существования таски
package functionality

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Создаем структуру для хранения данных в определенном виде. Пока неясно будет ли вообще использоваться
// Пока не использую
type fileDataStruct struct {
	taskNum   int
	shortName string
	comment   string
}

// Добавляем созданные данные в файл. Каждое новое добавление должно вызывать функцию Add
func Add(num int, shName, comm string) {
	dataForSerialize := &fileDataStruct{taskNum: num, shortName: shName, comment: comm}

	// Сериализовываем данные
	dataForFile, err := json.Marshal(dataForSerialize)
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
	err = ioutil.WriteFile(f.Name(), dataForFile, 0644)
	if err != nil {
		fmt.Println("Запись в файл", err)
		return
	}
}
