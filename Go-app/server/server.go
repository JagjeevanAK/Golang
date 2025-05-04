package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main()  {
  fmt.Println("This is golang server")
  r:=mux.NewRouter()
  r.HandleFunc("/",serve).Methods("GET")
  log.Fatal(http.ListenAndServe(":4000",r))
}

func serve(w http.ResponseWriter,r *http.Request) {
  w.Write([]byte("<h1>Hello world from golang</h1>"))
}


