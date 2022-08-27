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
				ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(1 * time.Hour)},
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
