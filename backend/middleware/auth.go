package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"music/api"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.GetHeader("token")
		claims, err := validateToken(token)

		if err != nil {
			api.Response(403, "", "token error", c)
			c.Abort()
		}

		c.Set("uid", claims["id"])
	}
}

func validateToken(token string) (jwt.MapClaims, error) {
	// 解析并验证 JWT
	jwtSignKey := viper.GetString("gin.jwt")
	var signingKey = []byte(jwtSignKey)
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// 返回签名密钥
		return signingKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse JWT: %v", err)
	}

	// 验证声明
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid JWT")
	}
}
