package entity

import (
	"errors"
	"html"
	"strings"

	"github.com/gofrs/uuid"
)

type Provider struct {
	ID           string  `gorm:"primary_key" json:"id"`
	CodeTypeProd string  `gorm:"size:255;not null" json:"codetypeprod"`
	Name         string  `gorm:"size:255;not null" json:"name"`
	Address      string  `gorm:"size:255;not null" json:"address"`
	Cellphone    float64 `json:"cellphone"`
}

func (d *Provider) Prepare() {
	uuid, _ := uuid.NewV4()
	id := uuid.String()
	d.ID = id
	d.CodeTypeProd = html.EscapeString(strings.TrimSpace(d.CodeTypeProd))
	d.Name = html.EscapeString(strings.TrimSpace(d.Name))
	d.Address = html.EscapeString(strings.TrimSpace(d.Address))

}

func (d *Provider) Validate() error {
	if d.CodeTypeProd == "" {
		return errors.New("required code type provider")
	}
	if d.Name == "" {
		return errors.New("required name")
	}
	if d.Address == "" {
		return errors.New("required address")
	}
	if d.Cellphone == 0 {
		return errors.New("required cellphone")
	}

	return nil
}

type FindAllProviderRequest struct {
}

type GetByIDProviderRequest struct {
	ID string `json:"id"`
}

type CreateProviderRequest struct {
	Provider Provider `json:"provider"`
}

type DeleteProviderRequest struct {
	ID string `json:"id"`
}

type FindAllProviderResponse struct {
	Providers []Provider `json:"providers"`
	Err       error      `json:"error,omitempty"`
}

func (r FindAllProviderResponse) error() error { return r.Err }

type GetByIDProviderResponse struct {
	Provider Provider `json:"provider"`
	Err      error    `json:"error,omitempty"`
}

func (r GetByIDProviderResponse) error() error { return r.Err }

type CreateProviderResponse struct {
	Err error `json:"error,omitempty"`
}

func (r CreateProviderResponse) error() error { return r.Err }

type DeleteProviderResponse struct {
	Err error `json:"error,omitempty"`
}

func (r DeleteProviderResponse) error() error { return r.Err }
