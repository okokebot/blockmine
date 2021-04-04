package importid

import (
	"regexp"
	"strconv"

	"github.com/okokebot/blockmine/pkg/redmine"
)

func PulloutIssueIdFromReleaseNote(s string) []int {
	r := regexp.MustCompile("\\#\\d{5}")
	match := r.FindAllStringSubmatch(s, -1)
	var ids []int
	for _, v := range match {
		stoi, _ := strconv.Atoi(v[0][1:6])
		ids = append(ids, stoi)
	}

	return ids
}

func CreateBlockNote(s string) string {
	ids := PulloutIssueIdFromReleaseNote(s)
	c := redmine.NewClient()
	blocks := ""
	for _, id := range ids {
		p := c.GetIssue(id)
		c.GetChildrenInfo(p)
		blocks += p.CreateReleaseBlock(*c)
	}
	return blocks
}
