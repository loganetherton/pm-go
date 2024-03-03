package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/loganetherton/pm-go/web/middleware"
	"net/http"
)

type LoginCredentials struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type loginController struct {
	loginService middleware.LoginService
	jwtService   middleware.JWTService
}

type ValError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func getErrorMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "let":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	}
	return "Unknown error"
}

func LoginHandler(loginService middleware.LoginService, jwtService middleware.JWTService) *loginController {
	return &loginController{
		loginService: loginService,
		jwtService:   jwtService,
	}
}

func (controller *loginController) Login(c *gin.Context) string {
	var credentials LoginCredentials
	// @TODO change to must bind
	if err := c.BindJSON(&credentials); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ValError, len(ve))
			for i, fe := range ve {
				out[i] = ValError{
					Field:   fe.Field(),
					Message: getErrorMessage(fe),
				}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
			return ""
		}
		return "Unknown error"
	}
	if controller.loginService.LoginUser(credentials.Email, credentials.Password) {
		return controller.jwtService.GenerateToken(credentials.Email, true)
	}
	return ""
}
