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
	for i := 0; ; {

		if i > len(deserData) {
			break
		}

		// На счёт строчек ифа на строчках 34-36 не уверен, что оно вообще срабатывает
		if len(deserData) == 0 {
			break
		} else if deserData[i].TaskNum == num {
			if len(deserData) == 1 {
				deserData = nil // Возможны ошибки добавления нулевого значения в файл
			} else {
				deserData = append(deserData[:i], deserData[i+1:]...) // Возможна ошибка выхода за пределы массива
				i = 0
				fmt.Println(len(deserData), i)
				fmt.Println()
				continue
			}
		}
		i++
	}

	fmt.Println(deserData)
	fmt.Println()
	fmt.Println(len(deserData))

	// Запись в файл
	err = WriteIntoFile(deserData)
	if err != nil {
		return fmt.Errorf("Ошибка записи в файл, %w", err)
	}

	return err
}

// По поводу реализации через хэшмапу. Ощущение что придется писать много кода для тогоо,
// Чтобы сначала её создать, а потом её же вернуть в тип структуры
