package service

import (
	"fmt"
	"self_gqlgen/models"
)

func RegisterPlayer(input models.RegisterUser) (*models.Player, error) {

	return nil, nil

}

func PrintPlayerName(input models.RegisterUser) error {
	fmt.Printf("名字: %s 年龄: %d", input.Name, input.Age)
	return nil
}
