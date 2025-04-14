// Функция удаления. Получает на вход номер задачи. По нему находит задачу и
// удаляет всю информацию об этой задаче
package functionality

import (
	"encoding/json"
	"fmt"
	"os"
)

func TodoDelete(num int) {
	deserData := &FileDataStruct{}

	// Чтение файла
	data, err := os.ReadFile("./OurToDoList.json")
	if err != nil {
		fmt.Println("Ошибка чтения", err)
		return
	}

	// Десериализация
	err = json.Unmarshal(data, deserData)
	if err != nil {
		fmt.Println("Десериализация", err)
	}

	// Добавить поиск из прочитанного десериализованного файла. Скорее всего
	// реализовать структуру хэш мапы????

}
