package userservice

import (
	"fmt"
	"game_app-traning/entity"
	"game_app-traning/pkg/phonenumber"
)

type Repositoty interface {
	IsPhoneNumberUniqe(phonenumber string) (bool, error)
	Register(u entity.User) (entity.User, error)
}

type Service struct {
	repo Repositoty
}

type RegisterRequest struct {
	PhoneNumber string
	Name        string
}

type RegisterResponse struct {
	User entity.User
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

	//validate name

	if len(req.Name) < 3 {
		return RegisterResponse{}, fmt.Errorf("name length should be grater tahn 3")
	}

	//create new user

	user := entity.User{
		ID:          0,
		PhoneNumber: req.PhoneNumber,
		Name:        req.Name,
	}

	createdUser, err := s.repo.Register(user)

	if err != nil {
		return RegisterResponse{}, fmt.Errorf("unexpected user : %w", err)
	}

	return RegisterResponse{
		User: createdUser,
	}, nil

}
