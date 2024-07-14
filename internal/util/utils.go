package util

import (
	"fmt"
	"strings"
	"time"

	"github.com/RobinHoodArmyHQ/robin-api/internal/repositories/user"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateJwt(emailId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email_id": emailId,
		"exp":      time.Now().Add(24 * time.Hour),
	})

	tokenString, err := token.SignedString([]byte(viper.GetString("auth.jwt_secret")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJwt(signedToken string) error {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(signedToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("auth.jwt_secret")), nil
	})

	if err != nil {
		return nil
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	expiresAt, err := claims.GetExpirationTime()
	if err != nil {
		return fmt.Errorf("error parsing claims")
	}

	if expiresAt.Unix() < time.Now().Local().Unix() {
		return fmt.Errorf("token expired")
	}

	return nil
}

func GetFirstNameAndLastName(fullName string) (string, string) {
	firstName, lastName := "", ""

	if len(fullName) == 0 {
		return firstName, lastName
	}

	nameArr := strings.Split(fullName, " ")
	firstName = nameArr[0]

	if len(nameArr) > 1 {
		lastName = strings.Join(nameArr[1:], " ")
	}

	return firstName, lastName
}

func SendEmailVerificationCode(userData *user.CreateUserResponse) {
	// TODO: send verification link via AWS-SES.

}
