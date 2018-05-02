package main

import (
	"net/http"
	"fmt"
)

func main() {
	adminMux := http.NewServeMux()
	adminMux.HandleFunc("/admin/", adminIndex)
	adminMux.HandleFunc("/admin/panic", panicPage)

	//set middleware
	adminHandler := adminAuthMiddleware(adminMux)

	siteMux := http.NewServeMux()
	siteMux.Handle("/admin/", adminHandler)
	siteMux.HandleFunc("/login", loginPage)
	siteMux.HandleFunc("/logout", logoutPage)
	siteMux.HandleFunc("/", mainPage)

	//set middleware
	siteHandler := accessLogMiddleware(siteMux)
	siteHandler = panicMeddleware(siteHandler)

	fmt.Println("Staarting server at :8080")
	http.ListenAndServe(":8080", siteHandler)
}


// middleware

func panicMeddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("panicMiddleware", r.URL.Path)
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("recovered", err)
				http.Error(w, "Internal server error", 500)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func accessLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("accessLogMiddleware", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func adminAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("adminAuthMiddleware", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

// handler

func adminIndex(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w,"adminINdex")
}

func panicPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"panicPage")
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"loginPage")
}

func logoutPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"logoutPage")
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"mainPage")
}

