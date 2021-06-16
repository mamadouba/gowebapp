package auth

import (
	"gorestapi/config"
	"gorestapi/utils"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

type JwtToken struct {
	Access  string `json:"access_token"`
	Refresh string `json:"refresh_token"`
	Expire  int64  `json:"expire_at"`
}
type Token struct {
	UserId       string
	RefreshToken string
	ExpireAt     int64
}

func (t *Token) Valid() bool {
	return t.ExpireAt > time.Now().Unix()
}

func (t *Token) SetRefreshToken() {
	charset := strings.Replace(t.UserId, "-", "", -1)
	t.RefreshToken = utils.GenerateToken(32, charset)
}

func (t *Token) GenerateJWT(db *gorm.DB, userId string) (*JwtToken, error) {
	secretKey := config.Configuration.SecretKey
	expTime := time.Now().Add(24 * 60 * time.Minute).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userId,
		"exp": expTime,
	})
	accessToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}
	db.Find(&t, "user_id = ?", userId)
	db.Delete(&t)
	t.UserId = userId
	t.ExpireAt = expTime
	t.SetRefreshToken()
	db.Create(&t)
	return &JwtToken{
		Access:  accessToken,
		Refresh: t.RefreshToken,
		Expire:  expTime,
	}, nil
}

type PwdResetToken struct {
	Token    string
	Email    string
	ExpireAt int64
}

func (t *PwdResetToken) Generate(db *gorm.DB, email string) string {
	expTime := time.Now().Add(15 * time.Minute).Unix()
	token := utils.GenerateToken(24, "")
	pwdToken := PwdResetToken{Token: token, Email: email, ExpireAt: expTime}
	db.Create(&pwdToken)
	return token
}

func (t *PwdResetToken) Valid(db *gorm.DB, token string) (string, bool) {
	var pwdToken PwdResetToken
	if db.Where("token = ?", token).First(&pwdToken); pwdToken.Email == "" {
		return "", false
	}
	if pwdToken.ExpireAt < time.Now().Unix() {
		return "", false
	}
	return "", true
}

type PwdReset struct {
	Password string `validate:"string,min=3"`
}
