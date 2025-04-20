// Функция добавления новых данных в файл. Функция получает номер задачи int,
// краткое название string и комментарий с пояснением string. Дальше можно накинуть проверку
// существования таски
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Создаем структуру для хранения данных
type FileDataStruct struct {
	TaskNum   int    `json:"Task Number"`
	ShortName string `json:"Name"`
	Comment   string `json:"Poxyi"`
}

// Название файла с данными
const FileName = "OurToDoList.json"

// Добавляем созданные данные в файл.
// Каждое новое добавление должно вызывать функцию Add
func AddTask(num int, shName, comm string) error {

	// Создаю обычный json файл С проверкой существования файла
	if _, err := os.Stat(FileName); os.IsNotExist(err) {
		_, err := os.Create(FileName)
		if err != nil {
			return fmt.Errorf("todoAdd Ошибка создания файла, %w", err)
		}
	}

	// Считываем данные из файла и возвращаем срез
	dataForSerialize, err := ReadFromFile()
	if err != nil {
		return fmt.Errorf("todoAdd Ошибка чтения файла, %w", err)
	}

	// Собственно сам процесс добавления данных к уже прочитанным
	newAddedData := &FileDataStruct{TaskNum: num, ShortName: shName, Comment: comm}
	dataForSerialize = append(dataForSerialize, *newAddedData)

	// Вызов функции записи даннных в файл
	WriteIntoFile(dataForSerialize)

	return nil
}

// Функция чтения данных из файла.
// Возвращает сразу срез данных созданной структуры с данными из файла
func ReadFromFile() ([]FileDataStruct, error) {
	// Ловим ошибку пустого файла
	size, err := os.Stat(FileName)
	if err != nil {
		return nil, fmt.Errorf("toDoRead Ошибка статов, %w", err)
	}
	if size.Size() == 0 {
		return nil, nil
	}

	var fileData []FileDataStruct

	// Считываем байты из файла
	tmpFileData, err := os.ReadFile(FileName)
	if err != nil {
		return nil, fmt.Errorf("toDoRead Ошибка чтения %w", err)
	}

	// Десериализуем данные из байтового формата в FileDataStruct
	err = json.Unmarshal(tmpFileData, &fileData)
	if err != nil {
		return nil, fmt.Errorf("toDoRead Ошибка десериализации %w", err)
	}

	return fileData, err
}

// Функция записи данных в файл
// Принимает слайс из данных типа структуры. Возвращает данные ошибки
// Сам сериализовывает полученный слайс
func WriteIntoFile(data []FileDataStruct) error {
	// Сериализовываем данные. Они хранятся в типе байт
	dataForFile, err := json.MarshalIndent(data, " ", " ")
	if err != nil {
		return fmt.Errorf("todoWrite Ошибка сериализации, %w", err)
	}

	// Запись данных в файл
	err = os.WriteFile(FileName, dataForFile, 0644)
	if err != nil {
		return fmt.Errorf("todoWrite Ошибка записи в файл, %w", err)
	}

	return nil
}

// Для дальнейшего масштабирования. Можно номер задачи задавать не вручную, а автоматически
// Можно добавить шаги для задач, их нумерацию и дать возможность удалять их и добавлять

// Вопрос на засыпку, зачем было реализовывать свои функции чтения, записи, удаления и добавления данных?
