// Includes functions: AddTask(int, string, string), DeleteTask(int), WriteIntoFile([]FileDataStruct), ReadFromFile()
package todoServices

import (
	"fmt"
	"os"
	fileChecker "projMod/internal/helpfullFunctions"
	fileStruct "projMod/internal/models"
)

// Добавляем созданные данные в файл. Можно указать конкретный файл, иначе будет
// задан по дефолту "OurToDoList.json"
// Каждое новое добавление должно вызывать функцию Add
func AddTask(num int, shName, comm string, fileName ...string) error {
	currFileName, err := fileChecker.CheckFileName(fileName...)
	if err != nil {
		return fmt.Errorf("todoAdd ошибка названия файла, %w", err)
	}

	// Создаю обычный json файл С проверкой существования файла
	// Надо задать в какой директории создается файл...
	if _, err := os.Stat(currFileName); os.IsNotExist(err) {
		_, err := os.Create(currFileName)
		if err != nil {
			return fmt.Errorf("todoAdd ошибка создания файла, %w", err)
		}
	}

	// Считываем данные из файла и возвращаем срез
	dataForSerialize, err := ReadFromFile(currFileName)
	if err != nil {
		return fmt.Errorf("todoAdd ошибка чтения файла, %w", err)
	}

	// Собственно сам процесс добавления данных к уже прочитанным
	newAddedData := &fileStruct.FileDataStruct{TaskNum: num, ShortName: shName, Comment: comm}
	dataForSerialize = append(dataForSerialize, *newAddedData)

	// Вызов функции записи даннных в файл
	WriteIntoFile(dataForSerialize, currFileName)

	return nil
}
