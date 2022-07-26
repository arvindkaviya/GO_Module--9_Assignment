package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server is running")
	http.HandleFunc("/welcome", WelcomeHandler)
	http.HandleFunc("/create-user", PostHandler)
	http.HandleFunc("/get-user", GetHandler)
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Your Welcome to my server!!")
}

type User struct {
	Name string `json:"name"`
	City string `json:"city"`
	Age  int    `json:"age"`
}

func PostHandler(w http.ResponseWriter, r *http.Request) {

	user1 := User{}

	err := json.NewDecoder(r.Body).Decode(&user1)
	if err != nil {
		fmt.Println(err)
	}
	// txtFile, err := os.OpenFile("Users.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModeAppend)

	// if err != nil {
	// 	log.Fatalf("os.Open() failed with '%s'\n", err)
	// }
	// defer func(f *os.File) {
	// 	err := f.Close()
	// 	if err != nil {
	// 	}
	// }(txtFile)

	fmt.Println(user1)
	err = json.NewEncoder(w).Encode(user1)
	if err != nil {
		fmt.Println(err)
	}

	//w.Header().Set("Content-Type", "Application/json")

}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	user1 := User{
		Name: "Arvind Singh Kaviya",
		City: "Jaipur",
		Age:  20,
	}

	err := json.NewEncoder(w).Encode(user1)

	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "Application/json")

}
