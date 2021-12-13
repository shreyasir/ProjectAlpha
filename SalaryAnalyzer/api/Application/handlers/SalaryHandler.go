package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Salary struct {
	l *log.Logger
}

func NewSalary(l *log.Logger) *Salary {
	return &Salary{l}
}

func (s *Salary) ServeHttp(rw http.ResponseWriter, req *http.Request) {
	s.l.Println("Hello")
	d, err := ioutil.ReadAll(req.Body)
	log.Printf("%s\n", d)
	if err != nil {
		http.Error(rw, "Ooops!", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "Data %s", d)
}
