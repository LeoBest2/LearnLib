package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey []byte = []byte("secret")

const (
	TOKEN_MAX_EXPIRE_HOUR      = 1  // token最长有效期
	TOKEN_MAX_REMAINING_MINUTE = 15 // token还有多久过期就返回新token
)

type customClaims struct {
	Username string `json:"username"`
	IsAdmin  bool   `json:"IsAdmin"`
	jwt.RegisteredClaims
}

//gin jwt 认证中间件
func AuthRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := strings.TrimPrefix(ctx.GetHeader("Authorization"), "Bearer ")
		token, err := jwt.ParseWithClaims(tokenString, &customClaims{}, func(t *jwt.Token) (interface{}, error) { return jwtKey, nil })
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"code": -1, "msg": fmt.Sprintf("access token parse error: %v", err)})
			return
		}
		if claims, ok := token.Claims.(*customClaims); ok && token.Valid {
			if !claims.VerifyExpiresAt(time.Now(), false) {
				ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"code": -1, "msg": "access token expired"})
				return
			}
			// 即将超过过期时间，则添加一个http header `new-token` 给前端更新
			if t := claims.ExpiresAt.Time.Add(-time.Minute * TOKEN_MAX_REMAINING_MINUTE); t.Before(time.Now()) {
				claims := customClaims{
					Username: claims.Username,
					IsAdmin:  claims.Username == "admin",
					RegisteredClaims: jwt.RegisteredClaims{
						ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(TOKEN_MAX_EXPIRE_HOUR * time.Hour)},
					},
				}
				token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
				tokenString, _ := token.SignedString(jwtKey)
				ctx.Header("new-token", tokenString)
			}
			ctx.Set("claims", claims)
		} else {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"code": -1, "msg": fmt.Sprintf("Claims parse error: %v", err)})
			return
		}
		ctx.Next()
	}
}

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	r := gin.Default()
	r.POST("/auth/login", func(ctx *gin.Context) {
		var req loginRequest
		ctx.BindJSON(&req)

		if req.Username != req.Password {
			ctx.JSON(http.StatusOK, gin.H{"code": -1, "msg": "incorrect username or password"})
			return
		}

		log.Printf("login user " + req.Username)

		claims := customClaims{
			Username: req.Username,
			IsAdmin:  req.Username == "admin",
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(TOKEN_MAX_EXPIRE_HOUR * time.Hour)},
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		if tokenString, err := token.SignedString(jwtKey); err != nil {
			ctx.JSON(http.StatusOK, gin.H{"code": -1, "msg": "generate access token failed: " + err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"code": 0, "msg": "", "data": tokenString})
		}
	})

	api := r.Group("/api")
	api.Use(AuthRequired())
	api.GET("/test", func(ctx *gin.Context) {
		claims := ctx.MustGet("claims").(*customClaims)
		ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": fmt.Sprintf("current user: %v , is admin: %v", claims.Username, claims.IsAdmin)})
	})

	r.Run(":8080")
}
