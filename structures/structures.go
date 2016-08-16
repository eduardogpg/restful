package structures

type User struct {
	Id   				int			`json:"message_id"`
  Username 		string	`json:"username"`
  First_Name 	string	`json:"first_name"`
  Last_Name 	string	`json:"last_name"`
}

type Response struct {
  Status 			int			`json:"status"`
  User				User 		`json:"users"`
}
