package ftapi

import "time"

type projectSamary struct {
	Name string `json:"name,omitempty"`
	ID   int    `json:"id,omitempty"`
	Slug string `json:"slug,omitempty"`
	URL  string `json:"url,omitempty"`
}

type scale struct {
	ID               int  `json:"id,omitempty"`
	CorrectionNumber int  `json:"correction_number"`
	IsPrimary        bool `json:"is_primary"`
}

type upload struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type session struct {
	ID                int       `json:"id,omitempty"`
	Solo              bool      `json:"solo"`
	BeginAt           time.Time `json:"begin_at,omitempty"`
	EndAt             time.Time `json:"end_at,omitempty"`
	EstimateTime      string    `json:"estimate_time,omitempty"`
	Difficulty        int       `json:"difficulty,omitempty"`
	Objectives        []string  `json:"objectives,omitempty"`
	Description       string    `json:"description,omitempty"`
	DurationDays      string    `json:"duration_days,omitempty"` // TODO
	TerminateingAfter int       `json:"terminating_after,omitempty"`
	ProjectID         int       `json:"project_id,omitempty"`
	CampusID          int       `json:"campus_id,omitempty"`
	CursusID          int       `json:"cursus_id,omitempy"`
	CreatedAt         time.Time `json:"created_at,omitempty"`
	UpdatedAt         time.Time `json:"updated_at,omitempty"`
	MaxPeople         int       `json:"max_poeple,omitepmpty"`
	IsSubscriptable   bool      `json:"is_subscriptable"`
	Scales            []*scale  `json:"scales,omitempty"`
	Uploads           []*upload `json:"uploads,omitempty"`
	TeamBehaviour     string    `json:"team_behaviour,omitempty"`
	Commit            string    `json:"commit,omitempty"`
}

type Project struct {
	ID              int              `json:"id,omitempty"`
	Name            string           `json:"name,omitempty"`
	Slug            string           `json:"slug,omitempty"`
	Parent          *projectSamary   `json:"parent,omitempty"`
	Children        []*projectSamary `json:"children,omitempty"`
	Attachments     []string         `json:"attachments,omitempty"` // TODO
	CreatedAt       time.Time        `json:"created_at,omitempty"`
	UpdatedAt       time.Time        `json:"updated_at,omitempty"`
	Exam            bool             `json:"exam"`
	GitID           int              `json:"git_id,omitempty"`
	Repogitory      string           `json:"repository,omitempty"`
	Cursus          []*cursus        `json:"cursus,omitempty"`
	Campus          []*Campus        `json:"campus,omitempty"`
	Videos          []string         `json:"videos,omitempty"` // TODO
	ProjectSessions []*session       `json:"project_sessions,omitempty"`
}
