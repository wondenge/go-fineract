package models

// GetUsersResponse GetUsersResponse
type GetUsersResponse struct {
	Id                   int64      `json:"id,omitempty"`
	Username             string     `json:"username,omitempty"`
	OfficeId             int64      `json:"officeId,omitempty"`
	OfficeName           string     `json:"officeName,omitempty"`
	Firstname            string     `json:"firstname,omitempty"`
	Lastname             string     `json:"lastname,omitempty"`
	Email                string     `json:"email,omitempty"`
	PasswordNeverExpires bool       `json:"passwordNeverExpires,omitempty"`
	Staff                StaffData  `json:"staff,omitempty"`
	SelectedRoles        []RoleData `json:"selectedRoles,omitempty"`
}
