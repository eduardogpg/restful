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

func CreateUser(user structures.User) structures.Users {
  connection.Create(&user)
  users := structures.Users{
      user,
    }
  return users
}

func GetUser(id string) structures.Users {
  users := structures.Users{}
  connection.Where("id = ?", id).Find(&users)
  return users
}

func CloseConnection(){
  connection.Close()
  log.Println("La Conexión con la base de datos se ha exitosamente.")
}

func GetConnectionString() string{
  return username + ":" + password + "@/" + database
}

