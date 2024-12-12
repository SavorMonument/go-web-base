package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRole string

const (
	SUPERADMIN_ROLE UserRole = "SuperAdminRole"
	ORGADMIN_ROLE   UserRole = "OrgAdminRole"
	USER_ROLE       UserRole = "UserRole"
)

type User struct {
	gorm.Model
	Username string       `json:"id" schema:"username"`
	Password []byte       `json:"password" schema:"password"`
	Role     UserRole     `json:"role" schema:"role"`
	Settings UserSettings `json:"settings" schema:"settings" gorm:"embedded"`
}

type UserSettings struct {
	Language *Language `json:"language" schema:"language" gorm:"embedded"`
}

func (u *User) MergeDefaults() {
	if u.Settings.Language == nil {
		u.Settings.Language = &EN_LANG
	}
}

func NewSuperAdminUser() *User {

	pass, _ := bcrypt.GenerateFromPassword([]byte("Admin"), 15)
	return &User{
		Username: "Admin",
		Role:     SUPERADMIN_ROLE,
		Password: pass,
		Settings: UserSettings{
			Language: &EN_LANG,
		},
	}
}

func (u *User) CanDelete() bool {
	if u.IsSuperAdmin() {
		return false
	}
	return true
}

func (u *User) IsSuperAdmin() bool {
	return u.Role == SUPERADMIN_ROLE
}

func (u *User) IsOrgAdmin() bool {
	return u.Role == ORGADMIN_ROLE
}

func (u *User) IsAdmin() bool {
	return u.Role == SUPERADMIN_ROLE
}
