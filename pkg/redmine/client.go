package redmine

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

func createURL(e string, id int) string {
	return e + "/issues/" + strconv.Itoa(id) + ".json?include=children"
}

func (c Client) GetIssue(id int) Issue {
	hc := &http.Client{}
	req, err := http.NewRequest("GET", createURL(c.Endpoint, id), nil)
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
	r := issueRequest{}
	err = decoder.Decode(&r)

	return r.Issue
}

// include=children" では、子Issue のIDやTitle は取得できても
// ステータスは取得できない
// よって、この関数で 子Issue に対しての詳細情報を取得し、上書きする
func (c *Client) GetChildrenInfo(parent Issue) {
	for i, child := range parent.Children {
		parent.Children[i] = c.GetIssue(child.Id)
	}
}

func contains(s []string, e string) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}
