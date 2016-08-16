package structures

type User struct {
	Id   				int			`json:"message_id"`
  Username 		string	`json:"username"`
  First_Name 	string	`json:"first_name"`
  Last_Name 	string	`json:"last_name"`
}

type Users []User

type Response struct {
  Status 			int			`json:"status"`
  Message 		string 	`json:"messages"`
  Users				Users		`json:"users"`
}
