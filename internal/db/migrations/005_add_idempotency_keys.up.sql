-- SQL migration for adding idempotency keys to transactions table

CREATE TABLE IF NOT EXISTS idempotency_keys (
    id SERIAL PRIMARY KEY,
    key VARCHAR(255) NOT NULL UNIQUE,
    user_id INT NOT NULL,
    transaction_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_user
        FOREIGN KEY(user_id) 
        REFERENCES users(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_transaction
        FOREIGN KEY(transaction_id) 
        REFERENCES transactions(id)
        ON DELETE CASCADE
);

CREATE INDEX idx_idempotency_key ON idempotency_keys(key);