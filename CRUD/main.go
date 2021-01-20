package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Book struct {
	ID    int    `json: "id"`
	Title string `json: "title"`
	Year  int    `json: "year"`
}

// Books list

var books []Book = []Book{
	Book{
		ID:    1,
		Title: "A câmara secreta",
		Year:  1997,
	},
	Book{
		ID:    2,
		Title: "dom casmurro",
		Year:  1899,
	},
}

func serverConfig() {
	httpAddr := ":5000"

	//routes
	routes()

	fmt.Println("server is running in the port", httpAddr)
	err := http.ListenAndServe(httpAddr, nil)
	if err != nil {
		log.Fatal("err")
	}

}

func routes() {
	//http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/Books", httpControler)
	http.HandleFunc("/Books/", listBookByIDHandler)
}

func main() {
	serverConfig()
}

// Handlers

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome")
}

func listBooksHandler(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	encoder.Encode(books)
}

func listBookByIDHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(r.URL.Query())

	id := r.URL.Query().Get("id")

	fmt.Println(id)

	NewID, _ := strconv.Atoi(id)

	for _, uniqueBook := range books {
		if uniqueBook.ID == NewID {
			fmt.Println("temos aqui")
			encoder := json.NewEncoder(w)
			encoder.Encode(uniqueBook)
		}
		return
	}
	w.WriteHeader(http.StatusNotFound)

}

func newBookHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("err")
	}
	var NewBook Book
	json.Unmarshal(body, &NewBook)
	fmt.Println(NewBook)
	books = append(books, NewBook)

	encoder := json.NewEncoder(w)
	encoder.Encode(NewBook)

}

func httpControler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		listBooksHandler(w, r)

	} else if r.Method == "POST" {
		newBookHandler(w, r)
	}

}
