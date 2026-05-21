INSERT INTO promotions (
    sku,
    type,
    config
) VALUES

(
    '43N23P',
    'FREE_ITEM',
    JSON_OBJECT(
        'threshold_qty', 1,
        'free_item_sku', '234234',
        'free_item_qty', 1
    )
),

(
    '120P90',
    'BUY_X_FOR_Y',
    JSON_OBJECT(
        'threshold_qty', 3,
        'promotion_qty', 2
    )
),

(
    'A304SD',
    'DISCOUNT',
    JSON_OBJECT(
        'threshold_qty', 3,
        'discount', 0.10
    )
);