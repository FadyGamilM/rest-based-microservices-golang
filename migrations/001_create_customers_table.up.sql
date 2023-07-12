CREATE TABLE IF NOT EXISTS customers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    city VARCHAR(70) NOT NULL,
    zip_code INTEGER CHECK(zip_code <= 99999 AND zip_code >= -99999),
    date_of_birth DATE
);