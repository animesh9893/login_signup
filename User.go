package login_signup

import (
        "encoding/hex"
        "golang.org/x/crypto/bcrypt"
        "math/rand"
)
// This structure is used to store user details
type User struct {
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


func EncryptPassword(password string) (string,error) {
	bytes,err := bcrypt.GenerateFromPassword([]byte(password),14);
	return string(bytes),err;
}

func CheckEqualPassword(password,hashPassword string) bool {
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
