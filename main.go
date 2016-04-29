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
	var data struct {
		Query  string
		Result string
	}

	data.Query = strings.ToLower(r.URL.Query().Get("query"))
	if data.Query != "" {
		boont, exists := englishToBoont[data.Query]
		if exists {
			data.Result = boont
		} else {
			data.Result = "No translation available :("
		}
	}
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, data)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Current dictionary</h1>")
	for english, boont := range englishToBoont {
		fmt.Fprintf(w, "%s -> %s<br>", english, boont)
	}
}
