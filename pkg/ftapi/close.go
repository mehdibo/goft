package ftapi

import "time"

type communityService struct {
	ID int `json:"id,omitempty"`
	Duration int64 `json:"duration,omitempty"`
	ScheduledAt time.Time `json:"schedule_at,omitempty"`
	Occupation string `json:"occupation,omitempty"`
	State string `json:"state,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// Close represents a close
type Close struct {
	ID int `json:"id,omitempty"`
	Kind string `json:"kind,omitempty"`
	Reason string `json:"reason,omitempty"`
	State string `json:"state,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	CommunityServices []*communityService `json:"community_services,omitempty"`
	User *User `json:"user,omitempty"`
	Closer *User `json:"closer,omitempty"`
}
