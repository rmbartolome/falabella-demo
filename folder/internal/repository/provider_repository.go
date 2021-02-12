package repository

import (
	"context"
)

type ProviderRepository interface {
	FindAll(ctx context.Context) ([]Provider, error)
	GetByID(ctx context.Context, id string) (Provider, error)
	Create(ctx context.Context, provider Provider) error
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, id string) (Provider, error)
}
type providerRepository struct {
	conn Connection
}

func NewProviderRepository(conn Connection) ProviderRepository {
	return &providerRepository{
		conn: conn,
	}
}
func GetByID(ID string) (found datos.Provider) {
	provider := datos.Provider{}
	db := utils.GetConnection()
	defer db.Close()
	db.Find(&provider, ID)
	return provider
}

func Delete(ID string) (found datos.Provider) {
	provider := datos.Provider{}
	db := utils.GetConnection()
	defer db.Close()
	db.Find(&provider.ID)
	db.Delete(provider)
	return provider
}

func Create(provider) (found datos.Provider) {
	provider := datos.Provider{}
	db := utils.GetConnection()
	defer db.Close()
	err = db.Create(&provider).Error
	return provider
}

func Update(ID string) (found datos.Provider) {
	providerFind := datos.Provider{}
	providerData := datos.Provider{}
	db := utils.GetConnection()
	defer db.Close()
	db.Find(&providerFind, ID)
	db.Model(&providerFind).Updates(providerData)

}

func FindAll( (found datos.Provider) {
	provider := datos.Provider{}
	db := utils.GetConnection()
	defer db.Close()
	db.Find(&provider)
	return provider

}
