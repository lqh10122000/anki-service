package service

import (
	"complaint-service/internal/model"
	"complaint-service/internal/repository"

	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthService interface {
	Login(req LoginRequest) (string, error)
	Register(req LoginRequest) error
}

type authService struct {
	repo repository.AuthRepository
	ser  MailService
}

var jwtKey = []byte("your_secret_key")

func NewAuthService(repo repository.AuthRepository, ser MailService) AuthService {
	return &authService{repo: repo, ser: ser}
}

func (s *authService) Register(req LoginRequest) error {
	if req.Username == "" || req.Password == "" {
		log.Println("Username or password is empty")
		return nil
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return err
	}

	auth := &model.Auth{
		Username: req.Username,
		Password: string(hashedPassword),
	}

	err = s.repo.Register(auth)

	s.ser.SendMail(req.Username, "Registration Successful", "You have successfully registered.")

	if err != nil {
		log.Printf("Error registering user: %v", err)
		return err
	}

	return nil
}

func (s *authService) Login(req LoginRequest) (string, error) {
	if req.Username == "" || req.Password == "" {
		log.Println("Username or password is empty")
		return "", model.ErrCustomerNotFound
	}

	auths, err := s.repo.FindByUsername(req.Username)

	if err != nil {
		log.Printf("Error finding user: %v", err)
		return "", model.ErrCustomerNotFound
	}

	if len(auths) == 0 {
		log.Println("User not found")
		return "", model.ErrCustomerNotFound
	}

	hashedPassword := auths[0].Password

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password))

	if err != nil {
		log.Println("Invalid password")
		return "", model.ErrCustomerNotFound
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": req.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		log.Println("Error generating token:", err)
		return "", err
	}

	log.Println("Login successful for user:", req.Username)
	return tokenString, nil

}
