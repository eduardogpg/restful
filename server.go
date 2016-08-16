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
    r.Get("/:id", GetUser)
		r.Post("/new", NewUser)
	})
	server.RunOnAddr(":8000")
}


func GetUser(w http.ResponseWriter, r *http.Request, params martini.Params){
	user_id := params["id"]
	users := connect.GetUser(user_id)
	response := structures.Response {200, "", users }

	json.NewEncoder(w).Encode(response)
}


func NewUser(w http.ResponseWriter, r *http.Request) {
	var user structures.User
	
	log.Println("Eduardo Ismael")

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	users := connect.CreateUser(user)
	response := structures.Response {200, "", users }

	json.NewEncoder(w).Encode(response)
}



