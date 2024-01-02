package define

import "github.com/golang-jwt/jwt/v4"

type UserClaim struct {
	Id       int64
	Username string
	jwt.StandardClaims
}

var TokenExpire = 36000
