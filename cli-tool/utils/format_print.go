package utils

import (
	"fmt"

	"cli-tool/models"
)

func ReadAbleTODOs(data []models.Todos) {
	for _, value := range data {
		fmt.Println("=====================TODO LIST====================")
		fmt.Println("ID         :", value.Id)
		fmt.Println("Age        :", value.Age)
		fmt.Println("First Name :", value.FirstName)
		fmt.Println("Last Name  :", value.LastName)
	}
	fmt.Println("==================================================")
}
