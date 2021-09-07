package server

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"pkg"
	"pkg/configuration"
)

// ---- CreateToken ----
func CreateToken(userToken root.UserToken, config configuration.Configuration, exp int64, remoteAddr string) string {
	var MySigningKey = []byte(config.Secret)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims[ "user_id" ] = userToken.Userid
	claims[ "username" ] = userToken.Username
	claims[ "email" ] = userToken.Email
	claims["remote_addr"] = remoteAddr
	claims[ "exp" ] = exp
	tokenString, _ := token.SignedString(MySigningKey)
	return tokenString
}

// ---- DecodeJWT ----
func DecodeJWT(curToken string, config configuration.Configuration) root.UserToken {
	var userToken root.UserToken
	var MySigningKey = []byte(config.Secret)
	token, err := jwt.Parse(curToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("error")
		}
		return []byte(MySigningKey), nil
	})
	if err != nil {
		return userToken
	}
	tokenClaims := token.Claims.(jwt.MapClaims)
	userToken.Userid = tokenClaims["user_id"].(string)
	userToken.Username = tokenClaims["username"].(string)
	userToken.Email = tokenClaims["email"].(string)
	userToken.RemoteAddr = tokenClaims["remote_addr"].(string)
	return userToken
}
