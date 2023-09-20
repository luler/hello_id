package helper

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"strconv"
	"strings"
	"time"
)

// 生成token
func GenerateToken(data map[string]any) string {
	jwt_expire, _ := strconv.Atoi(os.Getenv("JWT_EXPIRE"))
	jwt_secret := os.Getenv("JWT_SECRET")
	exp := time.Now().Add(time.Second * time.Duration(jwt_expire)).Unix()
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["data"] = data
	claims["exp"] = exp

	tokenString, err := token.SignedString([]byte(jwt_secret))
	if err != nil {
		fmt.Println(err)
		CommonException("生成token失败")
	}

	return tokenString
}

// 解析token
func ParseToken(tokenString string, ignore_exp ...bool) map[string]any {
	jwt_secret := os.Getenv("JWT_SECRET")
	// 解析和验证 JWT
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 验证签名密钥，通常是从配置文件中获取
		// 这里简单起见，使用静态字符串作为签名密钥，实际应用中应该更安全地存储密钥
		return []byte(jwt_secret), nil
	})

	ignore_exp_value := false
	if len(ignore_exp) > 0 {
		ignore_exp_value = ignore_exp[0]
	}

	if !ignore_exp_value && (err != nil || !token.Valid) {
		if strings.Contains(err.Error(), "expired") {
			CommonException("token已过期")
		} else {
			CommonException("token无效")
		}
	}

	// 将解析出的用户信息存储到上下文中，供后续处理程序使用
	claims, _ := token.Claims.(jwt.MapClaims)

	return claims
}
