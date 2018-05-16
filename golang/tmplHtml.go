package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl = `<html>
<head>
    <title>Hello</title>
</head>
<body>
    Hello {{ . }}
</body>
</html>
`

func handler(w http.ResponseWriter, r *http.Request) {
	t := template.New("main")
	t, _ = t.Parse(tmpl)
	log.Println(r.RemoteAddr, r.RequestURI[len("/"):])
	t.Execute(w, r.RequestURI[len("/"):])
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", handler)
	log.Printf("Started")
	server.ListenAndServe()
}
