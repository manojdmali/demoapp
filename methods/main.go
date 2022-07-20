package main

import (
	"fmt"
)

func main() {
	fmt.Println("lets learn method")

	manu := users{"Manoj", 23, "manoj@mgail.con", true}

	fmt.Println(manu)
	manu.IsUserActive()
	manu.NewEmail()
	fmt.Println(manu)
}

type users struct {
	Name   string
	Age    int
	Email  string
	Status bool
}

func (u users) IsUserActive() {
	println(" IS user active :", u.Status)
}

func (u users) NewEmail() {
	u.Email = "dmeo@gmail.com"
	println(" IS user email set to  :", u.Email)
}
