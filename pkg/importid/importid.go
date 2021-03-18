package importid

import (
	"regexp"
	"strconv"
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
