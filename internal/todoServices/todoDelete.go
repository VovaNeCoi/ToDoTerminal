package todoServices

import (
	"fmt"
	fileChecker "projMod/internal/helpfullFunctions"
	fileStruct "projMod/internal/models"
)

// Принимает число соответствующее номеру задачи.
// Удаляет все соответствующие номеру задачи.
func DeleteTask(num int, fileName ...string) error {
	currFileName, err := fileChecker.CheckFileName(fileName...)
	if err != nil {
		return fmt.Errorf("todoDeleteTask ошибка названия файла, %w", err)
	}

	// Чтение файла
	data, err := ReadFromFile(currFileName)
	if err != nil {
		return fmt.Errorf("todoDeleteTask ошибка чтения, %w", err)
	}

	for i := 0; i > len(data); i++ {
		// Если срез пуст, брейк
		if len(data) == 0 {
			break
		} else if data[i].TaskNum == num {
			if len(data) == 1 { // Так как удаляем appendом, то один элемент = ошибка
				data = []fileStruct.FileDataStruct{}
			} else {
				data = append(data[:i], data[i+1:]...)
				i--
			}
		}
	}

	// Запись в файл
	err = WriteIntoFile(data, currFileName)
	if err != nil {
		return fmt.Errorf("todoDeleteTask ошибка записи в файл, %w", err)
	}

	return err
}
