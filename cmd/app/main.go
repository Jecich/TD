package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"tst/pkg/slist"

	_ "github.com/go-sql-driver/mysql"
)

type Task struct {
	Id      int    //id в БД
	Content string // содержаине таски
	Idf     int    // id для вывода
}

var tasks []Task

func main() {

	// Реализация связного списка добавлена на будущее как возможность распределения данных в памяти
	l := slist.List{}

	l.Push(1, "ivan", "ewt@mail.com")
	l.Push(2, "iva321n", "ewt@mail.com")
	l.Push(19, "ivawt11n", "ewt@mail.com")
	l.Push(19239, "ivawt11n", "ewt@mail.com")

	elm := l.Head

	for elm != nil {
		fmt.Println(elm)
		elm = elm.Next
	}

	var data []slist.Element
	elm = l.Head

	for elm != nil {
		data = append(data, *elm)
		elm = elm.Next
	}
	// Обработка стилей
	http.Handle("/web/style/", http.StripPrefix("/web/style/", http.FileServer(http.Dir("web/style/"))))

	// Для главной страницы
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		tmpl, _ := template.ParseFiles("web/index.html")

		db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/tst")
		err = db.Ping()

		if err != nil {
			panic(err.Error())
		}
		defer db.Close()

		res, err := db.Query("SELECT * FROM `posts`")

		if err != nil {
			panic(err)
		}

		tasks = []Task{}
		for res.Next() {
			var pst Task
			err = res.Scan(&pst.Id, &pst.Content)
			if err != nil {
				panic(err)
			}
			pst.Idf = len(tasks) + 1
			tasks = append(tasks, pst)
			fmt.Println(fmt.Sprintf("Id: %d Task: %s", pst.Id, pst.Content))
		}

		tmpl.Execute(w, tasks)
	})
	http.HandleFunc("/create", create)
	http.HandleFunc("/save_article", saveArticle)

	// старт HTTP-сервера на порту 8080 протокола TCP с маршрутизатором запросов по умолчанию
	http.ListenAndServe(":8080", nil)
}

// для страницы с созданием заметки
func create(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("web/create.html")
	if err != nil {
		fmt.Println(err)
	}
	t.ExecuteTemplate(w, "create", nil)
}

// страница с сохранением заметки в БД
func saveArticle(w http.ResponseWriter, r *http.Request) {
	cont := r.FormValue("cont")
	if cont == "" {
		fmt.Fprintf(w, "Пустой ввод данных")
	} else {
		db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/tst")
		err = db.Ping()

		if err != nil {
			panic(err.Error())
		}
		defer db.Close()

		insert, err := db.Query(fmt.Sprintf("INSERT INTO `posts` (`content`) VALUES ('%s')", cont))
		if err != nil {
			fmt.Println(err)
		}
		defer insert.Close()

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

// HTTP-обработчик
func mainHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<html><body><h2>ds</h2></body></html>`))
}
