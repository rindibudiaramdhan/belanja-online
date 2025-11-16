CREATE TABLE IF NOT EXISTS cart (
    id SERIAL PRIMARY KEY,
    item_id INT NOT NULL REFERENCES items(id),
    amount INT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
