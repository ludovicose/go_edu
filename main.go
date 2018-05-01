package main

import (
	"net/http"
	"fmt"
	"strings"
	"log"
	"html/template"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("schema", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	fmt.Fprintf(w, "Hello Aibek")

}

func SayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello World")

}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./static/login.gtpl")
		t.Execute(w, nil)
	}else{
		r.ParseForm()
		fmt.Println("username:",r.Form["username"])
		fmt.Println("password:",r.Form["password"])
	}
}

func main() {
	http.HandleFunc("/", SayHello)
	http.HandleFunc("/say", SayHelloName)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
