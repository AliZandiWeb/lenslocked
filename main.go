package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset:utf-8")
	fmt.Fprint(w, "<h1>welcome to my website</h1><h2><a href=\"/contact\">Contact</a></h2>")
	fmt.Fprint(w, "<h2><a href=\"/faq\">FAQ</a></h2>")
	fmt.Fprint(w, "<h2><a href=\"/any\">Any</a></h2>")

}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset:utf-8")

	fmt.Fprint(w, "<h1>Contact Page</h1><h2><a href=\"/\">Home</a></h2><p>to get in touch, email me at <a href=\"mailto:ali.zandi.web@gmail.com\">ali.zandi.web@gmail.com</a>.</p>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset:utf-8")

	fmt.Fprint(w, `<h1>FAQ Page</h1>
	<h2><a href="/">Home</a></h2>
	<ul>
		<li><b>Is there a free version?</b>Yes! We offer trial for 30 days on any paid plans.</li>
		<li><b>What are your support Hours?</b>We have support staff answering emails 24.7, though response times may be a bit slower on weekends..</li>
		<li><b>How Do I contact support?</b>Email us - <a href="support@lenslocked.com">support@lenslocked.com</a></p></li>
	</ul>
	`)
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset:utf-8")

	w.WriteHeader(404)
	fmt.Fprint(w, "<h1 style='color:red;position:absolute;top:50%;left: 50%;transform: translate(-50%,50%);'>NOT FOUND 404.</h1>")
}

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
	case "/faq":
		faqHandler(w, r)
	default:
		// http.Error(w, "Page Not Found", http.StatusNotFound)
		errorHandler(w, r)
	}
}
func main() {
	// var router Router
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(errorHandler)
	fmt.Println("Starting the server on http://localhost:4000...")
	http.ListenAndServe(":4000", r)
}
