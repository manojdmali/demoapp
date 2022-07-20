package main

import (
	"fmt"
	"net/http"
	"net/url"
)

const urlchk = "https://lco.dev"
const url2 = "https://lco.dev:3000/learn?coursename=react&paymentid=sdfgsdfgsdgf&user=manu"

func main() {
	fmt.Println("learn web req and view page")

	resp, err := http.Get(urlchk)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Resp is of type %T \n ", resp)

	defer resp.Body.Close()

	/* 	databody, err := ioutil.ReadAll(resp.Body)

	   	if err != nil {
	   		panic(err)
	   	}

	   	content := string(databody)

	   	fmt.Println(content) */

	dataparse, err := url.Parse(url2)

	if err != nil {
		panic(err)
	}

	fmt.Println(dataparse.Scheme, " :<<:", dataparse.Host, " :>>:", dataparse.RequestURI(), " ::", dataparse.RawQuery)

	quertparam := dataparse.Query()

	fmt.Println(quertparam)

	for _, val := range quertparam {
		fmt.Println(val)
	}

	partsofUrl := &url.URL{
		Scheme:  "https",
		Host:    "loc.dev",
		Path:    "/tutcss",
		RawPath: "user=manoj&dept=IT",
	}

	fmt.Printf("the type is %T and value %v \n", partsofUrl, partsofUrl)

	anotherURL := partsofUrl.String()
	fmt.Println(anotherURL)

}
