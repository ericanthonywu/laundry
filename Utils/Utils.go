package Utils

import (
	"crypto/rand"
	"errors"
	"gorm.io/gorm"
	"io"
	"os"
	"strconv"
)

func GetEnvInt(key string) int {
	value, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		panic(err)
	}

	return value
}

func GenerateOtpCode() string {
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	max := GetEnvInt("OTP_LENGTH")
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

func IsDBNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
