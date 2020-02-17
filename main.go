package main

import (
  "fmt"
  "net/http"
  "os"

  "github.com/urfave/negroni"
)

func main() {
  port := os.Getenv("PORT")
  mux := http.NewServeMux()
  mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "Welcome to the home page!")
  })

  mux.HandleFunc("/callrail", func(w http.ResponseWriter, req *http.Request) {
	 fmt.Fprintln(w,"method=",req.Method)
	 if (req.Method != "POST") {
		return
	 }
    fmt.Fprintf(w, "Welcome to the call rail page!")
  })

  n := negroni.Classic() // Includes some default middlewares
  n.UseHandler(mux)

  http.ListenAndServe(port, n)
}
