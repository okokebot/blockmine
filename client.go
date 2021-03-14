package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
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
	dumpResp, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", dumpResp)

	return []string{"Golang", "Java"} // とりあえずかいている
}
