package query

import (
	"log"
	"fmt"
	"database/sql"
	"structs"
)

func QueryUsers(users []structs.Users, database *sql.DB) []structs.Users { //Подключение и выборка БД
	rows, err := database.Query("select login, password, token from users") //получение данных
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	p := structs.Users{}
	for rows.Next() { //последовательно перебираем все строки
		err := rows.Scan(&p.Login, &p.Password, &p.Token) //считываем все полученные данные в переменные
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, p) //считываем данные в структуру tasks и затем добавляем ее в срез
	}
	return users
}
