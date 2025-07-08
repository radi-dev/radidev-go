package handlers

import (
	"fmt"
	"net/http"
	"radidev/config"
	"radidev/repository"

	"github.com/gorilla/mux"
)

func GetUserById(a *config.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := a.DB.Query("SELECT id, name FROM users")
		if err != nil {
			http.Error(w, "DB error", 500)
			return
		}
		defer rows.Close()

		// ...
	}
}

func CreateUserForm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, headr+htmlString+"<p>Welcome to the user creation page!</p></html>")
}

func CreateUser(a *config.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		v := r.FormValue
		user := repository.User{Username: v("username"), PasswordHash: v("password")}
		if user.Username == "" || user.PasswordHash == "" {
			http.Error(w, "Username and password are required", http.StatusBadRequest)
			return
		}
		id, err := user.Create(a.DB)
		if err != nil {
			http.Error(w, "Error creating user: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "%s%s<p>User created with ID: %s</p></html>", headr, htmlString, id)
	}
}

var htmlString = `
               <body> <h1>Create User</h1>
                <form method="POST" action="/admin/users/create">
                    <label for="username">Username:</label>
                    <input type="text" id="username" name="username" required><br>
                    <label for="password">Password:</label>
                    <input type="password" id="password" name="password" required><br>
                    <input type="submit" value="Create User">
                </form><p><a href='/admin/users'>Show all users</a></p></body>
    `

func GetAllUsers(a *config.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := repository.User{}
		users, err := user.List(a.DB)
		if err != nil {
			http.Error(w, "Error getting users: "+err.Error(), http.StatusInternalServerError)
			return
		}
		listItems := ""
		for _, user := range users {
			text := fmt.Sprintf("<li><b>%s</b> %s <a href='/admin/users/%s'><button>Open</button></a></li>", user.Username, user.Id, user.Id)
			listItems += text
		}
		listIBlock := fmt.Sprintf("%s<div><ul>%s</ul><p><a href='/admin/users/create'>Create a new user</a></p></div></html>", headr, listItems)
		fmt.Fprint(w, listIBlock)
	}

}

func GetUser(a *config.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user_id := mux.Vars(r)["id"]
		fmt.Println("Getting user with ID:", user_id)
		user := repository.User{}
		user, err := user.Get(a.DB, user_id)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error getting user: %s", user_id)+err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "%s<p>ID is: %s</p><p>Username is: %s</p><p>PW is: %s</p><p>Created At is: %s</p><p><a href='/admin/users'>Show all users</a></p><p><form method='POST' action='/admin/users/%s/delete'><input type='submit' value='Delete User'></form></p></html>", headr, user.Id, user.Username, user.PasswordHash, user.CreatedAt, user.Id)
	}
}

func DeleteUser(a *config.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user_id := mux.Vars(r)["id"]
		fmt.Println("Deleting user with ID:", user_id)
		user := repository.User{}
		err := user.Delete(a.DB, user_id)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error deleting user: %s", user_id)+err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "%s<p>User with ID %s deleted successfully.</p><p><a href='/admin/users'>Show all users</a></p></html>", headr, user_id)
	}
}

var headr = `<html><head>
<title>Radi Dev</title>
<script src='https://unpkg.com/htmx.org@1.9.10'></script>
<link href="https://fonts.googleapis.com/css2?family=Montserrat&display=swap" rel="stylesheet">
<link href="/static/styles.css" rel="stylesheet">
  </head>`
