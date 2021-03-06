package repository

import (
	"context"
)

type ProductRepository interface {
	FindAll(ctx context.Context) ([]Product, error)
	GetByID(ctx context.Context, id string) (Product, error)
	Create(ctx context.Context, product Product) error
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, id string) (Product, error)
}
type productRepository struct {
	conn Connection
}

func NewBProductRepository(conn Connection) ProductRepository {
	return &productRepository{
		conn: conn,
	}
}

func GetByID(ID string) (found datos.Product) {
	product := datos.Product{}
	db := utils.GetConnection()
	defer db.Close()
	db.Find(&product, ID)
	return product
}

func Delete(ID string) (found datos.Product) {
	product := datos.Product{}
	db := utils.GetConnection()
	defer db.Close()
	db.Find(&product.ID)
	db.Delete(product)
	return product
}

func Create(Product) (found datos.Product) {
	product := datos.Product{}
	db := utils.GetConnection()
	defer db.Close()
	err = db.Create(&product).Error
	return product
}

func Update(ID string) (found datos.Product) {
	find := datos.Product{}
	data := datos.Product{}
	db := utils.GetConnection()
	defer db.Close()
	db.Find(&find, ID)
	db.Model(&find).Updates(data)

}

func FindAll() (found datos.Product) {
	product := datos.Product{}
	db := utils.GetConnection()
	defer db.Close()
	db.Find(&product)
	return product

}
