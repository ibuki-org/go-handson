package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	hh := func(w http.ResponseWriter, rq *http.Request) {
		w.Write([]byte("Hello, This is GO-server!"))
	}

	http.HandleFunc("/hello", hh)

	// template を利用する
	tf, er := template.ParseFiles("template/hello.html")
	if er != nil {
		tf, _ = template.New("index").Parse("<html><body><h1>NO TEMPLATE.</h1></body></html>")
	}
	item := struct {
		Title   string
		Message string
	}{
		Title:   "send values",
		Message: "this is a message.",
	}
	th := func(w http.ResponseWriter, rq *http.Request) {
		er = tf.Execute(w, item)
		if er != nil {
			log.Fatal(er)
		}
	}
	http.HandleFunc("/template", th)

	http.ListenAndServe("", nil)
}
