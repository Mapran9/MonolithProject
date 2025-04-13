ALTER TABLE `orders`
  ADD CONSTRAINT fk_order_customer FOREIGN KEY (customer_id) REFERENCES customers(customer_id) ON DELETE CASCADE,
  ADD CONSTRAINT fk_order_product FOREIGN KEY (product_id) REFERENCES products(product_id) ON DELETE CASCADE,
  ADD CONSTRAINT fk_order_payment FOREIGN KEY (payment_id) REFERENCES payments(payment_id) ON DELETE CASCADE;