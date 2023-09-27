package v1

import (
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var JWT = "JWT"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateTokenPair(usrId, secret, username, domain, appId string, uc UserUseCase) (string, string, error) {

	//create a token
	token := jwt.New(jwt.SigningMethodHS256)
	//set the claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = username
	claims["sub"] = usrId
	claims["aud"] = domain + "_" + appId
	claims["iss"] = uc.JwtIssuer
	claims["iat"] = time.Now().UTC().Unix()
	claims["typ"] = JWT

	//set expriry for JWT
	claims["exp"] = time.Now().UTC().Add(time.Minute * time.Duration(uc.JwtTokenExpire)).Unix()
	//create singed token
	signedAccessToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", "", err
	}

	//create a refresh token and set clailms
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	refreshClaims := refreshToken.Claims.(jwt.MapClaims)
	refreshClaims["sub"] = usrId
	refreshClaims["iat"] = time.Now().UTC().Unix()
	refreshClaims["iss"] = uc.JwtIssuer
	//set the expiry for the refresh token
	refreshClaims["exp"] = time.Now().UTC().Add(time.Hour * time.Duration(uc.JwtRefreshExpire)).Unix()
	//create signed refresh token
	signedRefreshAccessToken, err := refreshToken.SignedString([]byte(secret))
	if err != nil {
		return "", "", err
	}

	return signedAccessToken, signedRefreshAccessToken, nil
}
