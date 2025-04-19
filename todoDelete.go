package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func DeleteTask(num int) error {
	var deserData []FileDataStruct

	// Чтение файла
	data, err := os.ReadFile("./OurToDoList.json")
	if err != nil {
		return fmt.Errorf("Ошибка чтения, %w", err)
	}

	// Десериализация
	err = json.Unmarshal(data, &deserData)
	if err != nil {
		return fmt.Errorf("Ошибка десериализации, %w", err)
	}

	// Добавить поиск из прочитанного десериализованного файла. Скорее всего
	// реализовать структуру хэш мапы????
	// Нужно обработать ошибки выхода за пределы массива
	for i, _ := range deserData {
		if deserData[i].TaskNum == num {
			deserData = append(deserData[:i], deserData[i+1:]...)
		}
	}

	return err
}

// Функция создания среза по i элемент и от i
// func createNewSlice(currSlice []FileDataStruct, index int) ([]FileDataStruct, error) {
// 	return (append(currSlice[0:index], currSlice[index:]...)), nil
// }
