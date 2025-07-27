package main

import (
	"fmt"
	"game_app-traning/entity"
	"game_app-traning/repository/mysql"
)

func main() {

}

func TestuserMySqlRepo() {
	mysqlRepo := mysql.New()

	createdUser, err := mysqlRepo.Register(entity.User{
		ID:          0,
		Name:        "reza",
		PhoneNumber: "0937278",
	})

	if err != nil {
		fmt.Println("cant find user")
	} else {
		fmt.Println("created user ", createdUser)
	}

	isUniqe, err := mysqlRepo.IsPhoneNumberUniqe(createdUser.PhoneNumber)

	if err != nil {
		fmt.Println("uniqe err", err)
	}
	fmt.Println("isuniqe", isUniqe)
}
