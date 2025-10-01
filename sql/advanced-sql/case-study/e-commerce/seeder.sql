INSERT INTO customers (name, email, city) VALUES ('Andi', 'andi@mail.com', 'Jakarta'),('Huda', 'huda@mail.com', 'Malang'), ('rafly', 'rafly@mail.com', 'Surabaya'), ('Yaya', 'yaya@mail.com', 'Jakarta');

INSERT INTO products (name, category, price) VALUES ('Laptop Lenovo', 'Elektronik', 10000000), ('Iphone 17', 'Elektronik', 21000000),('Headphone Sony', 'Aksesoris', 500000), ('Meja Belajar', 'Furniture', 750000), ('Kursi Gaming', 'Furniture', 1200000);

INSERT INTO orders (customer_id, order_date, total) VALUES (1, '2025-09-01', 750000), (2, '2025-09-02', 1200000), (3, '2025-09-02', 31500000), (4, '2025-09-03', 33400000);

INSERT INTO order_items (order_id, product_id, quantity, subtotal) VALUES (1, 4, 1, 750000), (2, 3, 1, 500000), (2, 4, 1, 750000), (3, 1, 1, 10000000), (3, 2, 1, 21000000), (3, 3, 1, 500000), (4, 1, 1, 10000000), (4, 2, 1, 21000000), (4, 3, 1, 500000), (4, 4, 1, 750000), (4, 5, 1, 1200000);