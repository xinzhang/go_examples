package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	response, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("Error", err)
	}

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	length := len(body)
	fmt.Println(string(body[:length]))
	fmt.Println("Total length: ", length)
}
