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
	return dns;
}

func CreateDatabaseObject(server,user,password,database string) *Database{
	var obj Database
	obj.SERVER_NAME      = server;
	obj.SERVER_USER_NAME = user;
	obj.SERVER_PASSWORD  = password;
	obj.DATABASE_NAME    = database;
	obj.DNS_URL	     = DNSstring(server,user,password,database);
	return &obj;
}


func ConnectDB(server,url,user,password,database string) *Database{
	obj := CreateDatabaseObject(url,user,password,database);
	if(server == "mysql"){
		obj.DB,obj.ERROR = sql.Open(server,obj.DNS_URL);
	}
	return obj;
}





