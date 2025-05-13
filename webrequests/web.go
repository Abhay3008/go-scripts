package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	uri := "https://reqres.in/api/users"
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
