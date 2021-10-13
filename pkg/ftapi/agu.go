package ftapi

import "time"

// Agu represents an AGU entity
type Agu struct {
	ID int `json:"id,omitempty"`
	UserID int `json:"user_id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	CloseID int `json:"close_id,omitempty"`
	IsFree bool `json:"is_free,omitempty"`
	Reason string `json:"reason,omitempty"`
	EndDate string `json:"end_date,omitempty"`
	ExpectedEndDate string `json:"expected_end_date,omitempty"`
	BeginDate string `json:"begin_date,omitempty"`
	AguID int `json:"anti_grav_unit_id,omitempty"`
	InternshipId int `json:"internship_id,omitempty"`
}
