package main

import (
	"time"
	"net/http"
	"fmt"
)

func main() {
	// Ser cookie
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		exp := time.Now().Add(365 * 24 * time.Hour)
		cookie := http.Cookie{Name: "Aibek", Value: "Ludovicose02", Expires: exp}
		http.SetCookie(w,&cookie)
	})


	// get cookie
	http.HandleFunc("/test",func(w http.ResponseWriter, r *http.Request) {
		cookie, _ := r.Cookie("Aibek")
		fmt.Fprintln(w,cookie)
	})

	http.ListenAndServe(":8080", nil)

}
