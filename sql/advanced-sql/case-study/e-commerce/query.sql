-- CTE (Common Table Expression) 
WITH TotalBelanja AS (SELECT c.city, SUM(o.total) AS total_per_kota FROM customers c JOIN orders o ON c.id = o.customer_id GROUP BY c.city) 
SELECT city, total_per_kota, RANK() OVER (ORDER BY total_per_kota DESC) AS rangking FROM TotalBelanja;

-- Window Function
SELECT c.name, c.city, SUM(o.total) AS total_belanja, RANK() OVER (PARTITION BY c.city ORDER BY SUM(o.total) DESC) FROM customers c JOIN orders o ON c.id = o.customer_id GROUP BY c.id, c.name, c.city;

-- Subquery
SELECT c.name, (SELECT SUM(total) FROM orders WHERE customer_id = c.id) AS total_belanja FROM customers c WHERE (SELECT SUM(total) FROM orders WHERE customer_id = c.id) > (SELECT AVG(total) FROM orders);

-- JSON
SELECT * FROM customers WHERE preferences->>notif = 'sms';

-- ROLLUP
SELECT p.category, SUM(oi.subtotal) AS total_sales FROM products p JOIN order_items oi ON p.id = oi.product_id GROUP BY ROLLUP(p.category);

-- RANK
SELECT p.name, SUM(oi.quantity) AS total_terjual, RANK() OVER(ORDER BY SUM(oi.quantity) DESC) AS rank_product FROM products p JOIN order_items oi ON p.id = oi.product_id GROUP BY p.id, p.name;