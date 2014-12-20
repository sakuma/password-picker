package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"path"
	"sync"
	"time"

	"github.com/flosch/pongo2"
	_ "github.com/flosch/pongo2-addons"
	_ "github.com/mattn/go-sqlite3"
	"github.com/naoina/genmai"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
)

type PasswordPicker struct {
	URL string
	DB  *genmai.DB
}

type Password struct {
	Id    int64  `db:"pk" column:"id"`
	Title string `column:"title"`
	Body  string `column:"body"`
	URL   string `db:"-"`
	genmai.TimeStamp
}

func (password *Password) BeforeInsert() error {
	n := time.Now()
	password.CreatedAt = n
	password.UpdatedAt = n
	return nil
}

func (password *Password) BeforeUpdate() error {
	n := time.Now()
	password.UpdatedAt = n
	return nil
}

func (passwordPicker *PasswordPicker) PasswordURL(password *Password) string {
	return (&url.URL{Path: path.Join(passwordPicker.URL, "/passwords", string(password.Id))}).Path
}

func getPasswordPicker(c web.C) *PasswordPicker {
	return c.Env["PasswordPicker"].(*PasswordPicker)
}

func main() {
	// setup tables
	db, err := genmai.New(&genmai.SQLite3Dialect{}, "./passwordPicker.db")
	if err != nil {
		log.Fatalln(err)
	}
	if err := db.CreateTableIfNotExists(&Password{}); err != nil {
		log.Fatalln(err)
	}

	// setup pongo
	pongo2.DefaultSet.SetBaseDirectory("views")

	passwordPicker := &PasswordPicker{
		URL: "/",
		DB:  db,
	}
	pongo2.Globals["passwordPicker"] = passwordPicker
	goji.Use(middleware.Recoverer)
	goji.Use(middleware.NoCache)
	goji.Use(func(c *web.C, h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			c.Env["PasswordPicker"] = passwordPicker
			h.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)

	})

	goji.Get("/", Root)
	goji.Get("/passwords", showPasswords)
	goji.Get("/passwords/new", newPassword)
	goji.Post("/passwords", postPassword)

	goji.Serve()
}

func Root(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome Password picker world!")
}

func newPassword(c web.C, w http.ResponseWriter, r *http.Request) {
	var password Password
	tpl, err := pongo2.DefaultSet.FromFile("new.tpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tpl.ExecuteWriter(pongo2.Context{"password": password}, w)
}

func postPassword(c web.C, w http.ResponseWriter, r *http.Request) {

	passwordPicker := getPasswordPicker(c)
	db := passwordPicker.DB
	var mutex sync.Mutex
	mutex.Lock()
	defer mutex.Unlock()

	log.Println("--title")
	log.Println(r.FormValue("title"))
	log.Println("--body")
	log.Println(r.FormValue("body"))

	password := Password{
		Title: r.FormValue("title"),
		Body:  r.FormValue("body"),
	}
	// TODO: validation
	_, err := db.Insert(&password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, r.URL.RequestURI(), http.StatusFound)
}

func showPasswords(c web.C, w http.ResponseWriter, r *http.Request) {
	passwordPicker := getPasswordPicker(c)
	db := passwordPicker.DB

	var passwords []Password
	err := db.Select(&passwords) //, db.OrderBy("Id", genmai.ASC))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tpl, err := pongo2.DefaultSet.FromFile("index.tpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for i := range passwords {
		passwords[i].URL = passwordPicker.PasswordURL(&passwords[i])
	}
	tpl.ExecuteWriter(pongo2.Context{"passwords": passwords}, w)
}
