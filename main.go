package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>To get in touch, please send an email to <a href=\"mailto:vign1310@hotmail.com\">vign1310@hotmail.com")
}

func pageNotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>Page Not Found</h1>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, `<h1>FAQ</h1>
<ul>
	<li>
		<b>Is there a free version?</b>
		Yes! We offer a free trial for 30 days on any paid plans.
	</li>
	<li>
		<b>What are your support hours?</b>
		We have support staff answering emails 24/7, though response
		times may be a bit slower on weekends.
	</li>
	<li>
		<b>How do I contact support?</b>
		Email us - <a href="mailto:support@photodump.com">support@photodump.com</a>
	</li>
</ul>`)
	
}

func main() {
	// http.HandleFunc("/", handlerFunc)
	// http.HandleFunc("/contact", contactHandler)
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	switch r.URL.Path {
	// 	case "/":
	// 		handlerFunc(w, r)
	// 	case "/contact":
	// 		contactHandler(w, r)
	// 	case "/faq":
	// 		faqHandler(w, r)
	// 	default:
	// 		pageNotFound(w, r)
	// }
	// })
	r := chi.NewRouter()
	r.Get("/", handlerFunc)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func (w http.ResponseWriter, r *http.Request)  {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Server is running on port 3000")
	http.ListenAndServe(":3000", nil)
}
