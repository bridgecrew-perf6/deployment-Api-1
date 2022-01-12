package custmiddlewares

import (
	"deployment/config"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// //BASIC AUTH
// func DBBasicAuth(email, password string, c echo.Context) (bool, error) {
// 	checkUser := models.User{
// 		Email:    email,
// 		Password: password,
// 	}

// 	if err := config.DB.Where("email = ? AND password = ?", checkUser.Email, checkUser.Password).Find(&checkUser).Error; err != nil || checkUser.ID == 0 {
// 		return false, nil
// 	}
// 	return true, nil
// }

//FOR LOGIN

func CreateToken(id int, secret string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JWT_SECRET))
}

func ExtractToken(e echo.Context) int {
	token := e.Get("user").(*jwt.Token)
	if token != nil && token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		id := claims["id"]
		switch id.(type) {
		case float64:
			return int(id.(float64))
		default:
			return id.(int)
		}
	}
	return -1 //invalid
}
