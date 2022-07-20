package main

import (
	"encoding/json"
	"fmt"
)

type courses struct {
	Name     string `json:"username"`
	Email    string
	Password string   `json:"-"`
	Age      int      `json:"user-age"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("let paly with json ")
	//jsonEncode()
	jsonDecode()

}

func jsonEncode() {

	localusr := []courses{
		{"MANOJ", "manoj@ggg.cc", "suppp3453", 3, []string{"demo", "test", "hofo"}},
		{"tony", "tony@ggg.cc", "suppp3453", 3, []string{"dsfsafdo"}},
		{"sony", "sonyj@ggg.cc", "", 3, nil},
	}

	finaljson, err := json.MarshalIndent(localusr, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s \n", finaljson)

}

func jsonDecode() {

	jsondataweb := []byte(`
	{
	"username": "Go lang",
	"Email": "dagadga@gmail.com",
	"Age": "3",
	"Password": "test@asd",
	"tags": ["tag2","tage3"]
	}	
	`)

	var locadata courses

	chkValide := json.Valid(jsondataweb)

	if chkValide {

		fmt.Println("Valide JSON")
		json.Unmarshal(jsondataweb, &locadata)
		fmt.Printf("%v \n", locadata)

	} else {
		fmt.Println("IN-Valide JSON")
	}

	var myOnlinedata map[string]interface{}
	json.Unmarshal(jsondataweb, &myOnlinedata)
	fmt.Printf("value is %v", myOnlinedata)

	for k, v := range myOnlinedata {
		fmt.Printf("key : %v and val is : %v and type is %T \n", k, v, v)
	}
}
