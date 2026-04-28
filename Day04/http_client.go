package main

import (
	"fmt"
	"io"
	"ioutil"
	"net/http"
)

func main() {
	// Base URL for httpbin
	baseURL := "https://httpbin.org"

	// 1. GET Request
	resp, err := http.Get(baseURL + "/get")
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("GET Response:", string(body))

	// 2. POST Request
	postBody := []byte('{"key":"value"}')
	resp, err := http.Post(baseURL+"/post", "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		fmt.Println("Error making POST request:", err)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("POST Response:", string(body))

	// 3. PUT Request
	client := &http.Client{}
	putBody := []byte('{"key":"updated value"}')
	req, err := http.NewRequest(http.MethodPut, baseURL+"/put", bytes.NewBuffer(postBody))
	req.Header.Set("Content-Type", "application/json")
	req, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making PUT request:", err)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("PUT Response:", string(body))

	// 4. DELETE Request
	req, err := http.NewRequest(http.MethodDelete, baseURL+"/delete", nil)
	req, err := client.Do(req)
	if err != nil {
		fmt.Println("Error on DELETE request:", err)
		return
	}
	defer.req.Body.Close()
	body, _ := io.ReadAll(req.Body)
	fmt.Println("DELETE Response:", string(bodyfunc))
	
}
