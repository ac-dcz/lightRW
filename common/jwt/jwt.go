package jwt

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/golang-jwt/jwt/v4/request"
	"net/http"
	"time"
)

type Option struct {
	AccessSecret string
	AccessExpire int64 //seconds
}

func BuildToken(opt *Option, payload map[string]any) (string, error) {
	claims := jwt.MapClaims{}
	now := time.Now().Unix()
	//过期时间
	claims["exp"] = now + opt.AccessExpire
	//签发时间
	claims["iat"] = now
	for k, v := range payload {
		claims[k] = v
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(opt.AccessSecret))
}

func VerifyToken(opt *Option, tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(opt.AccessSecret), nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims.(jwt.MapClaims), nil
}

func ParseTokenFromRequest(opt *Option, r *http.Request) (string, error) {
	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (any, error) {
			return []byte(opt.AccessSecret), nil
		})
	if err != nil {
		return "", err
	}
	return token.Raw, nil
}
