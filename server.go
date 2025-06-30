package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handler1)
	http.HandleFunc("/{owner}", handler1)
	http.HandleFunc("/submit", handler2)
	http.Handle("/static/", http.StripPrefix("/static/", fileServ))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

var fileServ = http.FileServer(http.Dir("files/"))

func handler1(w http.ResponseWriter, r *http.Request) {
	owner := r.PathValue("owner")
	if owner == "" {
		owner = "Radi"
	}
	fmt.Fprintf(w, `<html>%s<img src='/static/RadiDev_banner.gif'/><h1>Welcome to the website of %s.</h1><br/><hr/><br/><h3>Currently serving from my <b>Raspberry Pi</b> Home Server.</h3><br/><hr/><br/><p>This site is still a work in progress.</p><br/><div style='background-color:#f2f9f2;padding:10px'><hr/><br/>Reach me via email at <a href='mailto:evaristusanarado@gmail.com' target='__blank'>evaristusanarado@gmail.com</a> or via <a href='https://wa.me/2348138686782' target='__blank'>WhatsApp</a><br/><hr/></div><br/>%s<br/><hr/><br/><img style='width:200px;height:200;border-radius:200px' src='https://github.com/radi-dev.png' alt='github profile photo'/><br/><hr/><br/><br/><hr/><br/></html>`, headr, owner, form)
}
func handler2(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
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

		file_line := fmt.Sprintf("%s, %s, %s\n", v("name"), v("email"), v("message"))
		file.WriteString(file_line)
		fmt.Fprintf(w, `%s<div id="form-container">
  <img src="https://media.giphy.com/media/l0MYB8Ory7Hqefo9a/giphy.gif" alt="Funny" width="200" /><br/><br/>
  <p>😂 Thanks %s, for submitting! Your message was well received.✅✅✅</p>
</div>`, form, name)
	}

}

var headr = `<head>
<title>HTMX Form</title>
<script src='https://unpkg.com/htmx.org@1.9.10'></script>
<link href="https://fonts.googleapis.com/css2?family=Montserrat&display=swap" rel="stylesheet">
<link href="/static/styles.css" rel="stylesheet">
	<style>
	body {
	 font-family: 'Montserrat', sans-serif;
	 margin: 0;
	 padding: 0;
	 width: 100vw;
	 max-width: 1000px;
	 margin-left: auto;
	 margin-right: auto;
	text-align: center;
	 padding: 20px;
	 background: #f9f9f9;
	 }

	 #form-container {
		padding: 30px 40px;
		border-radius: 12px;
		width: 100%;
		background-color: #fff;
		box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
		box-sizing: border-box;
		max-width: 500px;
		margin: 0 auto;
		text-align: left;
	 }

	 h2 {
		text-align: center;
		margin-bottom: 20px;
		color: #333;
	 }

	 label {
		display: block;
		margin-bottom: 6px;
		font-weight: 600;
	 }

	 input[type="text"],
	 input[type="email"],
	 textarea {
		width: 100%;
		padding: 10px 14px;
		margin-bottom: 16px;
		border: 1px solid #ccc;
		border-radius: 6px;
		font-size: 14px;
		box-sizing: border-box;
	 }

	 input[type="submit"] {
      background-color: #4CAF50;
      color: white;
      padding: 12px 18px;
      border: none;
      border-radius: 6px;
      cursor: pointer;
      width: 100%;
      font-size: 15px;
      font-weight: bold;
      transition: background-color 0.2s ease;
    }

    input[type="submit"]:hover {
      background-color: #45a049;
    }


	 textarea {
		resize: vertical;
		min-height: 80px;
	 }

	 img {
		display: block;
		margin: 0 auto 20px;
		max-width: 100%;
		border-radius: 8px;
	 }

	 p {
		font-size: 16px;
	 }

	 @media (max-width: 900px) {
	 	body{
		max-width: 100vw;
		width: 100vw;}
		#form-container {
		max-width: 100vw;
			width: 100vw;}
  </style>
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
