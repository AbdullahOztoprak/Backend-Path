#!/bin/bash

# Seed the database with initial data

# Load environment variables from .env file
if [ -f .env ]; then
    export $(cat .env | xargs)
fi

# Database connection details
DB_HOST=${DB_HOST:-localhost}
DB_PORT=${DB_PORT:-5432}
DB_USER=${DB_USER:-postgres}
DB_PASSWORD=${DB_PASSWORD:-your_secure_password}
DB_NAME=${DB_NAME:-go_backend_db}

# Function to execute SQL commands
execute_sql() {
    PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -U $DB_USER -d $DB_NAME -c "$1"
}

# Seed users
execute_sql "INSERT INTO users (username, email, password_hash, role) VALUES 
('admin', 'admin@example.com', '$(echo -n 'admin_password' | openssl passwd -stdin -1)', 'admin'),
('user1', 'user1@example.com', '$(echo -n 'user1_password' | openssl passwd -stdin -1)', 'user'),
('user2', 'user2@example.com', '$(echo -n 'user2_password' | openssl passwd -stdin -1)', 'user');"

# Seed roles
execute_sql "INSERT INTO roles (name) VALUES 
('admin'),
('user');"

# Seed transactions
execute_sql "INSERT INTO transactions (from_user_id, to_user_id, amount, description) VALUES 
(1, 2, 100.00, 'Initial transfer from admin to user1'),
(2, 3, 50.00, 'Initial transfer from user1 to user2');"

echo "Database seeded successfully."