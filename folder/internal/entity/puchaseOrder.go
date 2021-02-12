package entity

import (
	"errors"
	"html"
	"strings"

	"github.com/gofrs/uuid"
)

type PuchaseOrder struct {
	ID                string  `gorm:"primary_key" json:"id"`
	CodProd           string  `gorm:"size:255;not null" json:"codprod"`
	Provider          string  `gorm:"size:255;not null" json:"provider"`
	Price             float64 `json:"price"`
	ReplenishQuantity float64 `json:"replenishquantity"`
}

func (d *PuchaseOrder) Prepare() {
	uuid, _ := uuid.NewV4()
	id := uuid.String()
	d.ID = id
	d.CodProd = html.EscapeString(strings.TrimSpace(d.CodProd))
	d.Provider = html.EscapeString(strings.TrimSpace(d.Provider))

}

func (d *PuchaseOrder) Validate() error {
	if d.CodProd == "" {
		return errors.New("required product code")
	}
	if d.Provider == "" {
		return errors.New("required provider")
	}
	if d.Price == 0 {
		return errors.New("required price")
	}
	if d.ReplenishQuantity == 0 {
		return errors.New("required quantity replenish")
	}

	return nil
}

type FindAllPuchaseOrderRequest struct {
}

type GetByIDPuchaseOrderRequest struct {
	ID string `json:"id"`
}

type CreatePuchaseOrderRequest struct {
	PuchaseOrder PuchaseOrder `json:"puchaseOrder"`
}

type DeletePuchaseOrderRequest struct {
	ID string `json:"id"`
}

type FindAllPuchaseOrderResponse struct {
	PuchaseOrders []PuchaseOrder `json:"puchaseOrders"`
	Err           error          `json:"error,omitempty"`
}

func (r FindAllPuchaseOrderResponse) error() error { return r.Err }

type GetByIDPuchaseOrderResponse struct {
	PuchaseOrder PuchaseOrder `json:"puchaseOrder"`
	Err          error        `json:"error,omitempty"`
}

func (r GetByIDPuchaseOrderResponse) error() error { return r.Err }

type CreatePuchaseOrderResponse struct {
	Err error `json:"error,omitempty"`
}

func (r CreatePuchaseOrderResponse) error() error { return r.Err }

type DeletePuchaseOrderResponse struct {
	Err error `json:"error,omitempty"`
}

func (r DeletePuchaseOrderResponse) error() error { return r.Err }
