package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

const (
	//	ДБ
	DB_HOST     = "192.168.1.50"
	DB_USER     = "qz0.ru"
	DB_PASSWORD = ""
	DB_NAME     = "qz0.ru"
)

//	подключение к postgres
func connect_to_postgres() {
	//	формируем структурку для подключения
	dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		DB_HOST, DB_USER, DB_PASSWORD, DB_NAME)
	//	проверка ошибок
	db, err := sqlx.Connect("postgres", dbinfo)
	//checkErr(err)
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("подключился к БД")
	}
	defer db.Close()
}

//	подключение к postgres
func loadUsersFromDB(users *[]usersStruct) {
	//	формируем структурку для подключения
	dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		DB_HOST, DB_USER, DB_PASSWORD, DB_NAME)
	//	проверка ошибок
	db, err := sqlx.Connect("postgres", dbinfo)
	//checkErr(err)
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("подключился к БД")
	}
	defer db.Close()

	//	получаем таблицу со списком пользователей
	err = db.Select(users, "SELECT login, password, name, secgroup, properties FROM users")
	// err = db.Select(login, "SELECT `login` FROM `users` LIMIT 1")
	if err != nil {
		fmt.Print("Security Error :", err)
		log.Fatalln(err)
	}
}
