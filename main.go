package main

import (
	"fmt"
	"log"
	"net/http"
)

// Main Function
func main() {
	fmt.Println("Server Started!")
	http.HandleFunc("/create", CreateList)
	http.HandleFunc("/add", Add)
	http.HandleFunc("/get", GetList)
	http.HandleFunc("/deletetask", DeleteTask)
	http.HandleFunc("/update", UpdateTask)
	http.HandleFunc("/deletelist", DeleteList)
	log.Fatal(http.ListenAndServe(":8000", nil))
}