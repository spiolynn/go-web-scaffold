package security

import (
	"github.com/dgrijalva/jwt-go"
	"go-web-scaffold/pkgs/setting"
	"go-web-scaffold/pkgs/util"
	"time"
)

var (
	jwtSecret        = []byte(setting.G_cfg_yaml.Security.Jwt.Jwtsecret)
	jwtEffectiveTime = time.Duration(setting.G_cfg_yaml.Security.Jwt.Jwteffectivetime)
)

type Claims struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	OtherInfo string `json: otherinfo`
	jwt.StandardClaims
}

// 生成token
// GenerateToken generate tokens used for auth
func GenerateToken(username, password string) (string, error) {

	nowTime := time.Now()

	expireTime := nowTime.Add(jwtEffectiveTime * time.Hour)

	claims := Claims{
		username,
		util.EncodeMD5(password),
		"OtherInfo",
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "go-web-scaffold",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// 解析token
// ParseToken parsing token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
