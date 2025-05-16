package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type user struct {
	Id   int
	Name string
}

var data = []user{{Id: 1, Name: "abhay"}, {Id: 2, Name: "roan"}}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	data = append(data, user{Id: 3, Name: "rota"})
}

func adduser(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	n := r.URL.Query().Get("id")
	// if name
	id, _ := strconv.Atoi(n)
	for _, d := range data {
		currid := d.Id
		if currid == id {
			fmt.Fprintf(w, "the id %d alreadu exist for user %s !!!", id, d.Name)
			return
		}
	}
	data = append(data, user{Id: id, Name: name})
	fmt.Fprintf(w, "user %s added with id %d!!!", name, id)
	fmt.Println(data)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/add", adduser)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
