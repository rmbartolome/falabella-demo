package entity

import (
	"errors"
	"html"
	"strings"

	"github.com/gofrs/uuid"
)

type Users struct {
	ID       string `gorm:"primary_key" json:"id"`
	TypeUser string `gorm:"size:255;not null" json:"typeuser"`
}

func (d *Users) Prepare() {
	uuid, _ := uuid.NewV4()
	id := uuid.String()
	d.ID = id
	d.TypeUser = html.EscapeString(strings.TrimSpace(d.TypeUser))

}

func (d *Users) Validate() error {
	if d.TypeUser == "" {
		return errors.New("required type user")
	}

	return nil
}

type FindAllUsersRequest struct {
}

type GetByIDUsersRequest struct {
	ID string `json:"id"`
}

type CreateUsersRequest struct {
	Users Users `json:"users"`
}

type DeleteUsersRequest struct {
	ID string `json:"id"`
}

type FindAllUsersResponse struct {
	Userss []Users `json:"userss"`
	Err    error   `json:"error,omitempty"`
}

func (r FindAllUsersResponse) error() error { return r.Err }

type GetByIDUsersResponse struct {
	Users Users `json:"users"`
	Err   error `json:"error,omitempty"`
}

func (r GetByIDUsersResponse) error() error { return r.Err }

type CreateUsersResponse struct {
	Err error `json:"error,omitempty"`
}

func (r CreateUsersResponse) error() error { return r.Err }

type DeleteUsersResponse struct {
	Err error `json:"error,omitempty"`
}

func (r DeleteUsersResponse) error() error { return r.Err }
