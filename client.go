package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
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
	return e + "/issues/" + strconv.Itoa(id) + ".json?include=children,parent"
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
// よって、この関数で 子Issue に対しての詳細情報を取得する
func (c *Client) GetChildrenInfo(parent Issue) {
	for i, child := range parent.Children {
		parent.Children[i] = c.GetIssue(child.Id)
	}
}

// 未動作確認があれば、その親issue、いれば担当者を列挙する
func (c Client) ReleaseBlockIssues(parent Issue) {
	r := regexp.MustCompile(`.*動作確認.*`)
	unfinStatus := []string{"新規", "進行中", "フィードバック"}

	for _, child := range parent.Children {
		if r.MatchString(child.Subject) && contains(unfinStatus, child.Status.Name) {
			fmt.Println(child)
			fmt.Println(child.AssignedTo)
		}
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
