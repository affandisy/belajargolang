Table customers {
  id int [pk, increment] // Primary Key
  name varchar
  email varchar [unique]
  city varchar
  preferences jsonb
}

Table products {
  id int [pk, increment]
  name varchar
  category varchar
  price decimal
}

Table orders {
  id int [pk, increment]
  customer_id int [ref: > customers.id] // Many orders per customer
  order_date date
  total decimal
}

Table order_items {
  id int [pk, increment]
  order_id int [ref: > orders.id]    // Many order_items per order
  product_id int [ref: <> products.id] // Many order_items per product
  quantity int
  subtotal decimal
}