UPDATE payments py
JOIN customers cu ON py.customer_id IS NULL AND cu.customer_id IS NOT NULL
JOIN orders o ON py.order_id IS NULL AND o.order_id IS NOT NULL
SET py.customer_id = cu.customer_id, py.order_id = o.order_id
WHERE py.customer_id IS NULL OR py.order_id IS NULL;