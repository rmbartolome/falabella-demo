package entity

import (
	"errors"
	"html"
	"strings"

	"github.com/gofrs/uuid"
)

type Client struct {
	ID        string  `gorm:"primary_key" json:"id"`
	FullName  string  `gorm:"size:255;not null" json:"fullname"`
	Rif       string  `gorm:"size:255;not null" json:"rif"`
	Address   string  `gorm:"size:255;not null" json:"address"`
	Cellphone float64 `json:"cellphone"`
	Email     string  `gorm:"size:255;not null" json:"email"`
}

func (d *Client) Prepare() {
	uuid, _ := uuid.NewV4()
	id := uuid.String()
	d.ID = id
	d.FullName = html.EscapeString(strings.TrimSpace(d.FullName))
	d.Rif = html.EscapeString(strings.TrimSpace(d.Rif))
	d.Address = html.EscapeString(strings.TrimSpace(d.Address))
	d.Email = html.EscapeString(strings.TrimSpace(d.Email))

}

func (d *Client) Validate() error {
	if d.Address == "" {
		return errors.New("required address")
	}
	if d.Email == "" {
		return errors.New("required email")
	}

	if d.FullName == "" {
		return errors.New("required fullname")
	}

	if d.Rif == "" {
		return errors.New("required rif")
	}

	if d.Cellphone == 0 {
		return errors.New("required cellphone")
	}

	return nil
}

type FindAllClientRequest struct {
}

type GetByIDClientRequest struct {
	ID string `json:"id"`
}

type CreateClientRequest struct {
	Client Client `json:"client"`
}

type DeleteClientRequest struct {
	ID string `json:"id"`
}

type FindAllClientResponse struct {
	Clients []Client `json:"clients"`
	Err     error    `json:"error,omitempty"`
}

func (r FindAllClientResponse) error() error { return r.Err }

type GetByIDClientResponse struct {
	Client Client `json:"client"`
	Err    error  `json:"error,omitempty"`
}

func (r GetByIDClientResponse) error() error { return r.Err }

type CreateClientResponse struct {
	Err error `json:"error,omitempty"`
}

func (r CreateClientResponse) error() error { return r.Err }

type DeleteClientResponse struct {
	Err error `json:"error,omitempty"`
}

func (r DeleteClientResponse) error() error { return r.Err }
