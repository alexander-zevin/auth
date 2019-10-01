package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
	_ "github.com/mattn/go-sqlite3"
	"token"
	"query"
	"structs"
)

var database *sql.DB

func authHandler(w http.ResponseWriter, r *http.Request) { // "w" - поток ответа; "r" - информация о запросе

	// Возвращаем html
	t, _ := template.New("tmpl").ParseFiles("index.html", "templates/main.html", "templates/auth.html") //Свзываем шаблоны
	var dataTMPL = structs.TMPL{Auth: true, Title: "Auth"}
	t.ExecuteTemplate(w, "index", dataTMPL)
}

func indexHandler(w http.ResponseWriter, r *http.Request) { // "w" - поток ответа; "r" - информация о запросе

	// Ищем токен у клиента
	c, err := r.Cookie("token")
	if err != nil {
		authHandler(w, r)
	} else {
		login, password := token.DecodeToken(c.Value)
		if query.CheckValid(login, password, database) != false {

			// Возвращаем html
			t, _ := template.New("tmpl").ParseFiles("index.html", "templates/main.html", "templates/auth.html") //Свзываем шаблоны
			var dataTMPL = structs.TMPL{Auth: false, Title: "Main"}
			t.ExecuteTemplate(w, "index", dataTMPL)
		} else {
			authHandler(w, r)
		}
	}
}

func signIn(w http.ResponseWriter, r *http.Request) {

	time.Sleep(2 * time.Second)

	if r.Method == "POST" {

		// Принимаем данные с клиента
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		var dc structs.DataClient
		obj := r.FormValue("obj")
		err = json.Unmarshal([]byte(obj), &dc)

		//Проверяем введенные данные пользователем
		var signinData = structs.SigninData{}
		if query.CheckValid(dc.Login, dc.Password, database) == true {

			// Cоздаём токен и устанавливаем в куки
			cookie := http.Cookie{
				Name:  "token",
				Value: token.CreateToken(dc.Login, dc.Password),
			}
			http.SetCookie(w, &cookie)

			// Создаём структуру данных
			signinData = structs.SigninData{HTML: delGoSyntax("main"), MsgErr: "", Title: "Main"}
		} else {

			// Создаём структуру данных
			signinData = structs.SigninData{HTML: "", MsgErr: "Wrong login or password"}
		}

		// Возвращаем ответ клиенту
		jsonData, _ := json.Marshal(signinData)
		w.Write([]byte(jsonData))
	}
}

func signUp(w http.ResponseWriter, r *http.Request) {

	// Имитируем задержку сервера
	time.Sleep(2 * time.Second)

	if r.Method == "POST" {
		var reLogin bool = false

		// Принимаем данные с клиента
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		var dc structs.DataClient
		obj := r.FormValue("obj")
		err = json.Unmarshal([]byte(obj), &dc)

		// Создаём токен
		tokenString := token.CreateToken(dc.Login, dc.Password)

		//Проверяем, нет ли порльзователя с одниковым логином
		var users = []structs.Users{}
		users = query.QueryUsers(users, database)
		for _, p := range users {
			if p.Login == dc.Login {
				reLogin = true
			}
		}

		if reLogin == false {

			// Записываем в базу данных
			_, err = database.Exec("insert into users (login, password, token) values (?, ?, ?)", dc.Login, dc.Password, tokenString)
			if err != nil {
				fmt.Println(err)
			}
		}

		// Создаём структуру данных
		var data = structs.Data{HTML: delGoSyntax("main"), Token: tokenString, ReLogin: reLogin}

		// Возвращаем ответ клиенту
		jsonData, _ := json.Marshal(data)
		w.Write([]byte(jsonData))
	}
}

func changeHTML(w http.ResponseWriter, r *http.Request) {

	// Имитируем задержку сервера
	time.Sleep(2 * time.Second)

	if r.Method == "POST" {

		var signinData = structs.SigninData{}

		// Ищем токен у клиента
		c, err := r.Cookie("token")
		if err != nil {
			signinData = structs.SigninData{HTML: delGoSyntax("auth"), MsgErr: "", Title: "Auth"}
		} else {

			// Принимаем данные с клиента
			err1 := r.ParseForm()
			if err1 != nil {
				log.Println(err1)
			}
			obj := r.FormValue("obj")
			var dc structs.DataClient
			err = json.Unmarshal([]byte(obj), &dc)

			login, password := token.DecodeToken(c.Value)
			if query.CheckValid(login, password, database) != false {
				signinData = structs.SigninData{HTML: delGoSyntax(dc.Page), MsgErr: "", Title: strings.Title(dc.Page)}
			}

			if dc.DelCookie == true {
				// Удаляем куки
				cookie := http.Cookie{
					Name:   "token",
					Value:  "deleted",
					MaxAge: -1,
				}
				http.SetCookie(w, &cookie)
			}
		}

		// Возвращаем ответ клиенту
		jsonData, _ := json.Marshal(signinData)
		w.Write([]byte(jsonData))
	}
}

func delGoSyntax(tmpl string) string {
	htmlTMPL, _ := ioutil.ReadFile("templates/" + tmpl + ".html")
	r := strings.NewReplacer(
		"{{define \""+tmpl+"\" }}", "",
		"{{end}}", "")
	html := r.Replace(string(htmlTMPL))
	return html
}

func main() {

	fmt.Println("Слушаю порт 8181")

	db, err := sql.Open("sqlite3", "../db/database.db")
	if err != nil {
		log.Println(err)
	}
	database = db
	defer db.Close()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	// http.Handle("/static/" , http.FileServer(http.Dir("static")))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/auth", authHandler)
	http.HandleFunc("/signup", signUp)
	http.HandleFunc("/signin", signIn)
	http.HandleFunc("/changeHTML", changeHTML)
	http.ListenAndServe(":8181", nil)
}
