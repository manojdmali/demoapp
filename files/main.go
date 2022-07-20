package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	fmt.Print("Lets Learn about files")

	contentforfile := "This date need to add in file"

	file, err := os.Create("./testfile.txt")

	/* if err != nil {
		panic(err)
	} */
	checkErr(err)

	length, err := io.WriteString(file, contentforfile)

	checkErr(err)

	fmt.Println("Length is", length)
	readFile("./testfile.txt")
	defer file.Close()
}

func readFile(filename string) {
	databytes, err := ioutil.ReadFile(filename)
	checkErr(err)

	println("Data in file is : \n", string(databytes))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}

}
