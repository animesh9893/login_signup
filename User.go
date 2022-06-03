package login_signup

import (
	"fmt"
        "encoding/ex"
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

func EmailValidate(email string) bool {
    _, err := mail.ParseAddress(email)
    return err == nil
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
	Println("Create user sql function called")
	err := db.CreateUserSQL(obj);
	Println("Function calling done")
	return err
}

func (db *Database)CheckUserPresent(data string) (bool,error) {
	var obj User;
	json.Unmarshal([]byte(data),&obj)
	return db.CheckUserPresentSQL(obj)
}


