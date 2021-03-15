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
	Id          int    `json:"id"`
	Subject     string `json:"subject"`
	Description string `json:"description"`
	ProjectId   int    `json:"project_id"`
	// Project    *IdName `json:"project"`
	Project   IdName `json:"project"`
	TrackerId int    `json:"tracker_id"`
	Tracker   IdName `json:"tracker"`
	ParentId  int    `json:"parent_issue_id,omitempty"`
	Parent    int    `json:"parent"`
	StatusId  int    `json:"status_id"`
	Status    IdName `json:"status"`
	// PriorityId int `json:"priority_id,omitempty"`
	Priority IdName  `json:"priority"`
	Author   IdName  `json:"author"`
	Children []Child `json:"children"`
}

type issueRequest struct {
	Issue Issue `json:"issue"`
}
