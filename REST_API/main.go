package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type User struct {
	Id          string
	Name        string
	DateOfBirth string
	Friends     []User
}

var Database []User

func main() {
    http.HandleFunc("/users/", getUser)
	http.HandleFunc("/create/user", createUser)
	http.HandleFunc("/", Home)
	log.Fatal(http.ListenAndServe(":4300", nil))

	
}

func Home(w http.ResponseWriter, r *http.Request) {
	user_proto_type := User{Id: "SomeLongHashValue", Name: "Your Name", DateOfBirth: "DD-MM-YYYY"}
	data, err := json.MarshalIndent(user_proto_type, "", "    ")
	if err != nil {
		log.Println(err)
	}
	log.Println("Home     ",r.URL.Path)
	fmt.Fprintf(w, string(data))

}

func createUser(w http.ResponseWriter, r *http.Request) {
	log.Println("createUser     ",r.URL.Path)
	err := r.ParseForm()
	if err != nil {
		log.Println(w, "ParseForm error  %v\n", err)
	}
	id := r.FormValue("id")
	name := r.FormValue("name")
	date_of_birth := r.FormValue("date_of_birth")

	if len(name) > 0 && len(date_of_birth) > 0 {
		new_user := User{Id: id, Name: name, DateOfBirth: date_of_birth}
		Database = append(Database, new_user)
		log.Println("Data \n", new_user)

	}

}

func getUser(w http.ResponseWriter, r *http.Request) {

	log.Println("getAll_user   ",r.URL.Path)
	found := 0
	param := strings.TrimPrefix(r.URL.Path, "/users/")
	log.Println(param)
    if len(param)>0 {
			        
				for _, value := range Database {
					if value.Id == param {
						data, err := json.MarshalIndent(value, "", "    ")
						if err != nil {
							log.Println(err)
						}
						log.Println(string(data))
						fmt.Fprintf(w, string(data))
						found = 1
						break
					}
				}
				if found==0{
					fmt.Fprintf(w, "user does not exist. ")
				}
    }else{
	    	data, err := json.MarshalIndent(Database, "", "    ")
			if err != nil {
				log.Println(err)
			}
			log.Println(string(data))
			if len(Database) == 0 {
				fmt.Fprintf(w, "No users")
			} else{
				  fmt.Fprintf(w, string(data))
				} 
    }

	

}

