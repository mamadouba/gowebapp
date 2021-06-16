package utils

import (
	"math/rand"
	"os"
	"strconv"
	"strings"

	uuid "github.com/satori/go.uuid"
)

func GetInt(key string) int64 {
	str := os.Getenv(key)
	val, err := strconv.ParseInt(str, 64, 0)
	if err != nil {
		return 0
	}
	return val
}
func GetFloat(key string) float64 {
	str := os.Getenv(key)
	val, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0
	}
	return val
}

func GetBool(key string) bool {
	str := os.Getenv(key)
	val, err := strconv.ParseBool(str)
	if err != nil {
		return false
	}
	return val
}

func GetUid() string {
	return uuid.NewV4().String()
}

func GenerateToken(size int, charset string) string {
	if charset == "" {
		charset = "abcdefghijklmnopqrstuvwxyz"
		charset += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		charset += "1234567890"
	}
	charsetArr := strings.Split(charset, "")
	str := make([]string, size)
	for i := 0; i < size; i++ {
		str[i] = charsetArr[rand.Intn(len(charset))]
	}
	return strings.Join(str, "")
}
