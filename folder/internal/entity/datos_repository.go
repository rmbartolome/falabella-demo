package entity

import (
	"context"
	"errors"

	"github.com/jinzhu/gorm"
)

type DatosRepository interface {
	FindAll(ctx context.Context) ([]Product, error)
	Create(ctx context.Context, dat Product) error
	Delete(ctx context.Context, id string) error
	//ADD READ AND UPDATE
}

type datosRepository struct {
	conn Connection
}

func NewDatosRepository(conn Connection) DatosRepository {
	return &datosRepository{
		conn: conn,
	}
}

func (r *datosRepository) FindAll(ctx context.Context) ([]Product, error) {
	produ := []Product{}
	err := r.conn.DB.Find(&produ).Error
	if err != nil {
		return []Product{}, err
	}
	return produ, nil
}

func (r *datosRepository) Create(ctx context.Context, produ Product) error {
	err := r.conn.DB.Create(&produ).Error
	if err != nil {
		return err
	}
	r.conn.DB.Save(&produ)
	return nil
}
func (r *datosRepository) Delete(ctx context.Context, id string) error {
	err := r.conn.DB.Delete(&Product{}, id).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("user not found")
		}
		return err
	}
	return nil
}
