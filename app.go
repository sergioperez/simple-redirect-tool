package main

import (
    "fmt"
    "net/http"
    "os"
)

var token = ""
var redirectTo = "https://www.google.es"

func main() {
    http.HandleFunc("/", rootHandler)
    token = os.Getenv("TOKEN")
    if token == "" {
	fmt.Println("Invalid token")
	os.Exit(-1)
    }
    http.ListenAndServe(":8080", nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
	http.Redirect(w, r, redirectTo, http.StatusSeeOther)
    case "POST":
	r.ParseForm()
	reqPassword := r.Form.Get("token")
	if token == reqPassword {
	    url := r.Form.Get("url")
	    fmt.Println("url:", url)
	    redirectTo = url
	    fmt.Fprintf(w, "Redirection updated successfully.\n")
	} else {
	    fmt.Fprintf(w, "Invalid Token.\n")
	}
    default:
	fmt.Fprintf(w, "Invalid method.\n")
    }
}
