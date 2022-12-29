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

	tf, er := template.ParseFiles("template/hello.html")
	if er != nil {
		tf, _ = template.New("index").Parse("<html><body><h1>NO TEMPLATE.</h1></body></html>")
	}
	th := func(w http.ResponseWriter, rq *http.Request) {
		er = tf.Execute(w, nil)
		if er != nil {
			log.Fatal(er)
		}
	}
	http.HandleFunc("/hello", hh)
	// template を利用する
	http.HandleFunc("/template", th)

	http.ListenAndServe("", nil)
}
