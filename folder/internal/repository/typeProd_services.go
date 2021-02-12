package repository

import (
	"context"
)

type TypeProdRepository interface {
	FindAll(ctx context.Context) ([]TypeProd, error)
	GetByID(ctx context.Context, id string) (TypeProd, error)
	Create(ctx context.Context, typeProd TypeProd) error
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, id string) (TypeProd, error)
}
type typeProdRepository struct {
	conn Connection
}

func NewTypeProdRepository(conn Connection) TypeProdRepository {
	return &typeProdRepository{
		conn: conn,
	}
}
func GetByID(ID string) (found datos.TypeProd) {
	typeProd := datos.TypeProd{}
	db := utils.GetConnection()
	defer db.Close()
	db.Find(&typeProd, ID)
	return typeProd
}

func Delete(ID string) (found datos.TypeProd) {
	typeProd := datos.TypeProd{}
	db := utils.GetConnection()
	defer db.Close()
	db.Find(&typeProd.ID)
	db.Delete(typeProd)
	return typeProd
}

func Create(TypeProd) (found datos.TypeProd) {
	typeProd := datos.TypeProd{}
	db := utils.GetConnection()
	defer db.Close()
	err = db.Create(&typeProd).Error
	return typeProd
}

func Update(ID string) (found datos.TypeProd) {
	find := datos.TypeProd{}
	data := datos.TypeProd{}
	db := utils.GetConnection()
	defer db.Close()
	db.Find(&find, ID)
	db.Model(&find).Updates(data)

}

func FindAll() (found datos.TypeProd) {
	typeProd := datos.TypeProd{}
	db := utils.GetConnection()
	defer db.Close()
	db.Find(&typeProd)
	return typeProd

}
