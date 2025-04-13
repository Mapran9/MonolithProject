CREATE TABLE IF NOT EXISTS carts (
    cart_id VARCHAR(10) PRIMARY KEY,
    customer_id VARCHAR(12),
    product_id VARCHAR(10),
    quantity INT NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    status ENUM('pending', 'ordered', 'canceled') DEFAULT 'pending',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
