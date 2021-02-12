package entity

import (
	"errors"
	"html"
	"strings"

	"github.com/gofrs/uuid"
)

type TypeProd struct {
	ID   string `gorm:"primary_key" json:"id"`
	Name string `gorm:"size:255;not null" json:"name"`
}

func (d *TypeProd) Prepare() {
	uuid, _ := uuid.NewV4()
	id := uuid.String()
	d.ID = id
	d.Name = html.EscapeString(strings.TrimSpace(d.Name))

}

func (d *TypeProd) Validate() error {
	if d.Name == "" {
		return errors.New("required name")
	}

	return nil
}

type FindAllTypeProdRequest struct {
}

type GetByIDTypeProdRequest struct {
	ID string `json:"id"`
}

type CreateTypeProdRequest struct {
	TypeProd TypeProd `json:"typeProd"`
}

type DeleteTypeProdRequest struct {
	ID string `json:"id"`
}

type FindAllTypeProdResponse struct {
	TypeProds []TypeProd `json:"typeProds"`
	Err       error      `json:"error,omitempty"`
}

func (r FindAllTypeProdResponse) error() error { return r.Err }

type GetByIDTypeProdResponse struct {
	TypeProd TypeProd `json:"typeProd"`
	Err      error    `json:"error,omitempty"`
}

func (r GetByIDTypeProdResponse) error() error { return r.Err }

type CreateTypeProdResponse struct {
	Err error `json:"error,omitempty"`
}

func (r CreateTypeProdResponse) error() error { return r.Err }

type DeleteTypeProdResponse struct {
	Err error `json:"error,omitempty"`
}

func (r DeleteTypeProdResponse) error() error { return r.Err }
