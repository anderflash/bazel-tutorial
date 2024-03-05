package main

import (
	"log"
	"net/http"

	"github.com/anderflash/bazel-tutorial/projects/go_hello_world"
	"github.com/gorilla/mux"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request")
	w.Write([]byte(go_hello_world.HelloWorld()))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", YourHandler)
	log.Println("Going to listen on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
