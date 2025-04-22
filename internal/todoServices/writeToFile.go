package todoServices

import (
	"encoding/json"
	"fmt"
	"os"
	fileStruct "projMod/internal/models"
)

// Функция записи данных в файл
// Принимает слайс из данных типа FileDataStruct.
// Сам сериализовывает полученный слайс
func WriteIntoFile(data []fileStruct.FileDataStruct, fileName string) error {
	// Сериализовываем данные. Они хранятся в типе байт
	dataForFile, err := json.MarshalIndent(data, " ", " ")
	if err != nil {
		return fmt.Errorf("todoWrite ошибка сериализации, %w", err)
	}

	// Запись данных в файл
	err = os.WriteFile(fileName, dataForFile, 0644)
	if err != nil {
		return fmt.Errorf("todoWrite ошибка записи в файл, %w", err)
	}

	return nil
}
