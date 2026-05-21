package promotion

import (
	"context"
	"fmt"
	"strings"

	"checkout/internal/promotion"
)

func (r *PromotionMySQLRepository) GetBySKUs(skus []string) ([]promotion.Promotion, error) {
	if len(skus) == 0 {
		return []promotion.Promotion{}, nil
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
			type,
			config
		FROM promotions
		WHERE sku IN (%s)
	`, strings.Join(placeholders, ","))

	rows, err := r.db.QueryContext(context.Background(), query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []promotion.Promotion

	for rows.Next() {
		var p promotion.Promotion

		err := rows.Scan(
			&p.ID,
			&p.SKU,
			&p.Type,
			&p.Config,
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
