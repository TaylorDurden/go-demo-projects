package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	FullName string
	UserName string
	Email    string
}

type Post struct {
	Title  string
	Body   string
	Author User
}

var data []Post = []Post{}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/posts", addPost).Methods("POST")
	router.HandleFunc("/posts/{id}", getPost).Methods("GET")
	router.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
	router.HandleFunc("/posts/{id}", patchPost).Methods("PATCH")
	router.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")
	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", router)
}

func getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	var idParam string = mux.Vars(r)["id"]

	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could  not be  converted  to  integer"))
		return
	}

	if id >= len(data) {
		w.WriteHeader(404)
		w.Write([]byte("No data found with  specified ID"))
		return
	}

	post := data[id]

	json.NewEncoder(w).Encode(post)
}

func addPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	var newPost Post

	json.NewDecoder(r.Body).Decode(&newPost)

	data = append(data, newPost)

	json.NewEncoder(w).Encode(data)
}

func getPosts(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "Application/json")

	fmt.Println("Your details")

	json.NewEncoder(w).Encode(data)

}

func updatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	var idParam string = mux.Vars(r)["id"]

	id, err := strconv.Atoi(idParam)

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not converted to Integer"))
		return
	}

	//error checking

	if id >= len(data) {
		w.WriteHeader(404)
		w.Write([]byte("No data founded with  specified ID"))
		return
	}

	var updatedItem Post

	json.NewDecoder(r.Body).Decode(&updatedItem)

	data[id] = updatedItem

	json.NewEncoder(w).Encode(updatedItem)

}

func patchPost(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "Application/json")

	var idParam string = mux.Vars(r)["id"]

	id, err := strconv.Atoi(idParam)

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not converted to Integer"))
		return
	}

	//error checking

	if id >= len(data) {
		w.WriteHeader(404)
		w.Write([]byte("No data founded with  specified ID"))
		return
	}

	// get the current value
	existingPost := data[id]

	// create a map to store the fields to be updated
	var updates map[string]interface{}
	json.NewDecoder(r.Body).Decode(&updates)

	// check and update fields
	if body, ok := updates["Body"].(string); ok {
		existingPost.Body = body
	}
	if title, ok := updates["Title"].(string); ok {
		existingPost.Title = title
	}
	if author, ok := updates["Author"].(map[string]interface{}); ok {
		if fullName, ok := author["FullName"].(string); ok {
			existingPost.Author.FullName = fullName
		}
		if userName, ok := author["UserName"].(string); ok {
			existingPost.Author.UserName = userName
		}
		if email, ok := author["Email"].(string); ok {
			existingPost.Author.Email = email
		}
	}

	data[id] = existingPost

	json.NewEncoder(w).Encode(existingPost)
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	var idParam string = mux.Vars(r)["id"]

	id, err := strconv.Atoi(idParam)

	if err != nil {

		w.WriteHeader(400)
		w.Write([]byte("ID could not converted to Integer"))
		return
	}

	//error checking

	if id >= len(data) {

		w.WriteHeader(404)
		w.Write([]byte("No data founded with  specified ID"))
		return

	}

	data = append(data[:id], data[id+1:]...)

	w.WriteHeader(200)
}
