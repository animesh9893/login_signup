package login_signup

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	SERVER_NAME      string  `json:"server_name"` // dns url
	SERVER_USER_NAME string  `json:"server_user_name"` // server name like root
	SERVER_PASSWORD  string  `json:"server_password"` // password for db
	DATABASE_NAME    string  `json:"database_name"` // which database to use
	DB               *sql.DB `json:"database_object"` // object to acess db
	ERROR		 error   `json:"error"` //error if occured
	DNS_URL		 string  `json:"dns_url"`
}


func DNSstring(url,user,password,database string) string {
	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s",user,password,url,database);
	Println("DNS string function called")
	return dns;
}

func CreateDatabaseObject(server,user,password,database string) *Database{
	var obj Database
	obj.SERVER_NAME      = server;
	obj.SERVER_USER_NAME = user;
	obj.SERVER_PASSWORD  = password;
	obj.DATABASE_NAME    = database;
	obj.DNS_URL	     = DNSstring(server,user,password,database);
	Println("Create Database object fucntion called")
	Println(obj)
	return &obj;
}


func ConnectDB(server,url,user,password,database string) *Database{
	obj := CreateDatabaseObject(url,user,password,database);
	if(server == "mysql"){
		obj.DB,obj.ERROR = sql.Open(server,obj.DNS_URL);
	}
	Println("Connect DB is called")
	Println(obj)
	return obj;
}


func (db *Database) CreateUserTable() error {
	query := "CREATE TABLE IF NOT EXISTS "+TABLE_NAME+`(
		user_id INT NOT NULL AUTO_INCREMENT,
		user_name varchar(100) NOT NULL,
		password varchar(100) NOT NULL,
		organization_name varchar(200),
		email varchar(200) NOT NULL UNIQUE,
		auth_token varchar(60) NOT NULL,
		mobile varchar(13),
		note text,
		pin varchar(16),
		object_token varchar(60),
		PRIMARY KEY (user_id)
	)`
	_,err := db.DB.Query(query)
	Println("Create user table function called\t Error is -- ")
	Println(err)
	return err;
}

func (db *Database) CreateUserSQL(obj User) error {
	/*query := `INSERT INTO `+TABLE_NAME+`(
		user_name,password,organization_name,email,auth_token,mobile,note,pin,object_token)`+` VALUES (`
	string(obj.Name)+", "
	string(obj.Password)+", "
	string(obj.Organization)+", "
	string(obj.Email)+", "
	string(obj.Auth_token)+", "
	string(obj.Mobile)+", "
	string(obj.Note)+", "
	string(obj.Pin)+", "
	string(obj.Object_token)+");"
*/
	query := fmt.Sprintf(`insert into %s ( user_name,password,             organization_name,email,auth_token,mobile,note,pin,object_token) VALUES ("%s","%s","%s","%s","%s","%s","%s","%s","%s")`,obj.Name,obj.Password,obj.Organization,obj.Email,obj.Auth_token,obj.Mobile,obj.Note,obj.Pin,obj.Object_token);

	Println(query)
	Println("Executing Query")

	_,err := db.DB.Query(query)
	return err;
}







