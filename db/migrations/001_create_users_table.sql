CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name TEXT NOT NULL,
    dob DATE NOT NULL
);

CREATE INDEX idx_users_id ON users(id);