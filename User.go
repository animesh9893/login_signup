package login_signup

import (
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

func Testing2() string {
	return "testing 2"
}
func EncryptPassword(password string) (string,error) {
	bytes,err := bcrypt.GenerateFromPassword([]byte(password),14);
	return string(bytes),err;
}

func CheckEqualPassword(password,hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash),[]byte(password));
	return err==nil;
}


func GenerateToken(length int) string {
    b := make([]byte, length)
    if _, err := rand.Read(b); err != nil {
        return ""
    }
    return hex.EncodeToString(b)
}

func CreateUser(data string ) {
	var obj User;
	json.Unmarshal([]byte(data),&obj)
	fmt.Println(obj)
}








