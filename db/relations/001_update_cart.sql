UPDATE carts c
JOIN customers cu ON c.customer_id IS NULL AND cu.customer_id IS NOT NULL
JOIN products p ON c.product_id IS NULL AND p.product_id IS NOT NULL
SET c.customer_id = cu.customer_id, c.product_id = p.product_id
WHERE c.customer_id IS NULL OR c.product_id IS NULL;