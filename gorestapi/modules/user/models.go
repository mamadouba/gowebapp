package user

import (
	"fmt"
	"gorestapi/config"
	"gorestapi/db"
	"time"

	"github.com/dgrijalva/jwt-go"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	db.Model
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	Role      int64  `json:"role"`
	Enabled   bool   `json:"enabled"`
}

func (u *User) String() string {
	return fmt.Sprintf("id=%s email=%s role=%d", u.Id, u.Email, u.Role)
}
func (u *User) HashPassword(password string) {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	u.Password = string(bytes)
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) GenerateJWT() (map[string]interface{}, error) {
	secretKey := config.Configuration.SecretKey
	expTime := time.Now().Add(120 * time.Minute).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": u.Email,
		"role":  u.Role,
		"exp":   expTime,
	})
	tokenStr, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}
	result := make(map[string]interface{})
	result["access_token"] = tokenStr
	result["expire_at"] = expTime
	return result, nil
}
