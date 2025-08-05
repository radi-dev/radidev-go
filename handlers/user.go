package handlers

import (
	"fmt"
	"net/http"
	"radidev/config"
	"radidev/repository"

	"github.com/gorilla/mux"
)

var user repository.User

func CreateUserForm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, headr+htmlString+"<p>Welcome to the user creation page!</p></html>")
}

func CreateUser(a *config.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		v := r.FormValue
		// user := repository.User{Username: v("username"), PasswordHash: v("password")}
		userMap := map[string]any{
			"username":      v("username"),
			"password_hash": v("password")}

		if userMap["username"] == "" || userMap["password_hash"] == "" {
			http.Error(w, "Username and password are required", http.StatusBadRequest)
			return
		}

		id, err := user.CreateUser(a.DB, userMap)
		// id, err := repository.Create(a.DB, "users", user)
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
		users, err := user.ListUsers(a.DB, "id", "username")

		if err != nil {
			http.Error(w, "Error getting users: "+err.Error(), http.StatusInternalServerError)
			return
		}
		listItems := ""
		for _, user := range users {
			text := fmt.Sprintf("<li><b>%s</b> %s <a href='/admin/users/%s'><button>Open</button></a></li>", user["username"], user["id"], user["id"])
			listItems += text
		}

		fmt.Println("Users:", users)
		listIBlock := fmt.Sprintf("%s<div><ul>%s</ul><p><a href='/admin/users/create'>Create a new user</a></p></div></html>", headr, listItems)
		fmt.Fprint(w, listIBlock)
	}

}

func GetUser(a *config.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user_id := mux.Vars(r)["id"]
		fmt.Println("Getting user with ID:", user_id)

		// user, err := user.Get(a.DB, user_id)
		user, err := user.GetUser(a.DB, user, user_id, "id", "username", "password_hash", "created_at")
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
		err := user.DeleteUser(a.DB, user_id)
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
