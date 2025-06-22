package main

import (
	"fmt"
	"net/http"
	// "time"
)

func main() {

	http.HandleFunc("/", handler1)
	http.HandleFunc("/hello", handler2)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func handler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><div style='max-width:500px; margin-left:auto;margin-right:auto;margin-top:30px'><div><b>Welcome to the website of Radi.</b></div><div><br/><hr/><br/>Currently serving from my <b>Raspberry Pi</b> Home Server.<br/><hr/><br/><p>The site is still a work in progress.</p><br/><hr/><br/><div>Reach me via email at <a href='mailto:evaristusanarado@gmail.com' target='__blank'>evaristusanarado@gmail.com</a> or via <a href='https://wa.me/2348138686782' target='__blank'>WhatsApp</a></div></div></p><br/><hr/><br/><img style='width:200px;height:200;border-radius:200px' src='https://github.com/radi-dev.png' alt='github profile photo'/><sub><i>How do I center this image?🫣</i></sub></div> </html>")
}
func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
