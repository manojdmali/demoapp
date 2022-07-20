package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("Hi team lets learn array and splice")

	var newArray = [4]string{}
	newArray[0] = "manu"
	newArray[1] = "manu2"
	newArray[2] = "manu3"
	newArray[3] = "manu4"
	//newArray[4] = "manu4"
	fmt.Println(newArray)

	newSlicearray := []int{}

	newSlicearray = append(newSlicearray, 22, 33, 55, 66, 22, 112, 423, 45, 3463, 67)
	newSlicearray = append(newSlicearray[:5])
	fmt.Println(newSlicearray)
	//sort(newSlicearray)
	//fmt.Println(newSlicearray)

	highScore := make([]int, 4) //this is slice

	highScore[0] = 112
	highScore[1] = 122
	highScore[2] = 142
	highScore[3] = 312
	//highScore[4] = 312 // show error out of the  range

	highScore = append(highScore, 33, 4, 5, 45, 456, 3, 453452, 2342, 4234)

	//fmt.Println(sort.IntsAreSorted(highScore))
	sort.Ints(highScore)
	fmt.Println(highScore)

	//fmt.Println(sort.IntsAreSorted(highScore))

	var courses = []string{"node", "php", "golanguage", "c", "dotnet", "javja", "andriod"}

	//fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

	////////////////////////////MAP //////////////////
	languages := make(map[string]string)

	languages["C"] = "c lang"
	languages["PY"] = "Pytohn"
	languages["RB"] = "Ruby"

	fmt.Println("All values in MAP ", languages)

	fmt.Println("All values in PY ", languages["PY"])

	delete(languages, "RB")
	fmt.Println("All values in MAP ", languages)

	// looops are interstng in map and golang

	for key, value := range languages {
		fmt.Printf("the key: %v and value is : %v  \n ", key, value)
	}

	//// with comma ok syntex here

	for _, value := range languages {
		fmt.Printf("the value is : %v  \n ", value)
	}

	type users struct {
		Name   string
		Age    int
		Email  string
		Status bool
	}

	manu := users{"Manoj", 23, "manoj@mgail.con", true}

	fmt.Println(manu)
	fmt.Printf("Details for manu is %+v \n", manu)

	fmt.Printf("Name is : %v and emai l is %v \n", manu.Name, manu.Email)
	///// stuct in go lang

	// if else in go lang

	if num := 3; num < 12 {
		fmt.Println("Numbe is less then 12")
	} else {
		fmt.Println("Numbe is greater then 12")
	}
	/*
		err := "Err is there"
		if err != nil {
			fmt.Println("NO err ")
		} else {
			fmt.Println("err")
		}
	*/

	///////Switch case in go /////

	mycase := 3

	switch mycase {
	case 1:
		fmt.Print("you are in 1")
	case 3:
		fmt.Print("you are in 3")
	default:
		fmt.Print("you are in default")
	}

	////////for loop in Python/////

	/* 	days := []string{"sunday", "Monday", "Tuesday", "wed", "thu", "Fri", "sat"}

	   	fmt.Println(days) */

	/* for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	} */

	/* for i := range days {
		fmt.Println(days[i])
	}
	*/

	/* for index_days, vals := range days {
		fmt.Printf("the key %v and val is %v \n", index_days, vals)
	} */

	/* for _, vals := range days {
		fmt.Printf("the  val is %v \n", vals)
	} */

	////// functons in golang ////////////////
	shomessage()

	fmt.Printf("normal add result is %v \n", adders(2, 5))

	addrespro, Mymessage := proAdder(2, 4, 2, 6, 34, 856)
	fmt.Printf("Proaddres result is %v with mssage: %v \n", addrespro, Mymessage)

}

func shomessage() {
	fmt.Println("Wlecome to function in Golang")
}

/* func (){
	fmt.Println("Wlecome to anominus function in Golang")
}()
*/

func adders(val1 int, val2 int) int {
	return val2 + val1
}

func proAdder(valuess ...int) (int, string) {
	total := 0
	for _, val := range valuess {
		total += val
	}
	return total, "hi this is mix here as string"

}
