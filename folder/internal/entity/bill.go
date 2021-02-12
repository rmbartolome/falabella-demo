package entity

import (
	"errors"
	"html"
	"strings"

	"github.com/gofrs/uuid"
)

type Bill struct {
	ID            string  `gorm:"primary_key" json:"id"`
	CodClient     string  `gorm:"size:255;not null" json:"codClient"`
	Date          string  `gorm:"size:255;not null" json:"date"`
	Address       string  `gorm:"size:255;not null" json:"address"`
	TotalProdList float64 `json:"totalprodlist"`
	SubTotal      float64 `json:"subtotal"`
	Texes         float64 `json:"texes"`
	FinalPrice    float64 `json:"finalprice"`
}

func (d *Bill) Prepare() {
	uuid, _ := uuid.NewV4()
	id := uuid.String()
	d.ID = id
	d.CodClient = html.EscapeString(strings.TrimSpace(d.CodClient))
	d.Address = html.EscapeString(strings.TrimSpace(d.Address))

}

func (d *Bill) Validate() error {
	if d.CodClient == "" {
		return errors.New("required type user")
	}
	if d.Address == "" {
		return errors.New("required address")
	}
	if d.Date == "" {
		return errors.New("required date")
	}
	if d.FinalPrice == 0 {
		return errors.New("required final price")
	}
	if d.SubTotal == 0 {
		return errors.New("required sub total")
	}
	if d.Texes == 0 {
		return errors.New("required texes")
	}
	if d.TotalProdList == 0 {
		return errors.New("required total product list")
	}

	return nil
}

type FindAllBillRequest struct {
}

type GetByIDBillRequest struct {
	ID string `json:"id"`
}

type CreateBillRequest struct {
	Bill Bill `json:"bill"`
}

type DeleteBillRequest struct {
	ID string `json:"id"`
}

type FindAllBillResponse struct {
	Bills []Bill `json:"bills"`
	Err   error  `json:"error,omitempty"`
}

func (r FindAllBillResponse) error() error { return r.Err }

type GetByIDBillResponse struct {
	Bill Bill  `json:"bill"`
	Err  error `json:"error,omitempty"`
}

func (r GetByIDBillResponse) error() error { return r.Err }

type CreateBillResponse struct {
	Err error `json:"error,omitempty"`
}

func (r CreateBillResponse) error() error { return r.Err }

type DeleteBillResponse struct {
	Err error `json:"error,omitempty"`
}

func (r DeleteBillResponse) error() error { return r.Err }
