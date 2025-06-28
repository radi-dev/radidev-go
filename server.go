package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", handler1)
	http.HandleFunc("/submit", handler2)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func handler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<html>%s<div style='max-width:500px; margin-left:auto;margin-right:auto;margin-top:30px'><div><h1>Welcome to the website of Radi.</h1></div><br/><hr/><br/><h3>Currently serving from my <b>Raspberry Pi</b> Home Server.</h3><br/><hr/><br/><p>This site is still a work in progress.</p><br/><div style='background-color:#f2f9f2;padding:10px'><hr/><br/>Reach me via email at <a href='mailto:evaristusanarado@gmail.com' target='__blank'>evaristusanarado@gmail.com</a> or via <a href='https://wa.me/2348138686782' target='__blank'>WhatsApp</a><br/><hr/></div><br/>%s<br/><hr/><br/><img style='width:200px;height:200;border-radius:200px' src='https://github.com/radi-dev.png' alt='github profile photo'/><br/><hr/><br/><br/><hr/><br/></div> </html>`, headr, form)
}
func handler2(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	if r.Method == "POST" {
		fmt.Fprintf(w, `<div id="form-container"><p>Hi %s.</p>
  <img src="https://media.giphy.com/media/l0MYB8Ory7Hqefo9a/giphy.gif" alt="Funny" width="300" /><br/><br/>
  <p>😂 Thanks for submitting! We needed that laugh.</p>
</div>`, name)
	}

}

var headr = `<head>
<title>HTMX Form</title>
<script src='https://unpkg.com/htmx.org@1.9.10'></script>
<link href="https://fonts.googleapis.com/css2?family=Montserrat&display=swap" rel="stylesheet">
   <style>
   body {
      font-family: 'Montserrat', sans-serif;
    }

    #form-container {
      padding: 30px 40px;
      border-radius: 12px;
      width: 100%;
      max-width: 400px;
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
      text-align: center;
      font-size: 16px;
    }
  </style>
  </head>`
var form = `<div id="form-container">
	<h2>Send Me a Message</h2>
    <form hx-post="/submit" hx-target="#form-container" hx-swap="outerHTML">
      <label for="name">Name:</label><br />
      <input type="text" id="name" name="name" /><br /><br />

      <label for="email">Email:</label><br />
      <input type="email" id="email" name="email" /><br /><br />

      <label for="message">Message:</label><br />
      <textarea id="message" name="message"></textarea><br /><br />

      <input type="submit" value="Submit" />
    </form>
  </div>`
