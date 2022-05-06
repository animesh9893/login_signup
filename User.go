package login_signup

import (
	"fmt"
        "encoding/hex"
        "golang.org/x/crypto/bcrypt"
        "math/rand"
	"encoding/json"
)
// This structure is used to store user details
type User struct {
	User_id int `json:"user_id"`
	Name string `json:"user_name"`
	Password string `json:"password"`
	Organization string `json:"organization_name"`
	Email string `json:"email"`
	Auth_token string `json:"auth_token"`
	Mobile string `json:"mobile"`
	Note string `json:"note"`
	Pin string `json:"pin"`
	Object_token string `json:"object_token"`
}

func Println(x interface{}){
	if (DEBUGGER) {
		fmt.Println(x);
	}
}
func EncryptPassword(password string) (string,error) {
	bytes,err := bcrypt.GenerateFromPassword([]byte(password),14);
	Println("Encrypting password")
	return string(bytes),err;
}

func CheckEqualPassword(password,hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash),[]byte(password));
	Println("Checking Equal Password")
	return err==nil;
}


func GenerateToken(length int) string {
    b := make([]byte, length)
    if _, err := rand.Read(b); err != nil {
        return ""
    }
	Println("Generating Token")
    return hex.EncodeToString(b)
}

func (db *Database )CreateUser(data string ) error {
	var obj User;
	Println("Create User function called")
	json.Unmarshal([]byte(data),&obj)
	obj.Auth_token = GenerateToken(60);
	obj.Object_token = GenerateToken(60);

	if obj.Name == nil {obj.Name=""}
	if obj.Password == nil {obj.Password=""}
	if obj.Organization == nil {obj.Organization=""}
	if obj.Email == nil {obj.Email=""}
	if obj.Auth_token == nil {obj.Auth_token=""}
	if obj.Mobile == nil {obj.Mobile=""}
	if obj.Note == nil {obj.Note=""}
	if obj.Pin == nil {obj.Pin=""}
	if obj.Object_token == nil {obj.Object_token=""}


	Println("Create user sql function called")
	err := db.CreateUserSQL(obj);
	Println("Function calling done")
	return err
}








