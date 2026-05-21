package promotion

import (
	"database/sql"
)

type PromotionMySQLRepository struct {
	db *sql.DB
}

func NewPromotionRepository(db *sql.DB) *PromotionMySQLRepository {
	return &PromotionMySQLRepository{
		db: db,
	}
}
