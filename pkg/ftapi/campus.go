package ftapi

import "time"

type language struct {
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	ISOIdentifier string `json:"identifier,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// Campus represents a campus entity
type Campus struct {
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	TimeZone string `json:"time_zone,omitempty"`
	Language *language `json:"language,omitempty"`
	UsersCount int `json:"users_count,omitempty"`
	VogsphereID int `json:"vogsphere_id,omitempty"`
}
