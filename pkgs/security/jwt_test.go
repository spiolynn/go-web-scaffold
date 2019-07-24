package security

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"go-web-scaffold/pkgs/util"
	"testing"
	"time"
)

// 数据定义
var tGenerateTokenData = []struct { // Test table
	username string
	password string
	err      error
}{
	{"tony", "stark", nil},
}

func TestGenerateToken(t *testing.T) {
	for i, tt := range tGenerateTokenData {
		out, err := GenerateToken(tt.username, tt.password)
		if err != tt.err {
			t.Errorf("%v. %v => %v, wanted: %v", i, tt.username, err, tt.password)
		} else {
			t.Log(out)
		}
	}
}

var tParseTokenData = []struct { // Test table
	token  string
	claims *Claims
	err    error
}{
	{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRvbnkiLCJwYXNzd29yZCI6IjJmYmI3MWIwNGIwMjczODMwMDQyNzg2NmQ2ZTMxODFhIiwiT3RoZXJJbmZvIjoiT3RoZXJJbmZvIiwiZXhwIjoxNTY0MTY1MDQ2LCJpc3MiOiJnby13ZWItc2NhZmZvbGQifQ.YwM5Smy-9aO7cufctnqiHQSgowxzTZ9tM7n2hlag7uc",
		&Claims{util.EncodeMD5("hupan"), util.EncodeMD5("hupan"), "OtherInfo", jwt.StandardClaims{
			ExpiresAt: time.Now().Unix(),
			Issuer:    "gin-blog",
		}}, nil},
	{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRvbnkiLCJwYXNzd29yZCI6IjJmYmI3MWIwNGIwMjczODMwMDQyNzg2NmQ2ZTMxODFhIiwiT3RoZXJJbmZvIjoiT3RoZXJJbmZvIiwiZXhwIjoxNTY0MTY1MDQ2LCJpc3MiOiJnby13ZWItc2NhZmZvbGQifQ.YwM5Smy-9aO7cufctnqiHQSgowxzTZ9tM7n2hlag7uc",
		&Claims{util.EncodeMD5("hupan"), util.EncodeMD5("hupan"), "OtherInfo", jwt.StandardClaims{
			ExpiresAt: time.Now().Unix(),
			Issuer:    "gin-blog",
		}}, nil},
}

func TestParseToken(t *testing.T) {
	for i, tt := range tParseTokenData {
		out, err := ParseToken(tt.token)
		if err != tt.err {
			t.Log(fmt.Sprintf("%v", err))
			t.Errorf("%v. %v => %v, wanted: %v", i, tt.token, err, tt.token)
		} else {
			t.Log(out.Password, out.Username, tt.claims.Username)
		}
	}
}
