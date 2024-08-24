package Models

import (
	"fmt"
)

type CoffeeManagment struct {
	CoffeeObject Coffee
}

var singleInstance *CoffeeManagment

func createManagment() *CoffeeManagment {
	if singleInstance == nil {
		singleInstance = &CoffeeManagment{}
	} else {
		fmt.Println("Its already created")
	}
	return &CoffeeManagment{}
}
