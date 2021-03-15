package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Client struct {
	Endpoint string
	Apikey   string
}

func NewClient() *Client {
	conf := getConfig()
	return &Client{conf.Endpoint, conf.Apikey}
}

func createURLWithCildren(e string, id int) string {
	return e + "/issues/" + strconv.Itoa(id) + ".json?include=children"
}

func (c Client) GetIssueWithCildren(id int) []string {
	hc := &http.Client{}
	req, err := http.NewRequest("GET", createURLWithCildren(c.Endpoint, id), nil)
	req.Header.Set("X-Redmine-API-Key", c.Apikey)
	if err != nil {
		fmt.Println(req)
	}
	resp, err := hc.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Println("Could not get API")
	}
	decoder := json.NewDecoder(resp.Body)
	var r issueRequest
	err = decoder.Decode(&r)
	fmt.Println(&r.Issue)

	return []string{"Golang", "Java"} // とりあえずかいている
}
