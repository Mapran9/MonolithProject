CREATE TABLE IF NOT EXISTS payments (
    payment_id VARCHAR(12) PRIMARY KEY,
    customer_id VARCHAR(12),
    order_id VARCHAR(12) ,
    payment_method ENUM('credit_card', 'paypal', 'bank_transfer') NOT NULL,
    amount DECIMAL(10,2) NOT NULL,
    status ENUM('pending', 'completed', 'failed') DEFAULT 'pending',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
