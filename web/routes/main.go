package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/loganetherton/pm-go/controllers"
	"github.com/loganetherton/pm-go/utils"
	"github.com/loganetherton/pm-go/web/middleware"
	"net/http"
)

func SetupRoutes() {
	defer utils.Recover("Unhandled exception", "And another")

	loginController := controllers.LoginHandler(middleware.StaticLoginService(), middleware.JWTAuthService())

	r := setupRouter()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "ping",
		})
	})

	r.POST("/login", func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": "err"})
		return
		token := loginController.Login(c)
		if c.IsAborted() {
			return
		}
		if token != "" {
			c.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			c.JSON(http.StatusUnauthorized, nil)
		}
	})

	authRoutes := r.Group("/")
	authRoutes.Use(middleware.AuthorizeJWT())
	authRoutes.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"profile": "",
		})
	})

	err := r.Run()
	if err != nil {
		panic(err)
	}
}
