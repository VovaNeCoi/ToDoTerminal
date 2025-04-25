package todoServices

import (
	"encoding/json"
	"fmt"
	"os"
	fileChecker "projMod/internal/helpfullFunctions"
	fileStruct "projMod/internal/models"
)

// Функция чтения данных из файла. Можно передавать пустое значение "", тогда
// чтение произведется из дефолтного файла "OurToDoList.json"
// Возвращает сразу срез данных FileDataStruct с данными из файла. Если данных
// нет вернет пустой срез FileDataStruct
func ReadFromFile(fileName ...string) ([]fileStruct.FileDataStruct, error) {
	// Проверка и передача названия файла в тек. функции
	currFileName, err := fileChecker.CheckFileName(fileName[0])
	if err != nil {
		return nil, fmt.Errorf("toDoRead ошибка проверки названия файла, %w", err)
	}

	// Ловим ошибку пустого файла
	size, err := os.Stat(currFileName)
	if err != nil {
		return nil, fmt.Errorf("toDoRead ошибка статов, %w", err)
	}
	if size.Size() == 0 {
		return []fileStruct.FileDataStruct{}, nil
	}

	var fileData []fileStruct.FileDataStruct

	// Считываем байты из файла
	tmpFileData, err := os.ReadFile(currFileName)
	if err != nil {
		return nil, fmt.Errorf("toDoRead ошибка чтения %w", err)
	}

	// Десериализуем данные из байтового формата в FileDataStruct
	err = json.Unmarshal(tmpFileData, &fileData)
	if err != nil {
		return nil, fmt.Errorf("toDoRead ошибка десериализации %w", err)
	}

	return fileData, err
}

// Можно добавить обработку неверного формата названия файла и автоматическое добавление .json
