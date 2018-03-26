package main

import (
	"net/http"
	"fmt"
	"strings"
	"log"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("schema", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for k,v := range r.Form{
		fmt.Println("key:",k)
		fmt.Println("val:",strings.Join(v,""))
	}

	fmt.Fprintf(w, "Hello Aibek")

}

func main() {
	http.HandleFunc("/",SayHello)
	err := http.ListenAndServe(":9090", nil)

	if err != nil{
		log.Fatal("ListenAndServer: ",err)
	}
}
