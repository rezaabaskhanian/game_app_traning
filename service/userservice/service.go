package userservice

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"game_app-traning/entity"
	"game_app-traning/pkg/phonenumber"
)

type Repositoty interface {
	IsPhoneNumberUniqe(phonenumber string) (bool, error)
	Register(u entity.User) (entity.User, error)
	GetUserByPhoneNumber(phonenumber string) (entity.User, bool, error)
}

type Service struct {
	repo Repositoty
}

type RegisterRequest struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password`
}

type RegisterResponse struct {
	User entity.User
}

func New(repo Repositoty) Service {
	return Service{repo: repo}

}

func (s Service) Register(req RegisterRequest) (RegisterResponse, error) {
	// TODO - we should verify phone number by verfication code

	// validate phonenumber

	if !phonenumber.IsValid(req.PhoneNumber) {
		return RegisterResponse{}, fmt.Errorf("phone number is not valid")

	}

	// check uniqeness of phone number
	if isuniqe, err := s.repo.IsPhoneNumberUniqe(req.PhoneNumber); err != nil ||
		!isuniqe {
		if err != nil {
			return RegisterResponse{}, err
		}

		if !isuniqe {
			return RegisterResponse{}, fmt.Errorf("phone number is not uniqe")
		}

	}
	// TODO check the password with regex pattern

	// validate password

	if len(req.Password) < 8 {
		return RegisterResponse{}, fmt.Errorf("password length should be grater tahn 8")
	}

	//validate name

	if len(req.Name) < 3 {
		return RegisterResponse{}, fmt.Errorf("name length should be grater tahn 3")
	}

	//create new user

	user := entity.User{
		ID:          0,
		PhoneNumber: req.PhoneNumber,
		Name:        req.Name,
		Password:    GetMD5Hash(req.Password),
	}

	createdUser, err := s.repo.Register(user)

	if err != nil {
		return RegisterResponse{}, fmt.Errorf("unexpected user : %w", err)
	}

	return RegisterResponse{
		User: createdUser,
	}, nil

}

type LoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type LoginResponse struct {
}

func (s Service) Login(req LoginRequest) (LoginResponse, error) {

	//check the existence of phonenumber from reopsitory
	user, exist, err := s.repo.GetUserByPhoneNumber(req.PhoneNumber)

	if err != nil {
		return LoginResponse{}, fmt.Errorf("unexpected user : %w", err)
	}

	if !exist {
		return LoginResponse{}, fmt.Errorf("user or passwordsss isnt correct : %w", err)

	}

	if user.Password != GetMD5Hash(req.Password) {
		return LoginResponse{}, fmt.Errorf("user or password isnt correct : %w", err)
	}

	return LoginResponse{}, nil

}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
