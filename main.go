package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Title string
}

type JournalEntry struct {
	Date         string `json:"date"`
	LastModified string `json:"lastModified"`
	Markdown     string `json:"markdown"`
}

func indexHandler(w http.ResponseWriter, r *http.Request, title string) {
	enableCsp(&w)
	p := &Page{
		Title: "What Got Done",
	}
	renderTemplate(w, "index.html", p)
}

var templates = template.Must(
	// Use custom delimiters so Go's delimiters don't clash with Vue's.
	template.New("client/dist/index.html").Delims("[[", "]]").ParseFiles(
		"client/dist/index.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, r.URL.Path)
	}
}

func entriesHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	entries := []JournalEntry{
		JournalEntry{
			Date:         "2019-03-22",
			LastModified: "2019-03-22T08:15:22.382Z",
			Markdown:     "Ate some crackers",
		},
		JournalEntry{
			Date:         "2019-03-15",
			LastModified: "2019-03-15T22:06:45.196Z",
			Markdown:     "Took a nap",
		},
		JournalEntry{
			Date:         "2019-03-08",
			LastModified: "2019-03-22T14:59:16.010Z",
			Markdown:     "Watched the movie *The Royal Tenenbaums*.",
		},
	}

	if err := json.NewEncoder(w).Encode(entries); err != nil {
		panic(err)
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func enableCsp(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Security-Policy", "default-src 'self'")
}

func main() {
	fs := http.FileServer(http.Dir("client/dist"))
	http.Handle("/css/", fs)
	http.Handle("/js/", fs)
	http.HandleFunc("/entries", entriesHandler)
	http.HandleFunc("/", makeHandler(indexHandler))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
