package main

import (
	"github.com/go-martini/martini"
	"log"

	"net/http"
	"encoding/json"

	"./connect"
	"./structures"
)
var port string = ":8000"

type Response struct{
	UserName string `json:"username"`
}

func main(){
	connect.InitializeDataBase()
	defer connect.CloseConnection()
	log.Println("El servidor se encuentra a la escucha en el puerto 8000")
	
	server := martini.Classic()
	server.Group("/users", func(r martini.Router) {
    r.Get("/:id", GetUser)// show something
		r.Post("/new", NewUser) // create something
		r.Patch("/update/:id", UpdateUser) // update something
		r.Delete("/delete/:id", DeleteUser) //delte something
	})
	server.RunOnAddr(":8000")
}


func GetUser(w http.ResponseWriter, r *http.Request, params martini.Params){
	user_id := params["id"]
	user := connect.GetUser(user_id)
	response := structures.Response {200, user }

	json.NewEncoder(w).Encode(response)
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	var user structures.User
	
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		log.Fatal(err)
	}

	response := structures.Response {200, connect.CreateUser(user) }
	json.NewEncoder(w).Encode(response)

}

func UpdateUser(w http.ResponseWriter, r *http.Request, params martini.Params){
	var user structures.User
	id := params["id"]

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		log.Fatal(err)
	}

	response := structures.Response {200, connect.UpdateUser(id, user) }
	json.NewEncoder(w).Encode(response)
}

func DeleteUser(w http.ResponseWriter, r *http.Request, params martini.Params){
	var user structures.User
	
	id := params["id"]
	connect.DeleteUser(id)
	response := structures.Response {200, user}

	json.NewEncoder(w).Encode(response)
}

