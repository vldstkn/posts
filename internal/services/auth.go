package services

import (
	"errors"
	"net/http"
	"p1/pkg/jwt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AuthServiceDeps struct {
	UserService *UserService
}

type AuthService struct {
	UserService *UserService
}

func NewAuthService(deps AuthServiceDeps) *AuthService {
	return &AuthService{
		UserService: deps.UserService,
	}
}

func (service *AuthService) Register(email, password string) (int, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return -1, err
	}
	userId, err := service.UserService.Create(email, string(hashedPassword))
	if err != nil {
		return -1, err
	}
	return userId, nil
}

func (service *AuthService) Login(email, password string) (int, string, error) {
	user := service.UserService.GetByEmail(email)
	if user == nil {
		return -1, "", errors.New("wrong email or password")
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return -1, "", errors.New("wrong email or password")
	}
	return user.Id, user.Role, nil
}

func (service *AuthService) GetNewTokens(refreshToken, secret string, expirationTime time.Time) (string, string, error) {
	isValid, data := jwt.NewJWT(secret).Parse(refreshToken)
	if !isValid {
		return "", "", errors.New("token not valid")
	}
	accessToken, refreshToken, err := service.IssueTokens(*data, secret, expirationTime)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

func (service *AuthService) IssueTokens(data jwt.JWTData, secret string, expirationTime time.Time) (string, string, error) {
	j := jwt.NewJWT(secret)
	accessToken, err := j.Create(data, expirationTime)
	if err != nil {
		return "", "", err
	}
	refreshToken, err := j.Create(data, expirationTime.Add(time.Hour*3))
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

func (service *AuthService) AddCookie(w http.ResponseWriter, name, value string, expirationTime time.Time) {
	cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		Expires:  expirationTime,
		Path:     "/auth",
	}
	http.SetCookie(w, cookie)
}
