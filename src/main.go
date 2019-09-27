package main

import (
	"database/sql" //База данных
	"encoding/json"
	"fmt"           //Ввод/вывод
	"html/template" //Шаблоны
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	// "strconv"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/mattn/go-sqlite3" //Драйвер
)

var database *sql.DB

// Секретный ключ
var mySigningKey = []byte("secret")

type Data struct {
	HTML  string
	Token string
	ReLogin bool
}

type Users struct {
	Login    string
	Password string
	Token    string
}

type TMPL struct {
	Auth  bool
	Title string
}

func authHandler(w http.ResponseWriter, r *http.Request) { // "w" - поток ответа; "r" - информация о запросе

	// Возвращаем html
	t, _ := template.New("tmpl").ParseFiles("index.html", "templates/main.html", "templates/auth.html") //Свзываем шаблоны
	var dataTMPL = TMPL{Auth: true, Title: "Auth"}
	t.ExecuteTemplate(w, "index", dataTMPL)
}

func indexHandler(w http.ResponseWriter, r *http.Request) { // "w" - поток ответа; "r" - информация о запросе

	// Ищем токен у клиента
	c, err := r.Cookie("token")
	if err != nil {
		authHandler(w, r)
	} else {

		// Получаем данные токена
		loginCLient, passwordClient := decodeToken(c.Value)

		// Проверка действительности сессии
		valid := checkValid(loginCLient, passwordClient, c.Value)

		if valid != false {

			// Возвращаем html
			t, _ := template.New("tmpl").ParseFiles("index.html", "templates/main.html", "templates/auth.html") //Свзываем шаблоны
			var dataTMPL = TMPL{Auth: false, Title: "Main"}
			t.ExecuteTemplate(w, "index", dataTMPL)
		} else {
			authHandler(w, r)
		}
	}
}

func queryUsers(users []Users) []Users { //Подключение и выборка БД
	rows, err := database.Query("select login, password, token from users") //получение данных
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	p := Users{}
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

func checkValid(login, password, token string) bool {
	var users = []Users{}
	users = queryUsers(users)
	var valid bool = false
	for _, p := range users {
		if p.Login == login && p.Password == password && p.Token == token {
			valid = true
		}
	}
	return valid
}

func decodeToken(value string) (string, string) {

	// Разбираем токен
	token, _ := jwt.Parse(value, func(token *jwt.Token) (interface{}, error) {
		return []byte(mySigningKey), nil
	})
	claims, _ := token.Claims.(jwt.MapClaims)
	login := claims["login"].(string)
	password := claims["password"].(string)
	return login, password
}

func createToken(login, password string) string {

	// Создаем новый токен
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"login":    login,
		"password": password,
	})
	tokenString, err := token.SignedString([]byte(mySigningKey))
	if err != nil {
		log.Println(err)
	}
	return tokenString
}

func signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var reLogin bool = false

		// Принимаем данные с клиента
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		login := r.FormValue("login")
		password := r.FormValue("password")

		// Копируем HTML шаблон и удаляем синтаксис go
		htmlTMPL, err := ioutil.ReadFile("templates/main.html")
		r := strings.NewReplacer(
			"{{define \"main\" }}", "",
			"{{end}}", "")
		html := r.Replace(string(htmlTMPL))

		// Создаём токен
		tokenString := createToken(login, password)

		//Проверяем, нет ли порльзователя с одниковым логином
		var users = []Users{}
		users = queryUsers(users)
		for _, p := range users {
			if (p.Login == login) {
				reLogin = true
			} 
		}

		if (reLogin == false) {
			// Записываем в базу данных
			_, err = database.Exec("insert into users (login, password, token) values (?, ?, ?)", login, password, tokenString)
			if err != nil {
				fmt.Println(err)
			}

			// Выводим в консоль
			fmt.Println("Регестрация нового пользователя. login: " + login + "; password: " + password)
		}
		
		// Создаём структуру данных
		var data = Data{HTML: html, Token: tokenString, ReLogin: reLogin}

		// Возвращаем ответ клиенту
		jsonData, _ := json.Marshal(data)
		w.Write([]byte(jsonData))
	}
}

func main() {
	fmt.Println("Слушаю порт 8181")

	db, err := sql.Open("sqlite3", "../database/database.db")
	if err != nil {
		log.Println(err)
	}
	database = db
	defer db.Close()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static")))) //Регистрирует обработчик для данного шаблона
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/auth", authHandler)
	http.HandleFunc("/signup", signup)
	http.ListenAndServe(":8181", nil)
}
