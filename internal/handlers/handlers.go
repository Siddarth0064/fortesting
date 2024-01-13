package handlers

import (
	"hospetal/internal/auth"
	//"hospetal/internal/middleware"
	"hospetal/internal/services"

	"github.com/gin-gonic/gin"
)

func ApiEndPoints(a *auth.Auth, s *services.Service) *gin.Engine {
	r := gin.New()
	h, _ := NewHandler(a, s, s)
	//m, _ := middleware.NewMiddleware(a)
	r.POST("/signUp", h.PatientSignup)
	r.POST("/login", h.PatientLogin)
	return r
}
