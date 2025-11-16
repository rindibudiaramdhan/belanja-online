CREATE TABLE IF NOT EXISTS items (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    stock INT NOT NULL DEFAULT 0
);

INSERT INTO items (name, stock) VALUES
('Sabun', 20),
('Shampoo', 15),
('Sikat Gigi', 30);
