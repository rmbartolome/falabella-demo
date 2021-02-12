package repository

import (
	"fmt"
	"log"
)
var logger log.Logger
{
logger = log.NewLogfmtLogger(os.Stderr)
logger = log.With(logger, "ts", log.DefaultTimestampUTC)
logger = log.With(logger, "caller", log.DefaultCaller)
}

// MigrateDB migra la base de datos
func MigrateDB() {
	db := NewConnection(logger)
	defer db.Close()

	fmt.Println("Migrating models....")
	// Automigrate se encarga de migrar la base de datos sí no se ha migrado, y lo hace a partir del modelo
	db.AutoMigrate(
		&Product{},
		&Stock{},
		&Provider{},
		&Bill{},
		&DeliveryOrder{},
		&PurchaseOrder{},
		&TypeProd{},
		&Users{},
		&Client{}
	)
}
