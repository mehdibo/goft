package ftapi

import "time"

type language struct {
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	ISOIdentifier string `json:"identifier,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type campus struct {
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	TimeZone string `json:"time_zone,omitempty"`
	Language *language `json:"language,omitempty"`
	UsersCount int `json:"users_count,omitempty"`
	VogsphereID int `json:"vogsphere_id,omitempty"`
}

type campusUser struct {
	ID int `json:"id,omitempty"`
	UserID int `json:"user_id,omitempty"`
	CampusID int `json:"campus_id,omitempty"`
	IsPrimary bool `json:"is_primary,omitempty"`
}

// User represents a user entity
type User struct {
	ID int `json:"id,omitempty"`
	Login string `json:"login,omitempty"`
	Email string `json:"email,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName string `json:"last_name,omitempty"`
	UsualFirstName string `json:"usual_first_name,omitempty"`
	Phone string `json:"phone,omitempty"`
	ImageURL string `json:"image_url,omitempty"`
	IsStaff bool `json:"staff?,omitempty"`
	Kind string `json:"kind,omitempty"`
	CampusID int `json:"campus_id,omitempty"`
	URL string `json:"url,omitempty"`
	PoolMonth string `json:"pool_month,omitempty"`
	PoolYear string `json:"pool_year,omitempty"`
	Campuses []*campus `json:"campus,omitempty"`
	CampusUsers []*campusUser `json:"campus_users,omitempty"`
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
