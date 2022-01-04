package ticket

import (
	"time"

	"github.com/merico-dev/lake/models/domainlayer"
)

type Issue struct {
	domainlayer.DomainEntity
	Url                     string
	Key                     string
	Title                   string
	Summary                 string
	EpicKey                 string
	Type                    string
	Status                  string
	StoryPoint              uint
	ResolutionDate          *time.Time
	CreatedDate             *time.Time
	UpdatedDate             *time.Time
	LeadTimeMinutes         uint
	ParentIssueId           string
	Priority                string
	OriginalEstimateMinutes int64
	TimeRemainingMinutes    int64
	CreatorId               string
	AssigneeId              string
	OwnerId                 string
}
