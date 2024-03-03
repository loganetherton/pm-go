package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/loganetherton/pm-go/config"
	"net/http"
	"time"
)

type LoginService interface {
	LoginUser(email string, password string) bool
}

type userLogin struct {
	Email    string
	Password string
}

func (info *userLogin) LoginUser(email string, password string) bool {
	return info.Email == email && info.Password == password
}

type JWTService interface {
	GenerateToken(email string, isUser bool) string
	ValidateToken(token string) (*jwt.Token, error)
}

type authCustomClass struct {
	Name   string `json:"name"`
	IsUser bool   `json:"user"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func StaticLoginService() LoginService {
	return &userLogin{
		Email:    "loganetherton@gmail.com",
		Password: "password",
	}
}

func JWTAuthService() JWTService {
	return &jwtService{
		secretKey: config.JwtSecretKey,
		issuer:    "logan",
	}
}

func (service *jwtService) GenerateToken(email string, isUser bool) string {
	claims := &authCustomClass{
		Name:   email,
		IsUser: isUser,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    service.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (service *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token: %v", token.Header["alg"])
		}
		return []byte(service.secretKey), nil
	})
}

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer"
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err := JWTAuthService().ValidateToken(tokenString)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			fmt.Println(claims)
		} else {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
