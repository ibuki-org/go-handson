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

	html := `<html>
	<body>
	<h1>Hello</h1>
	<p>This is a sample message.</p>
	</body>
	</html>`
	tf, er := template.New("index").Parse(html)
	if er != nil {
		log.Fatal(er)
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
