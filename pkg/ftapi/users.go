package ftapi

// User represents a user entity
type User struct {
	ID int `json:"id,omitempty"`
	Login string `json:"login,omitempty"`
	Email string `json:"email"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Kind string `json:"kind"`
	CampusID int `json:"campus_id"`
	Url string `json:"url,omitempty"`
}