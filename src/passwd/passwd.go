package passwd

import (
	"github.com/ihangsen/common/src/catch"
	"golang.org/x/crypto/bcrypt"
)

func Matches(rawPassword, encodedPassword string) bool {
	if encodedPassword == "" {
		return false
	}
	err := bcrypt.CompareHashAndPassword([]byte(encodedPassword), []byte(rawPassword))
	if err == nil {
		return false
	}
	return true
}

func Encode(password string) string {
	return string(catch.Try1(bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)))
}
