package product

import "database/sql"

type ProductMySQLRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductMySQLRepository {
	return &ProductMySQLRepository{
		db: db,
	}
}
