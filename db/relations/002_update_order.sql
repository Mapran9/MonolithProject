UPDATE orders o
JOIN customers cu ON o.customer_id IS NULL AND cu.customer_id IS NOT NULL
JOIN products p ON o.product_id IS NULL AND p.product_id IS NOT NULL
JOIN payments py ON o.payment_id IS NULL AND py.payment_id IS NOT NULL
SET o.customer_id = cu.customer_id, o.product_id = p.product_id, o.payment_id = py.payment_id
WHERE o.customer_id IS NULL OR o.product_id IS NULL OR o.payment_id IS NULL;