CREATE TABLE IF NOT EXISTS orders (
    order_id VARCHAR(12)PRIMARY KEY ,
    customer_id VARCHAR(12),
    product_id VARCHAR(10),
    payment_id VARCHAR(12),
    quantity INT NOT NULL,
    unit_price DECIMAL(10,2) NOT NULL,
    total_price DECIMAL(10,2) NOT NULL,
    order_status ENUM('pending', 'shipped', 'delivered', 'canceled') DEFAULT 'pending',
    create_date DATETIME DEFAULT CURRENT_TIMESTAMP,
    update_date DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
