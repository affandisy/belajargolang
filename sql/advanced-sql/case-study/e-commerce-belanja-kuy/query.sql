CREATE TABLE customers (customer_id SERIAL PRIMARY KEY, email varchar(100) UNIQUE, phone varchar(100));

belanjakuy=# CREATE TABLE products (product_id SERIAL PRIMARY KEY, name varchar(100), price numeric(12, 2) NOT NULL CHECK (price > 0), stock INT NOT NULL CHECK (stock >= 0), created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW());

belanjakuy=# CREATE TABLE product_logs (log_id SERIAL PRIMARY KEY, product_id INT NOT NULL REFERENCES products(product_id) ON DELETE CASCADE, action varchar(50), change_time TIMESTAMP WITH TIME ZONE DEFAULT NOW(), old_price NUMERIC(12, 2), new_price NUMERIC(12,2), old_name varchar(200), new_name varchar(200), old_stock INT, n
ew_stock INT);

belanjakuy=# CREATE FUNCTION log_product_update() RETURNS TRIGGER AS $$
belanjakuy$# BEGIN
belanjakuy$# INSERT INTO product_logs(product_id, action, change_time, old_price, new_price, old_name, new_name, old_stock, new_stock)
belanjakuy$# VALUES
belanjakuy$# (NEW.product_id, 'UPDATED', NOW(), OLD.price, NEW.price, OLD.name, NEW.name, OLD.stock, NEW.stock); RETURN NEW; END; $$ LANGUAGE plpgsql;
CREATE FUNCTION

belanjakuy=# CREATE TRIGGER trg_product_update AFTER UPDATE ON products FOR EACH ROW EXECUTE FUNCTION log_pro
duct_update();

