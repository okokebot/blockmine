package main

func main() {
	c := NewClient()
	p := c.GetIssue(65527)
	c.GetChildrenInfo(p)
	c.ReleaseBlockIssues(p)
}
