package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"radidev/database"

	"github.com/gorilla/mux"

	_ "github.com/lib/pq"

	"radidev/config"
	"radidev/handlers"
)

var cfg = config.Load()

func main() {
	db, err := database.ConnectDb()
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	defer db.Close()

	err = database.CreateTables(db)
	if err != nil {
		log.Fatal("Error creating tables:", err)
	}

	r := mux.NewRouter()
	app := &config.App{DB: db}

	r.HandleFunc("/admin/users", handlers.GetAllUsers(app)).Methods("GET")
	r.HandleFunc("/admin/users/create", handlers.CreateUserForm).Methods("GET")
	r.HandleFunc("/admin/users/create", handlers.CreateUser(app)).Methods("POST")
	r.HandleFunc("/admin/users/{id}/delete", handlers.DeleteUser(app)).Methods("POST")
	r.HandleFunc("/admin/users/{id}", handlers.GetUser(app)).Methods("GET")

	r.HandleFunc("/", handler1).Methods("GET")
	r.HandleFunc("/{owner}", handler1).Methods("GET")
	r.HandleFunc("/submit", handler2).Methods("POST")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServ)).Methods("GET")

	if err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), r); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

var fileServ = http.FileServer(http.Dir("templates/staticFiles/"))

func handler1(w http.ResponseWriter, r *http.Request) {
	// owner := r.PathValue("owner")
	owner := mux.Vars(r)["owner"] // Get the owner from the URL path
	fmt.Println("Owner:", owner)
	if owner == "" {
		owner = "Radi"
	}
	tmpl := handlers.Homepage()
	tmpl.Execute(w, nil)
	// fmt.Fprintf(w, `<html>%s<body><img src='/static/RadiDev_banner.gif'/><h1>Welcome to the website of %s.</h1><br/><hr/><br/><h3>Currently serving from my <b>Raspberry Pi</b> Home Server.</h3><br/><hr/><br/><p>This site is still a work in progress.</p><br/><div id=contact><hr/><br/>Reach me via email at <a href='mailto:evaristusanarado@gmail.com' target='__blank'>evaristusanarado@gmail.com</a> or via <a href='https://wa.me/2348138686782' target='__blank'>WhatsApp</a><br/><hr/></div><br/>%s<br/><hr/><br/><img id='pic' src='https://github.com/radi-dev.png' alt='github profile photo'/><br/><hr/><br/><br/><hr/><br/></body></html>`, headr, owner, form)
}
func handler2(w http.ResponseWriter, r *http.Request) {
	filename := "formdata.csv"
	v := r.FormValue
	name := v("name")
	file, er := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	if er == os.ErrNotExist {
		_file, _ := os.Create(filename)
		file = _file
		file.WriteString("name,email,message\n") // Header for CSV
	}

	defer file.Close()

	file_line := fmt.Sprintf("%s, %s, '%s\n", v("name"), v("email"), v("message"))
	file.WriteString(file_line)
	fmt.Fprintf(w, `%s<div id="form-container">
  <img src="https://media.giphy.com/media/l0MYB8Ory7Hqefo9a/giphy.gif" alt="Funny" width="200" /><br/><br/>
  <p>😂 Thanks %s, for submitting! Your message was well received.✅✅✅</p>
</div>`, form, name)

}

var headr = `<head>
<title>Radi Dev</title>
<script src='https://unpkg.com/htmx.org@1.9.10'></script>
<link href="https://fonts.googleapis.com/css2?family=Montserrat&display=swap" rel="stylesheet">
<link href="/static/styles.css" rel="stylesheet">
  </head>`
var form = `<div id="form-container">
	<h2>Send Me a Message</h2>
    <form hx-post="/submit" hx-target="#form-container" hx-swap="outerHTML">
      <label for="name">Name:</label><br />
      <input type="text" id="name" name="name" /><br /><br />

      <label for="email">Email:</label><br />
      <input required type="email" id="email" name="email" /><br /><br />

      <label for="message">Message:</label><br />
      <textarea required id="message" name="message"></textarea><br /><br />

      <input type="submit" value="Submit" />
    </form>
  </div>`
