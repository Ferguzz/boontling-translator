package boontling

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

func init() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/list", listHandler)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var result string

	input := strings.ToLower(r.URL.Query().Get("input"))
	if input != "" {
		boont, exists := englishToBoont[input]
		if exists {
			result = boont
		} else {
			result = "No translation available :("
		}
	}
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, result)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Current dictionary</h1>")
	for english, boont := range englishToBoont {
		fmt.Fprintf(w, "%s -> %s<br>", english, boont)
	}
}
