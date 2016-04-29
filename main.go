package boontling

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

var logger Logger

func init() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/list", listHandler)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Query  string
		Result string
	}

	data.Query = r.URL.Query().Get("query")
	if data.Query != "" {
		boont, err := translate(strings.ToLower(data.Query))
		if err == nil {
			logger.Debug(r, boont)
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

	fmt.Fprint(w, "<h3>Nouns</h3>")
	for english, boont := range nouns {
		fmt.Fprintf(w, "%s -> %s<br>", english, boont)
	}

	fmt.Fprint(w, "<h3>Verbs</h3>")
	for english, boont := range verbs {
		fmt.Fprintf(w, "%s -> %s<br>", english, boont)
	}
}
