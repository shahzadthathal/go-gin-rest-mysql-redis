package model

// MUserDetail struct
type MUserDetail struct {
	ID      int64  `json:"id"`
	Address string `json:"address"`
	DOB     string `json:"dob"`
	POB     string `json:"pob"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	UserID  int64  `json:"userId"`
	MUser   MUser  `json:"user"`
}

// MUserDetails array of MUserDetail
type MUserDetails []MUserDetail
