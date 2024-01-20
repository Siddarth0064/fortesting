package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"hospetal/internal/auth"
	model "hospetal/internal/models"
	"strconv"

	//"hospetal/internal/handlers"
	"hospetal/internal/middleware"
	"hospetal/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

type Handler struct {
	a  *auth.Auth
	ps services.Patient
	ds services.Doctor
}

func NewHandler(a *auth.Auth, ps services.Patient, ds services.Doctor) (*Handler, error) {
	if a == nil {
		return nil, errors.New("Error in NewHandler func")
	}
	return &Handler{a: a, ps: ps, ds: ds}, nil
}
func (h *Handler) PatientSignup(c *gin.Context) {
	ctx := c.Request.Context()

	traceId, _ := ctx.Value(middleware.TraceIdKey).(string)
	// if !ok {
	// 	log.Error().Str("traceId", traceId).Msg("trace id not found in UserSignin handler")
	// 	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
	// 	return
	// }
	var PatientCreate model.UserSignUp
	body := c.Request.Body
	err := json.NewDecoder(body).Decode(&PatientCreate)
	if err != nil {
		log.Error().Err(err).Msg("Error in SignUp Decodeing")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}
	validate := validator.New()
	err = validate.Struct(&PatientCreate)
	if err != nil {
		log.Error().Err(err).Msg("Error in validating struct signup")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Msg": "Invalid Input Please Provid Valid Struct"})
		return
	}
	ps, err := h.ps.PatientSignup(PatientCreate)
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId).Msg("user signup problem")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "user signup failed"})
		return
	}
	c.JSON(http.StatusOK, ps)
}

func (h *Handler) PatientLogin(c *gin.Context) {
	ctx := c.Request.Context()
	traceId, ok := ctx.Value(middleware.TraceIdKey).(string)
	if !ok {
		log.Error().Str("traceId", traceId).Msg("trace id not found in userLogin handler")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		//return
	}
	var PatientLogin model.UserLogin
	body := c.Request.Body
	err := json.NewDecoder(body).Decode(&PatientLogin)
	if err != nil {
		log.Error().Err(err).Msg("error in decoding Patient login struct")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}
	validate := validator.New()
	err = validate.Struct(&PatientLogin)
	if err != nil {
		log.Error().Err(err).Msg("error in validating patient login struct")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "invalid input"})
		return
	}
	regClaims, err := h.ps.Userlogin(PatientLogin)
	if err != nil {
		log.Error().Err(err).Msg("error in patient Loginin  ")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "invalid input"})
		return
	}
	token, err := h.a.GenerateToken(regClaims)
	if err != nil {
		log.Error().Err(err).Msg("error in Gneerating toek ")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return

	}

	c.JSON(http.StatusOK, gin.H{"TOKEN ": token})
}

func (h *Handler) AddingPatientDeatils(c *gin.Context) {
	ctx := c.Request.Context()
	traceId, ok := ctx.Value(middleware.TraceIdKey).(string)
	if !ok {
		log.Error().Str("traceId", traceId).Msg("trace id not found in Adding Patient Deatils handler")
		//c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		//return
	}
	id, err := strconv.ParseUint(c.Param("number"), 10, 32)
	if err != nil {
		log.Error().Stack().Msg("Error in parameter of adding patient details ")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": http.StatusText(http.StatusBadRequest)})
		return
	}
	fmt.Println(id, "======================================")
	var PatienDeatiles model.PatienDeatiles
	body := c.Request.Body
	err = json.NewDecoder(body).Decode(&PatienDeatiles)
	if err != nil {
		log.Error().Err(err).Msg("error in decoding Patient detailes")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}
	us, err := h.ps.CreatePatientDtls(PatienDeatiles, id)
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId).Msg("Error in creating data of patient deatils")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "Error in creating data of patient deatils"})
		return
	}
	c.JSON(http.StatusOK, us)
}
