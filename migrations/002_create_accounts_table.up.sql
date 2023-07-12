CREATE TABLE IF NOT EXISTS customers (
    id SERIAL PRIMARY KEY,
    customer_id INTEGER REFERENCES customers(id),
    opening_date DATE,
    account_type VARCHAR(10)
);