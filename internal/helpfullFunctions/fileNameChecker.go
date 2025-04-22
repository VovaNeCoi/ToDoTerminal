package helpfullFunctions

import "strings"

// Проверка имени файла на валидность и передача корректного варианта в случае
// ошибки. При получении пустого значения "" возвращает "OurToDoList.json"
func CheckFileName(fileNameSlice ...string) (string, error) {
	var fileName string

	// Проверка что переадно хоть что-то
	if len(fileNameSlice) == 0 {
		return "OurToDoList.json", nil
	} else {
		fileName = fileNameSlice[0]
	}

	if len(fileName) == 0 {
		return "OurToDoList.json", nil
	}

	if !strings.HasSuffix(fileName, ".json") {
		return fileName + ".json", nil
	}

	return fileName, nil
}
