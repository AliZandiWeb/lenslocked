package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset:utf-8")
	fmt.Fprint(w, "<h1>welcome to my website</h1><h2><a href=\"/contact\">Contact</a></h2>")
	fmt.Fprint(w, "<h2><a href=\"/any\">Any</a></h2>")

}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Contact Page</h1><h2><a href=\"/\">Home</a></h2><p>to get in touch, email me at <a href=\"mailto:ali.zandi.web@gmail.com\">ali.zandi.web@gmail.com</a>.</p>")
}

// func errorHandler(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(404)
// 	fmt.Fprint(w, "<h1 style='color:red;position:absolute;top:50%;left: 50%;transform: translate(-50%,50%);'>NOT FOUND 404.</h1>")
// }

// func pathHandler(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	default:
// 		http.Error(w, "Page Not Found", http.StatusNotFound)
// 	}
// }

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		http.Error(w, "Page Not Found", http.StatusNotFound)
	}
}
func main() {
	var router Router
	fmt.Println("Starting the server on http://localhost:4000...")
	http.ListenAndServe(":4000", router)
}
