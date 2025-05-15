package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	uri := "https://reqres.in/api/users"
	// GET request
	getRequest(uri)

	// POST request
	PostRequest(uri)
}
func getRequest(uri string) {
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("x-api-key", "reqres-free-v1")
	client := &http.Client{}
	resp, _ := client.Do(req)
	response, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Println(string(response))
}
func PostRequest(uri string) {
	body := `{
			"name": "morpheus",
			"job": "leader"
		}`
	r := strings.NewReader(body)
	req, err := http.NewRequest("POST", uri, r)
	req.Header.Add("x-api-key", "reqres-free-v1")
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	resp, _ := client.Do(req)
	response, _ := io.ReadAll(resp.Body)
	fmt.Println(string(response))
}
