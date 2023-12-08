package models

import (
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Model
	FirstName       string `json:"first_name" query:"first_name" form:"first_name" validate:"required"`
	LastName        string `json:"last_name" query:"last_name" form:"last_name" validate:"required"`
	Email           string `json:"email" gorm:"uniqueIndex" query:"email" form:"email" validate:"required,email"`
	PasswordHash    string `json:"password_hash" query:"-" form:"-"`
	Password        string `json:"password" gorm:"-" query:"password" form:"password" validate:"required,min=8,max=64"`
	PasswordConfirm string `json:"-" gorm:"-" query:"password_confirm" form:"password_confirm" validate:"required,min=8,max=64"`
	CompanyID       uuid.UUID
	Company         *Company
	Permissions     []*Permission `json:"permissions" query:"-" form:"-" gorm:"many2many:user_languages;"`
}

func (u *User) FlatPermissions(db *gorm.DB) ([]string, error) {
	var permissions []string
	var allPermission []Permission
	err := db.Find(&allPermission).Error
	if err != nil {
		return permissions, err
	}
	permissionSet := make(map[string]struct{})
	for _, permission := range u.Permissions {
		flatten, err := permission.Flatten(allPermission)
		if err != nil {
			return permissions, err
		}
		for _, p := range flatten {
			permissionSet[p] = struct{}{}
		}
	}
	for permission := range permissionSet {
		permissions = append(permissions, permission)
	}
	return permissions, nil
}

type Permission struct {
	Model
	Permission        string
	ParentPermissions []*Permission `gorm:"many2many:permission_parent_permissions;"`
}

func (p *Permission) Flatten(allPermissions []Permission) ([]string, error) {
	var permissions []string
	permissionsSet := make(map[string]struct{})
	for _, permission := range allPermissions {
		if permission.ID == p.ID {
			for _, parentPermission := range p.ParentPermissions {
				flatten, err := parentPermission.Flatten(allPermissions)
				if err != nil {
					return permissions, err
				}
				for _, p := range flatten {
					permissionsSet[p] = struct{}{}
				}
			}
		}
	}
	for permission := range permissionsSet {
		permissions = append(permissions, permission)
	}
	permissions = append(permissions, p.Permission)
	return permissions, nil
}

type Role struct {
	Model
	Name string `json:"name" query:"name" form:"name" validate:"required"`
}

type Company struct {
	Model
	Name              string   `json:"name" query:"name" form:"name" validate:"required" yaml:"name"`
	InternalReference string   `json:"internal_reference" query:"internal_reference" form:"internal_reference" validate:"required" yaml:"internal_reference"`
	CompanyCRN        string   `json:"company_crn" query:"company_crn" form:"company_crn" validate:"required" yaml:"company_crn"`
	VAT               string   `json:"vat" query:"vat" form:"vat" validate:"required" yaml:"vat"`
	Address           *Address `json:"address" query:"-" form:"-" yaml:"address"`
	Users             []*User  `json:"users" query:"-" form:"-"`
}

type Address struct {
	Model
	CompanyID uuid.UUID
	Company   *Company
	Line1     string `json:"line_1" query:"line_1" form:"line_1" yaml:"line_1" validate:"required"`
	Line2     string `json:"line_2" query:"line_2" form:"line_2" yaml:"line_2"`
	Town      string `json:"town" query:"town" form:"town" validate:"required" yaml:"town"`
	County    string `json:"county" query:"county" form:"county" validate:"required" yaml:"county"`
	Postcode  string `json:"postcode" query:"postcode" form:"postcode" validate:"required" yaml:"postcode"`
}

func (a *Address) String() string {
	var b strings.Builder
	b.WriteString(a.Line1)
	b.WriteString(", ")
	if a.Line2 != "" {
		b.WriteString(a.Line2)
		b.WriteString(", ")
	}
	if a.Town != "" {
		b.WriteString(a.Town)
		b.WriteString(", ")
	}
	if a.County != "" {
		b.WriteString(a.County)
		b.WriteString(", ")
	}
	b.WriteString(a.Postcode)
	return b.String()
}
