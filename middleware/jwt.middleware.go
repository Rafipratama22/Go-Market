package middleware

import (
	"fmt"
	"log"
	"time"

	"github.com/Rafipratama22/go_market/entity"
	"github.com/Rafipratama22/go_market/config"
	// "github.com/Rafipratama22/go_market/service"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	// gorm "github.com/jinzhu/gorm"
)

var (
	db *gorm.DB = config.SetupDatabase()
)

func ValidateToken(ctx *gin.Context) {
	tokenString := ctx.Request.Header.Get("Token")
	usered := entity.User{}
	var tokenMap = map[string]string{
		"user_id": "false",
	}
	if tokenString == "" {
		ctx.JSON(401, gin.H{
			"message": "Token is required",
		})
		ctx.Abort()
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", "secret")
		}
		// fmt.Println("token claims", token.Claims.(jwt.MapClaims))
		return []byte("secret"), nil
	})
	if claims, err := token.Claims.(jwt.MapClaims); err && token.Valid {
		fmt.Println("claims", claims)
		for key, val := range claims {
			if s, ok := val.(string); ok {
				fmt.Printf("%q is a string: %q\n", key, s)
				tokenMap[key] = s
				fmt.Println("tokenMap: ", tokenMap)
			}
		}
	}
	db.Model(&usered).Where("name = ?", tokenMap["user_id"]).First(&usered)
	fmt.Println("usered: ", usered)
	ctx.Set("user_id", tokenMap["user_id"])
	ctx.Next()
	// fmt.Println("token:", token)
	if err != nil {
		log.Fatal("error:", err)
	}
}

func CreateToken(user_id string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("secret"))
}
