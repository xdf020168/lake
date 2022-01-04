package factory

import (
	"github.com/merico-dev/lake/models/domainlayer/ticket"
)

func CreateWorklog(boardId string, issueId string) (*ticket.Worklog, error) {
	worklog := &ticket.Worklog{
		Id:       RandIntString(),
		IssueId:  issueId, // ref to issue
		AuthorId: "",
	}
	return worklog, nil
}
