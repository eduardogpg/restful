package connect

import(
	"log"
	"github.com/jinzhu/gorm"
   _ "github.com/jinzhu/gorm/dialects/mysql"

  "../structures"
)

var connection *gorm.DB
var engine_sql string = "mysql"

var username string = "root"
var password string = ""
var database string = "rest_go"

func InitializeDataBase(){
	connection = ConnectORM( GetConnectionString() )
  log.Println("Conexión con la base de datos exitosamente.")
}

func ConnectORM(stringConnection string) *gorm.DB {
	connection, err := gorm.Open(engine_sql,stringConnection)
  if err != nil {
      log.Fatal(err)
      return nil
  }
	return connection
}

func CreateUser(user structures.User) structures.User {
  connection.Create(&user)
  return user
}

func GetUser(id string) structures.User {
  user := structures.User{}
  connection.Where("id = ?", id).First(&user)
  return user
}

func UpdateUser(id string, user structures.User) structures.User {
  currentUser := structures.User{}
  connection.Where("id = ?", id).First(&currentUser)
  
  currentUser.Username = user.Username
  currentUser.First_Name = user.First_Name
  currentUser.Last_Name = user.Last_Name
  connection.Save(&currentUser)

  return currentUser
}

func DeleteUser(id string) {
  user := structures.User{}
  connection.Where("id = ?", id).First(&user)
  connection.Delete(&user)
}

func CloseConnection(){
  connection.Close()
  log.Println("La Conexión con la base de datos se ha exitosamente.")
}

func GetConnectionString() string{
  return username + ":" + password + "@/" + database
}

