package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	Id    int
	Title string
	Body  string
}

var db, err = sql.Open("mysql", "root:root@tcp(db:3306)/go_course")

func main() {

	// stmt, err := db.Prepare("INSERT INTO posts (title, body) VALUES (?, ?)")
	// checkErr(err)

	// _, err = stmt.Exec("My First Post", "My First Content")
	// checkErr(err)

	rows, err := db.Query("SELECT * FROM posts")
	checkErr(err)

	// items := []Post{}

	for rows.Next() {
		var id int
		var title string
		var body string

		rows.Scan(&id, &title, &body)
		fmt.Println(id, title, body)
	}

	db.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		post := Post{Id: 1, Title: "Unnamed Post", Body: "No content"}

		if title := r.FormValue("title"); title != "" {
			post.Title = title
		}

		t := template.Must(template.ParseFiles("templates/index.html"))
		if err := t.ExecuteTemplate(w, "index.html", post); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println(http.ListenAndServe("app:8080", nil))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
