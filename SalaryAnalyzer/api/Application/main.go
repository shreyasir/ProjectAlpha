package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	l := log.New(os.Stdout, "salary-api", log.Flags())
	hh := handlers.NewSalary(l)
	http.HandleFunc("/salary/get", hh)
	http.ListenAndServe("localhost:9090", nil)
}
