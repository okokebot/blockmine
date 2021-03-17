package main

import "fmt"

func main() {
	c := NewClient()
	p := c.GetIssue(65527)
	c.GetChildrenInfo(p)
	s := p.createReleaseBlock(*c)
	fmt.Println(s)
}
