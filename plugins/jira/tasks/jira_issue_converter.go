package tasks

import (
	lakeModels "github.com/merico-dev/lake/models"
	"github.com/merico-dev/lake/models/domainlayer/didgen"
	"github.com/merico-dev/lake/models/domainlayer/ticket"
	jiraModels "github.com/merico-dev/lake/plugins/jira/models"
	"gorm.io/gorm/clause"
)

func ConvertIssues(sourceId uint64, boardId uint64) error {

	jiraIssue := &jiraModels.JiraIssue{}
	// select all issues belongs to the board
	cursor, err := lakeModels.Db.Model(jiraIssue).
		Select("jira_issues.*").
		Joins("left join jira_board_issues on jira_board_issues.issue_id = jira_issues.issue_id").
		Where("jira_board_issues.source_id = ? AND jira_board_issues.board_id = ?", sourceId, boardId).
		Rows()
	if err != nil {
		return err
	}
	defer cursor.Close()

	issueIdGen := didgen.NewDomainIdGenerator(&jiraModels.JiraIssue{})
	userIdGen := didgen.NewDomainIdGenerator(&jiraModels.JiraUser{})

	boardIssue := &ticket.BoardIssue{
		BoardId: didgen.NewDomainIdGenerator(&jiraModels.JiraBoard{}).Generate(sourceId, boardId),
	}
	// clearn up board issue relationship altogether
	err = lakeModels.Db.Exec("DELETE from board_issues where board_id = ?", boardIssue.BoardId).Error
	if err != nil {
		return err
	}

	// iterate all rows
	for cursor.Next() {
		err = lakeModels.Db.ScanRows(cursor, jiraIssue)
		if err != nil {
			return err
		}
		issue := &ticket.Issue{
			Id:                      issueIdGen.Generate(jiraIssue.SourceId, jiraIssue.IssueId),
			Url:                     jiraIssue.Self,
			Key:                     jiraIssue.Key,
			Summary:                 jiraIssue.Summary,
			EpicKey:                 jiraIssue.EpicKey,
			Type:                    jiraIssue.StdType,
			Status:                  jiraIssue.StdStatus,
			StoryPoint:              jiraIssue.StdStoryPoint,
			OriginalEstimateMinutes: jiraIssue.OriginalEstimateMinutes,
			CreatorId:               userIdGen.Generate(sourceId, jiraIssue.CreatorAccountId),
			ResolutionDate:          jiraIssue.ResolutionDate,
			Priority:                jiraIssue.PriorityName,
			CreatedDate:             jiraIssue.Created,
			UpdatedDate:             jiraIssue.Updated,
			LeadTimeMinutes:         jiraIssue.LeadTimeMinutes,
		}
		if jiraIssue.AssigneeAccountId != "" {
			issue.AssigneeId = userIdGen.Generate(sourceId, jiraIssue.AssigneeAccountId)
		}
		if jiraIssue.ParentId != 0 {
			issue.ParentIssueId = issueIdGen.Generate(sourceId, jiraIssue.ParentId)
		}

		err = lakeModels.Db.Clauses(clause.OnConflict{UpdateAll: true}).Create(issue).Error
		if err != nil {
			return err
		}

		// convert board issue relationship
		boardIssue.IssueId = issue.Id
		err = lakeModels.Db.Clauses(clause.OnConflict{DoNothing: true}).Create(boardIssue).Error
		if err != nil {
			return err
		}
	}
	return nil
}
