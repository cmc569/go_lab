package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/cmc569/web/mysql"
)


type profile struct {
	Name string
	Data []string
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/db", mysql.Db)
	http.HandleFunc("/json", jsons)

	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("Path", r.URL.Path)
	fmt.Println("Scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_logn"])
	for k, v := range r.Form {
		fmt.Println("Key:", k)
		fmt.Println("Val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello")
}


func jsons(w http.ResponseWriter, r *http.Request) {
	// el := "<h2>Show List</h2>"
	// el := `{"Show": "List"}`
	el := profile{"jason", []string{"program", "tv", "movie"}}
	// profile := Profile{"Jason", []string{"programming", "moving"}}
	// js, _ := json.Marshal(profile)
	js, _ := json.Marshal(el)

	w.Header().Add("content-type", "application/json")
	// fmt.Fprintf(w, el)
	// w.Write(js)
	// js = []byte(js)
	// js := []byte(el)
	// w.Write(js)
	w.Write(js)
}
