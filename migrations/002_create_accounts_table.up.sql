CREATE TABLE IF NOT EXISTS accounts (
    id SERIAL PRIMARY KEY,
    customer_id INTEGER REFERENCES customers(id),
    opening_date DATE,
    account_type VARCHAR(10)
);