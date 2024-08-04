package util

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"github.com/RobinHoodArmyHQ/robin-api/pkg/nanoid"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type JWTData struct {
	*jwt.RegisteredClaims
	UserInfo map[string]interface{}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateJwt(userID nanoid.NanoID) (string, error) {
	// set expires at after 3 months
	expiresAt := time.Now().AddDate(0, 3, 0)

	// set user_info claims
	userInfo := map[string]interface{}{
		"user_id":   userID.String(),
		"user_role": "regular",
	}

	claims := &JWTData{
		&jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
		userInfo,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(viper.GetString("auth.jwt_secret")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJwt(signedToken string) (*JWTData, error) {
	token, err := jwt.ParseWithClaims(signedToken, &JWTData{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("auth.jwt_secret")), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTData)
	if !ok {
		return nil, fmt.Errorf("invalid token")
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	expiresAt, err := claims.GetExpirationTime()
	if err != nil {
		return nil, fmt.Errorf("error parsing claims")
	}

	if time.Now().Unix() > expiresAt.Unix() {
		return nil, fmt.Errorf("token expired")
	}

	return claims, nil
}

func GenerateOtp(length int) (string, error) {
	seed := "012345679"
	byteSlice := make([]byte, length)

	for i := 0; i < length; i++ {
		max := big.NewInt(int64(len(seed)))
		num, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}

		byteSlice[i] = seed[num.Int64()]
	}

	return string(byteSlice), nil
}
