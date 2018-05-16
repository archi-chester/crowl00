package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//	функции инициализации

//	константы инициализации
const (
	//	значение порта по умолчанию
	LISTENING_PORT = "12346"
	//	файл настроек по умолчанию
	SETTINGS_FILE_NAME = "cloud.conf"
)

//	переменные инициализации

//	Какой порт слушаем
var listeningPort string

//	Объявляем структурку настроек
var mySettings settingsStruct

//	Инициализация данных из командной строки
func initPromptArgs() {
	//	Представляемся
	fmt.Println("Вошли в initPromptArgs")
	//	Проверяем сколько аргументов в командной строке
	if len(os.Args) > 1 {
		//	Аргумента два - первый : порт подкючения
		listeningPort = os.Args[1]
	} else {
		//	Аргумент один - порт подключения берем по умолчанию
		listeningPort = LISTENING_PORT
	}
	fmt.Println(listeningPort)
}

//	программа загрузки настроек
func loadSettings() {

	//	Читаем содержимое файла настроек
	myFile, err := os.Open(SETTINGS_FILE_NAME)
	if err != nil {

		//	Файла нет - создаем
		myFile, err := os.OpenFile("./"+SETTINGS_FILE_NAME, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer myFile.Close()

		//	Объявляем для примера один параметр
		mySettings.ListeningPort = LISTENING_PORT
		//	Маршализируем
		buf, err := json.Marshal(mySettings)
		//	Копируем структурку в файлик
		myFile.Write(buf)

		return
	}
	//	Отложенно закрываем
	defer myFile.Close()

	// Получить размер файла
	stat, err := myFile.Stat()
	if err != nil {
		return
	}

	// Чтение файла
	buf := make([]byte, stat.Size())
	_, err = myFile.Read(buf)
	if err != nil {
		return
	}

	//	Маршалим прочитанное в структуру
	json.Unmarshal(buf, &mySettings)

	//	Выводим сообщение
	fmt.Println("Процесс инициализации завершен")
}
