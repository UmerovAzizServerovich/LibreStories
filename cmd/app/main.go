package main

import (
	"fmt"
	"librestories/repositories"
	"net/http"
)

func main() {
	if err := repositories.InitUsers(); err != nil {
		fmt.Println(err)
	}
	if err := repositories.InitPublications(); err != nil {
		fmt.Println(err)
	}
	if err := repositories.InitUsers(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Server is listening")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
