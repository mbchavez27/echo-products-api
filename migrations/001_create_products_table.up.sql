CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    price DECIMAL NOT NULL,
    ratings FLOAT NOT NULL,
    description TEXT NOT NULL,
    category VARCHAR(200) NOT NULL
);
