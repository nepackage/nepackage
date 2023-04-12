package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/nepackage/nepackage/auth"
	"github.com/nepackage/nepackage/models"
	"github.com/nepackage/nepackage/repository"
)

// GenerateToken	godoc
// @Summary			GenerateToken
// @Schemes
// @Description		Generate a valid token with user's mail and password
// @Tags			Authentication
// @Accept			json
// @Param			auth	body		models.TokenRequest		true	"User credentials"
// @Produce			json
// @Success			200		{object}	models.TokenResponse
// @Failure			401,500		{object}	models.ErrorMessage
// @Router			/token	[post]
func GenerateToken(c *gin.Context) {
	var request models.TokenRequest
	var user models.User
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	// Check if mail exists and password is correct
	record := models.DB.Where("mail = ?", request.Mail).First(&user)
	if record.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		c.Abort()
		return
	}

	credentialError := user.CheckPassword(request.Password)
	if credentialError != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		c.Abort()
		return
	}

	tokenString, err := auth.GenerateJWT(user.Mail, user.Username, user.ID, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

// RefreshToken	godoc
// @Summary			RefreshToken
// @Schemes
// @Description		Refresh JWT, validate that actual JWT is ok
// @Tags			Authentication
// @Accept			json
// @Param			auth	body		models.TokenRefresh		true	"User credentials"
// @Produce			json
// @Success			200				{object}	models.TokenRefresh
// @Failure			401,500			{object}	models.ErrorMessage
// @Router			/token/refresh	[post]
func RefreshToken(c *gin.Context) {

	var request models.TokenRefresh
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	_, err := auth.ValidateToken(request.Token)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	token, _ := jwt.Parse(request.Token, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return token, nil
	})

	claims := token.Claims.(jwt.MapClaims)
	stringUserID := fmt.Sprint(claims["id"])
	uintUserId, _ := strconv.Atoi(stringUserID)
	user, err := repository.GetUserById(uint(uintUserId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	tokenString, err := auth.GenerateJWT(user.Mail, user.Username, user.ID, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
