package product

import (
	"context"
	"fmt"
	"strings"

	"checkout/internal/product"
)

func (r *ProductMySQLRepository) GetBySKUs(skus []string) ([]product.Product, error) {
	if len(skus) == 0 {
		return []product.Product{}, nil
	}

	placeholders := make([]string, len(skus))
	args := make([]interface{}, len(skus))

	for i, sku := range skus {
		placeholders[i] = "?"
		args[i] = sku
	}

	query := fmt.Sprintf(`
		SELECT
			id,
			sku,
			name,
			price,
			stock
		FROM products
		WHERE sku IN (%s)
	`, strings.Join(placeholders, ","))

	rows, err := r.db.QueryContext(context.Background(), query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []product.Product

	for rows.Next() {
		var p product.Product

		err := rows.Scan(
			&p.ID,
			&p.SKU,
			&p.Name,
			&p.Price,
			&p.Stock,
		)
		if err != nil {
			return nil, err
		}

		result = append(result, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
