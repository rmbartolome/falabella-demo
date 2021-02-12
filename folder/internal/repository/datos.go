package repository

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
	CodProd           string  `gorm:"size:255;not null" json:"codprod"`
	Provider          string  `gorm:"size:255;not null" json:"provider"`
	Price             float64 `json:"price"`
	ReplenishQuantity float64 `json:"replenishquantity"`
}

type DeliveryOrder struct {
	ID          string `gorm:"primary_key" json:"id"`
	CodBill     string `gorm:"size:255;not null" json:"codBill"`
	Address     string `gorm:"size:255;not null" json:"address"`
	Responsable string `gorm:"size:255;not null" json:"responsable"`
}

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
type Product struct {
	ID          string  `gorm:"primary_key" json:"id"`
	CodTypeProd string  `gorm:"size:255;not null" json:"codtypeprod"`
	Name        string  `gorm:"size:255;not null" json:"name"`
	Maker       string  `gorm:"size:255;not null" json:"maker"`
	Model       string  `gorm:"size:255;not null" json:"model"`
	Year        string  `gorm:"size:255;not null" json:"year"`
	Price       float64 `json:"price"`
}

type Provider struct {
	ID           string  `gorm:"primary_key" json:"id"`
	CodeTypeProd string  `gorm:"size:255;not null" json:"codetypeprod"`
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

func (d *Client) Prepare() {
	uuid, _ := uuid.NewV4()
	id := uuid.String()
	d.ID = id
	d.FullName = html.EscapeString(strings.TrimSpace(d.FullName))
	d.Rif = html.EscapeString(strings.TrimSpace(d.Rif))
	d.Address = html.EscapeString(strings.TrimSpace(d.Address))
	d.Email = html.EscapeString(strings.TrimSpace(d.Email))

}

func (d *TypeProd) Prepare() {
	uuid, _ := uuid.NewV4()
	id := uuid.String()
	d.ID = id
	d.Name = html.EscapeString(strings.TrimSpace(d.Name))

}
func (d *PuchaseOrder) Prepare() {
	uuid, _ := uuid.NewV4()
	id := uuid.String()
	d.ID = id
	d.CodProd = html.EscapeString(strings.TrimSpace(d.CodProd))
	d.Provider = html.EscapeString(strings.TrimSpace(d.Provider))

}
func (d *DeliveryOrder) Prepare() {
	uuid, _ := uuid.NewV4()
	id := uuid.String()
	d.ID = id
	d.CodBill = html.EscapeString(strings.TrimSpace(d.CodBill))
	d.Responsable = html.EscapeString(strings.TrimSpace(d.Responsable))
	d.Address = html.EscapeString(strings.TrimSpace(d.Address))

}

func (d *Bill) Prepare() {
	uuid, _ := uuid.NewV4()
	id := uuid.String()
	d.ID = id
	d.CodClient = html.EscapeString(strings.TrimSpace(d.CodClient))
	d.Address = html.EscapeString(strings.TrimSpace(d.Address))

}

func (d *Product) Prepare() {
	uuid, _ := uuid.NewV4()
	id := uuid.String()
	d.ID = id
	d.CodTypeProd = html.EscapeString(strings.TrimSpace(d.CodTypeProd))
	d.Name = html.EscapeString(strings.TrimSpace(d.Name))
	d.Maker = html.EscapeString(strings.TrimSpace(d.Maker))
	d.Model = html.EscapeString(strings.TrimSpace(d.Model))
	d.Year = html.EscapeString(strings.TrimSpace(d.Year))

}
func (d *Provider) Prepare() {
	uuid, _ := uuid.NewV4()
	id := uuid.String()
	d.ID = id
	d.CodeTypeProd = html.EscapeString(strings.TrimSpace(d.CodeTypeProd))
	d.Name = html.EscapeString(strings.TrimSpace(d.Name))
	d.Address = html.EscapeString(strings.TrimSpace(d.Address))

}
func (d *Stock) Prepare() {
	uuid, _ := uuid.NewV4()
	id := uuid.String()
	d.ID = id
	d.ProductName = html.EscapeString(strings.TrimSpace(d.ProductName))
	d.TypeProdName = html.EscapeString(strings.TrimSpace(d.TypeProdName))

}

//INGRESO DE VALIDATE PARA CADA TABLA
func (d *Users) Validate() error {
	if d.TypeUser == "" {
		return errors.New("required type user")
	}

	return nil
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
func (d *TypeProd) Validate() error {
	if d.Name == "" {
		return errors.New("required name")
	}

	return nil
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
func (d *DeliveryOrder) Validate() error {
	if d.CodBill == "" {
		return errors.New("required bill code")
	}
	if d.Address == "" {
		return errors.New("required address")
	}
	if d.Responsable == "" {
		return errors.New("required responsable")
	}

	return nil
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
func (d *Product) Validate() error {
	if d.CodTypeProd == "" {
		return errors.New("required code type product")
	}
	if d.Maker == "" {
		return errors.New("required code maker")
	}
	if d.Model == "" {
		return errors.New("required code model")
	}
	if d.Name == "" {
		return errors.New("required code name")
	}
	if d.Year == "" {
		return errors.New("required code year")
	}
	if d.Price == 0 {
		return errors.New("required code price")
	}

	return nil
}
func (d *Provider) Validate() error {
	if d.CodeTypeProd == "" {
		return errors.New("required code type product")
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
func (d *Stock) Validate() error {
	if d.ProductName == "" {
		return errors.New("required type user")
	}
	if d.TypeProdName == "" {
		return errors.New("required type product name")
	}
	if d.QuantityExist == 0 {
		return errors.New("required quantity exist")
	}
	if d.PurchasePrice == 0 {
		return errors.New("required purchase price")
	}
	if d.SalePrice == 0 {
		return errors.New("required sales price")
	}

	return nil
}

type FindAllProductRequest struct {
}
type FindAllDatosResponse struct {
	TDatos []Product `json:"tdatos"`
	Err    error     `json:"error,omitempty"`
}

func (r FindAllDatosResponse) error() error { return r.Err }

type CreateDatosRequest struct {
	Product Product `json:"product"`
}
type CreateDatosResponse struct {
	Err error `json:"error,omitempty"`
}

func (r CreateDatosResponse) error() error { return r.Err }

type DeleteDatosRequest struct {
	ID string `json:"id"`
}
type DeleteDatosResponse struct {
	Err error `json:"error,omitempty"`
}

func (r DeleteDatosResponse) error() error { return r.Err }
