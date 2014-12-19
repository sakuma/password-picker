package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"github.com/naoina/genmai"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

// define a table schema.
type Passwords struct {
	Id          int64  `db:"pk" column:"id"`
	Description string `default:"me"`
	body        string
	UserName    string `db:"unique" size:"255"`
	genmai.TimeStamp
}

func main() {
	db, err := genmai.New(&genmai.SQLite3Dialect{}, "./password-picker.db")
	if err != nil {
		log.Fatalln(err)
	}
	if err := db.CreateTableIfNotExists(&Passwords{}); err != nil {
		log.Fatalln(err)
	}

	goji.Get("/", Root)
	goji.Serve()
}

func Root(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome Password picker world!")
}
