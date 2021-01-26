package ftapi

import "time"

type cursus struct {
	ID int `json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	Name string `json:"name,omitempty"`
	Slug string `json:"slug,omitempty"`
}

type skill struct {
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Level float32 `json:"level,omitempty"`
}

type cursusUser struct {
	ID int `json:"id,omitempty"`
	Grade string `json:"grade,omitempty"`
	Level float32 `json:"level,omitempty"`
	Skills []*skill `json:"skills,omitempty"`
	BlackholedAt *time.Time `json:"blackholed_at,omitempty"`
	BeginAt *time.Time `json:"begin_at,omitempty"`
	EndAt *time.Time `json:"end_at,omitempty"`
	HasCoalition bool `json:"has_coalition,omitempty"`
	Cursus *cursus `json:"cursus,omitempty"`
}

type campusUser struct {
	ID int `json:"id,omitempty"`
	UserID int `json:"user_id,omitempty"`
	CampusID int `json:"campus_id,omitempty"`
	IsPrimary bool `json:"is_primary,omitempty"`
}

type role struct {
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`

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
	Campuses []*Campus `json:"campus,omitempty"`
	CampusUsers []*campusUser `json:"campus_users,omitempty"`
	Roles []*role `json:"role,omitempty"`
	CursusUsers []*cursusUser `json:"cursus_user,omitempty"`
}

// GetPrimaryCampus returns the user's primary campus or nil if none found
func (u *User) GetPrimaryCampus() *Campus {
	if u.CampusUsers == nil || u.Campuses == nil {
		return nil
	}
	var primaryCampusID int
	for _, campusUser := range u.CampusUsers {
		if !campusUser.IsPrimary {
			continue
		}
		primaryCampusID = campusUser.CampusID
		break
	}
	for _, campus := range u.Campuses {
		if campus.ID == primaryCampusID {
			return campus
		}
	}
	return nil
}
