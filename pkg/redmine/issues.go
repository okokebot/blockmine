package redmine

import (
	"regexp"
	"strconv"
)

type IdName struct {
	Id   int
	Name string
}

type Issue struct {
	Id          int     `json:"id"`
	Subject     string  `json:"subject"`
	Description string  `json:"description"`
	ProjectId   int     `json:"project_id"`
	Project     IdName  `json:"project"`
	TrackerId   int     `json:"tracker_id"`
	Tracker     IdName  `json:"tracker"`
	ParentId    int     `json:"parent_issue_id,omitempty"`
	Parent      int     `json:"parent"`
	StatusId    int     `json:"status_id"`
	Status      IdName  `json:"status"`
	Priority    IdName  `json:"priority"`
	Author      IdName  `json:"author"`
	AssignedTo  IdName  `json:"assigned_to"`
	Children    []Issue `json:"children"`
}

type issueRequest struct {
	Issue Issue `json:"issue"`
}

// 未動作確認があれば、その親issue、いれば担当者を列挙する
func (parent Issue) CreateReleaseBlock(c Client) string {
	s := ""
	for _, child := range parent.Children {
		if checkReleaseBlock(child) {
			assigned := ""
			if child.AssignedTo.Name == "" {
				assigned = "なし"
			} else {
				assigned = child.AssignedTo.Name
			}
			s += c.Endpoint + "/issues/" + strconv.Itoa(parent.Id) + " における\n" + strconv.Itoa(child.Id) + " 担当者: " + assigned + " がブロック\n"
		}
	}
	return s
}

func checkReleaseBlock(child Issue) bool {
	r := regexp.MustCompile(`.*動作確認.*`)
	unfinStatus := []string{"新規", "進行中", "フィードバック"}
	return r.MatchString(child.Subject) && contains(unfinStatus, child.Status.Name)
}
