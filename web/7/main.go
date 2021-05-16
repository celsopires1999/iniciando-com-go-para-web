package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
)

type Post struct {
	Id    int
	Title string
	Body  string
}

var db, err = sql.Open("mysql", "root:root@tcp(db:3306)/go_course")

func main() {

	fmt.Println("App rodando...")

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)

	fmt.Println(http.ListenAndServe("app:8080", r))
}

func ListPosts() []Post {
	rows, err := db.Query("SELECT * FROM posts")
	checkErr(err)

	items := []Post{}

	for rows.Next() {
		post := Post{}
		rows.Scan(&post.Id, &post.Title, &post.Body)
		items = append(items, post)
	}

	return items
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("templates/index.html"))
	if err := t.ExecuteTemplate(w, "index.html", ListPosts()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// stmt, err := db.Prepare("INSERT INTO posts (title, body) VALUES (?, ?)")
// checkErr(err)

// _, err = stmt.Exec("My First Post", "My First Content")
// checkErr(err)
