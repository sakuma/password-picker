package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"path"
	"strconv"
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
	Id    int `db:"pk"`
	Title string
	Body  string
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
	return (&url.URL{Path: path.Join("/", "passwords", strconv.Itoa(password.Id))}).Path
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

	goji.Get("/", http.FileServer(http.Dir("./public")))
	goji.Get("/assets/*", http.FileServer(http.Dir(".")))
	goji.Get("/passwords", showPasswords)
	goji.Post("/passwords", postPassword)
	goji.Get("/passwords/:id", showPassword)
	goji.Post("/passwords/:id", updatePassword)
	// goji.Delete("/passwords/:id", deletePassword)
	goji.Get("/passwords/:id/delete", deletePassword)

	goji.Serve()
}

func postPassword(c web.C, w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var newPassword Password
	err := decoder.Decode(&newPassword)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	passwordPicker := getPasswordPicker(c)
	db := passwordPicker.DB
	var mutex sync.Mutex
	mutex.Lock()
	defer mutex.Unlock()

	// // TODO: validation
	_, err = db.Insert(&newPassword)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, _ := json.Marshal(newPassword)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
	// http.Redirect(w, r, r.URL.RequestURI(), http.StatusFound)
}

func showPasswords(c web.C, w http.ResponseWriter, r *http.Request) {
	passwordPicker := getPasswordPicker(c)
	db := passwordPicker.DB

	var passwords []Password
	err := db.Select(&passwords, db.OrderBy("Id", genmai.DESC))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, _ := json.Marshal(passwords)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func showPassword(c web.C, w http.ResponseWriter, r *http.Request) {
	passwordPicker := getPasswordPicker(c)
	db := passwordPicker.DB

	var passwords []Password
	err := db.Select(&passwords, db.From(&Password{}), db.Where("id", "=", c.URLParams["id"]))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(passwords) == 0 {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	password := passwords[0]
	password.URL = passwordPicker.PasswordURL(&password)

	tpl, err := pongo2.DefaultSet.FromFile("password.tpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tpl.ExecuteWriter(pongo2.Context{"password": password}, w)
}

func updatePassword(c web.C, w http.ResponseWriter, r *http.Request) {
	passwordPicker := getPasswordPicker(c)
	db := passwordPicker.DB
	var mutex sync.Mutex
	mutex.Lock()
	defer mutex.Unlock()

	var passwords []Password
	err := db.Select(&passwords, db.From(&Password{}), db.Where("id", "=", c.URLParams["id"]))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(passwords) == 0 {
		http.Redirect(w, r, "/passwords", http.StatusFound)
		return
	}
	password := passwords[0]
	password.Title = r.FormValue("title")
	password.Body = r.FormValue("body")
	_, err = db.Update(&password)
	if err == nil {
		log.Println("Update Successful")
		http.Redirect(w, r, passwordPicker.PasswordURL(&password), http.StatusFound)
	} else {
		log.Fatalln(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func deletePassword(c web.C, w http.ResponseWriter, r *http.Request) {
	passwordPicker := getPasswordPicker(c)
	db := passwordPicker.DB
	var mutex sync.Mutex
	mutex.Lock()
	defer mutex.Unlock()

	var passwords []Password
	err := db.Select(&passwords, db.From(&Password{}), db.Where("id", "=", c.URLParams["id"]))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	password := passwords[0]

	_, err = db.Delete(&password)
	if err == nil {
		http.Redirect(w, r, "/passwords", http.StatusFound)
	} else {
		log.Fatal("Failed Delete: ")
		log.Fatal(c.URLParams["id"])
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
