package define

import "github.com/golang-jwt/jwt/v4"

type UserClaim struct {
	Id       int64
	Username string
	jwt.StandardClaims
}

var TokenExpire = 36000

var JwtKey = "mini-tiktok"
var CosBucket = "https://1-1317002166.cos.ap-guangzhou.myqcloud.com"
var TencentSecretID = "AKID1qSr2ygGdniQuDI4dx50pQp22qQozweA"
var TencentSecretKey = "Yv89ze7NSL0Wqyuc7SaVDi2A8tUs31nU"
var TencentFilePrefix = "mini-tiktok/"
