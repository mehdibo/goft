package ftapi

import "time"

// User represents a user entity
type User struct {
	ID int `json:"id,omitempty"`
	Login string `json:"login,omitempty"`
	Email string `json:"email"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Kind string `json:"kind"`
	CampusID int `json:"campus_id"`
	URL string `json:"url,omitempty"`
}

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
