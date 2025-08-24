-- Users table: Stores user account information
CREATE TABLE users (
    id SERIAL PRIMARY KEY,                          -- Unique user ID
    username VARCHAR(50) NOT NULL UNIQUE,           -- Unique username
    email VARCHAR(100) NOT NULL UNIQUE,             -- Unique email address
    password_hash VARCHAR(255) NOT NULL,            -- Hashed password
    role VARCHAR(20) NOT NULL,                      -- User role (e.g., admin, user)
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),    -- Account creation timestamp
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()     -- Last update timestamp
);

-- Transactions table: Records financial transactions between users
CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,                          -- Unique transaction ID
    from_user_id INTEGER REFERENCES users(id),      -- Sender user ID
    to_user_id INTEGER REFERENCES users(id),        -- Receiver user ID
    amount NUMERIC(15,2) NOT NULL,                  -- Transaction amount
    type VARCHAR(20) NOT NULL,                      -- Transaction type (e.g., credit, debit, transfer)
    status VARCHAR(20) NOT NULL,                    -- Transaction status (e.g., pending, completed)
    created_at TIMESTAMP NOT NULL DEFAULT NOW()     -- Transaction creation timestamp
);

-- Balances table: Maintains current balance for each user
CREATE TABLE balances (
    user_id INTEGER PRIMARY KEY REFERENCES users(id), -- User ID (unique, references users)
    amount NUMERIC(15,2) NOT NULL,                    -- Current balance amount
    last_updated_at TIMESTAMP NOT NULL DEFAULT NOW()   -- Last balance update timestamp
);

-- Audit Logs table: Tracks actions performed on entities for auditing
CREATE TABLE audit_logs (
    id SERIAL PRIMARY KEY,                          -- Unique audit log ID
    entity_type VARCHAR(50) NOT NULL,               -- Type of entity (e.g., user, transaction)
    entity_id INTEGER NOT NULL,                     -- Entity ID
    action VARCHAR(50) NOT NULL,                    -- Action performed (e.g., create, update, delete)
    details TEXT,                                   -- Additional details about the action
    created_at TIMESTAMP NOT NULL DEFAULT NOW()     -- Audit log creation timestamp
);

-- Indices for performance: Improve query speed on frequently accessed columns
CREATE INDEX idx_transactions_from_user_id ON transactions(from_user_id);
CREATE INDEX idx_transactions_to_user_id ON transactions(to_user_id);
CREATE INDEX idx_balances_user_id ON balances(user_id);
CREATE INDEX idx_audit_logs_entity_id ON