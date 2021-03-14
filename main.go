package main

import (
	"fmt"
)

func main() {
	c := NewClient()
	fmt.Println(c.Apikey)
	fmt.Println(c.Endpoint)

	fmt.Println(c.GetIssueWithCildren(64976))
}
