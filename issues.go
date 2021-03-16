package main

type IdName struct {
	Id   int
	Name string
}

type Child struct {
	Subject string `json:"subject"`
	Id      int    `json:"id"`
	Tracker IdName `json:"tracker"`
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

type BlockIssue struct {
	Id         int
	Subject    string
	Status     IdName
	AssignedTo IdName
}

type BlockStory struct {
	BlockStory BlockIssue
	BlockTasks []BlockIssue
}
