package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"

	"github.com/aniizif/stack-mate/auth-service/internal/services"
)

type AuthHandler struct {
	services *services.AuthService
}

func NewAuthHandler(services *services.AuthService) *AuthHandler {
	return &AuthHandler{services: services}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6,max=32"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	user, token, err := h.services.Register(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data": gin.H{
			"id":           user.ID,
			"email":        user.Email,
			"access_token": token,
			"token_type":   "Bearer",
			"expires_in":   int((15 * time.Minute).Seconds()), // 900
		},
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6,max=32"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	token, err := h.services.Login(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": gin.H{
			"access_token": token,
			"token_type":   "Bearer",
			"expires_in":   int((15 * time.Minute).Seconds()),
		},
	})
}
