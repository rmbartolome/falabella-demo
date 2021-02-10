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

type Client struct {
	ID        string  `gorm:"primary_key" json:"id"`
	FullName  string  `gorm:"size:255;not null" json:"fullname"`
	Rif       string  `gorm:"size:255;not null" json:"rif"`
	Address   string  `gorm:"size:255;not null" json:"address"`
	Cellphone float64 `json:"cellphone"`
	Email     string  `gorm:"size:255;not null" json:"email"`
}

type TypeProd struct {
	ID   string `gorm:"primary_key" json:"id"`
	Name string `gorm:"size:255;not null" json:"name"`
}

type PuchaseOrder struct {
	ID                string  `gorm:"primary_key" json:"id"`
	codProd           string  `gorm:"size:255;not null" json:"codprod"`
	Provider          string  `gorm:"size:255;not null" json:"provider"`
	Price             float64 `json:"price"`
	ReplenishQuantity float64 `json:"replenishquantity"`
}

type DeliveryOrder struct {
	ID          string `gorm:"primary_key" json:"id"`
	codBill     string `gorm:"size:255;not null" json:"codBill"`
	Address     string `gorm:"size:255;not null" json:"address"`
	Responsable string `gorm:"size:255;not null" json:"responsable"`
}
type Bill struct {
	ID            string  `gorm:"primary_key" json:"id"`
	codClient     string  `gorm:"size:255;not null" json:"codClient"`
	Date          string  `gorm:"size:255;not null" json:"date"`
	Address       string  `gorm:"size:255;not null" json:"address"`
	TotalProdList float64 `json:"totalprodlist"`
	SubTotal      float64 `json:"subtotal"`
	Texes         float64 `json:"texes"`
	FinalPrice    float64 `json:"finalprice"`
}
type Product struct {
	ID          string  `gorm:"primary_key" json:"id"`
	codTypeProd string  `gorm:"size:255;not null" json:"codtypeprod"`
	Name        string  `gorm:"size:255;not null" json:"name"`
	Maker       string  `gorm:"size:255;not null" json:"maker"`
	Model       string  `gorm:"size:255;not null" json:"model"`
	Year        string  `gorm:"size:255;not null" json:"year"`
	Price       float64 `json:"price"`
}

type Provider struct {
	ID           string  `gorm:"primary_key" json:"id"`
	codeTypeProd string  `gorm:"size:255;not null" json:"codetypeprod"`
	Name         string  `gorm:"size:255;not null" json:"name"`
	Address      string  `gorm:"size:255;not null" json:"address"`
	Cellphone    float64 `json:"cellphone"`
}

type Stock struct {
	ID            string  `gorm:"primary_key" json:"id"`
	ProductName   string  `gorm:"size:255;not null" json:"productname"`
	TypeProdName  string  `gorm:"size:255;not null" json:"typeprodname"`
	QuantityExist float64 `json:"quantityexist"`
	PurchasePrice float64 `json:"purchaseprice"`
	SalePrice     float64 `json:"saleprice"`
}

//INGRESO DE PREPARE PARA CADA TABLA
func (d *Users) Prepare() {
	uuid, _ := uuid.NewV4()
	id := uuid.String()
	d.ID = id
	d.TypeUser = html.EscapeString(strings.TrimSpace(d.TypeUser))

}

//INGRESO DE VALIDATE PARA CADA TABLA
func (d *Users) Validate() error {
	if d.TypeUser == "" {
		return errors.New("required type user")
	}

	return nil
}
