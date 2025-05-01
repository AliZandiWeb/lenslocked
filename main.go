package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset:utf-8")
	fmt.Fprint(w, "<h1>welcome to my website</h1><h2><a href=\"/contact\">Contact</a></h2>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Contact Page</h1><h2><a href=\"/\">Home</a></h2><p>to get in touch, email me at <a href=\"mailto:ali.zandi.web@gmail.com\">ali.zandi.web@gmail.com</a>.</p>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.HandleFunc("/contact", contactHandler)
	fmt.Println("Starting the server on http://localhost:4000...")
	http.ListenAndServe(":4000", nil)
}
