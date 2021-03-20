package main

import (
	"fmt"

	"github.com/okokebot/blockmine/internal/importid"
	"github.com/okokebot/blockmine/pkg/redmine"
)

func main() {
	ids := importid.PulloutIssueIdFromReleaseNote("[#65527]hogehogeする")
	c := redmine.NewClient()
	p := c.GetIssue(ids[0])
	c.GetChildrenInfo(p)
	s := p.CreateReleaseBlock(*c)
	fmt.Println(s)
}
