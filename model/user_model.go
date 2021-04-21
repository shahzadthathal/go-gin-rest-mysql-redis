package model

// MUser struct
type MUser struct {
	ID                 int64  `json:"id" example:"1"`
	UserName           string `json:"userName" example:"userlogin"`
	Password           string `json:"password" example:"passlogin"`
	AccountExpired     bool   `json:"accountExpired" example:"false"`
	AccountLocked      bool   `json:"accountLocked" example:"false"`
	CredentialsExpired bool   `json:"credentialsExpired" example:"false"`
	Enabled            bool   `json:"enabled" example:"true"`
}

// MUsers array of MUser type
type MUsers []MUser
