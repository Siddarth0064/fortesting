package services

import (
	"errors"
	"fmt"
	model "hospetal/internal/models"
	"hospetal/internal/repository"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	p repository.Patient
	d repository.Doctor
}

func NewService(p repository.Patient, d repository.Doctor) (*Service, error) {
	if p == nil || d == nil {
		return nil, fmt.Errorf("Error in NewService Repository")
	}
	return &Service{p: p, d: d}, nil
}

type Patient interface {
	PatientSignup(ps model.UserSignUp) (model.PetienUser, error)
	Userlogin(ps model.UserLogin) (jwt.RegisteredClaims, error)
	CreatePatientDtls(pc model.PatienDeatiles, id uint64) (model.PatienDeatiles, error)
}

func (s *Service) PatientSignup(ps model.UserSignUp) (model.PetienUser, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(ps.PassHas), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Msg("error occured in hashing password")
		return model.PetienUser{}, errors.New("hashing password failed")
	}

	user := model.PetienUser{UserName: ps.UserName, Email: ps.Email, PhoNo: ps.PhoNo, PassHas: string(hashedPass)}
	// database.CreateTable()
	cu, err := s.p.CreatePatient(user)
	if err != nil {
		log.Error().Err(err).Msg("couldnot create user")
		return model.PetienUser{}, errors.New("user creation failed")
	}
	return cu, nil
}
func (s *Service) Userlogin(ps model.UserLogin) (jwt.RegisteredClaims, error) {
	fu, err := s.p.FetchUserByEmail(ps.Email)
	if err != nil {
		log.Error().Err(err).Msg("couldnot find user")
		return jwt.RegisteredClaims{}, errors.New("user login failed")
	}
	err = bcrypt.CompareHashAndPassword([]byte(fu.PassHas), []byte(ps.PassHas))
	if err != nil {
		log.Error().Err(err).Msg("password of user incorrect")
		return jwt.RegisteredClaims{}, errors.New("user login failed")
	}
	c := jwt.RegisteredClaims{
		Issuer:    "service project",
		Subject:   strconv.FormatUint(uint64(fu.ID), 10),
		Audience:  jwt.ClaimStrings{"users"},
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	return c, nil
}
func (s *Service) CreatePatientDtls(pc model.PatienDeatiles, id uint64) (model.PatienDeatiles, error) {
	fpho, err := s.p.FetchUserByPhono(id)
	fmt.Println(fpho.PhoNo, "===============================")
	if err != nil {
		log.Error().Err(err).Msg("couldnot fetch the parameter of phone number in creatring patient deatiels")
	}
	PatDetls := model.PatienDeatiles{BloodGroup: pc.BloodGroup, Age: pc.Age, Place: pc.Place, Disease: pc.Disease, PetienUser: model.PetienUser{UserName: fpho.UserName, Email: fpho.Email, PhoNo: fpho.PhoNo}}
	cu, err := s.p.CreatePD(PatDetls)
	if err != nil {
		log.Error().Err(err).Msg("couldnot create patient details")
		return model.PatienDeatiles{}, errors.New("patient  creation failed")
	}

	return cu, nil
}
