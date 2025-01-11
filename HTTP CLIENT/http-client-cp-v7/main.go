package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Quotes struct {
	Tags   []string `json:"tags"`
	Author string   `json:"author"`
	Quote  string   `json:"content"`
}

func ClientGet() ([]Quotes, error) {
	// Membuat transport dengan TLS yang tidak diverifikasi
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("GET", "https://api.quotable.io/quotes/random?limit=3", nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var quotes []Quotes
	err = json.Unmarshal(body, &quotes)
	if err != nil {
		return nil, err
	}

	return quotes, nil
}

type data struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type Postman struct {
	Data data
	Url  string `json:"url"`
}

func ClientPost() (Postman, error) {
	postBody, _ := json.Marshal(map[string]string{
		"name":  "Dion",
		"email": "dionbe2022@gmail.com",
	})
	requestBody := bytes.NewBuffer(postBody)

	client := http.Client{}
	req, err := http.NewRequest("POST", "https://postman-echo.com/post", requestBody)
	if err != nil {
		return Postman{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return Postman{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Postman{}, err
	}

	var postmanResponse Postman
	err = json.Unmarshal(body, &postmanResponse)
	if err != nil {
		return Postman{}, err
	}

	return postmanResponse, nil
}

func main() {
	get, err := ClientGet()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(get)
	}

	post, err := ClientPost()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(post)
	}
}
