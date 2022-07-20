package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Welcome to time  series in Golang")

	presentTime := time.Now()

	fmt.Println(presentTime.Format("02-01-2006 Monday 15:04:05 "))

	creatNewDate := time.Date(2022, time.May, 13, 10, 10, 0, 0, time.Now().Location())
	fmt.Println(creatNewDate.Format("02/01/2006 15:04:05 Monday"))

}
