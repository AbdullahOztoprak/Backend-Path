package postgres

import (
    "database/sql"
    "log"
    "os"

    _ "github.com/lib/pq"
)

type Migration struct {
    ID   string
    Up   string
    Down string
}

var migrations = []Migration{
    {
        ID:   "001_create_users",
        Up:   "CREATE TABLE users (id SERIAL PRIMARY KEY, username VARCHAR(50) UNIQUE NOT NULL, password_hash VARCHAR(255) NOT NULL, email VARCHAR(100) UNIQUE NOT NULL, created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP);",
        Down: "DROP TABLE users;",
    },
    {
        ID:   "002_create_transactions",
        Up:   "CREATE TABLE transactions (id SERIAL PRIMARY KEY, from_user_id INT NOT NULL, to_user_id INT NOT NULL, amount DECIMAL(10, 2) NOT NULL, description TEXT, created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, FOREIGN KEY (from_user_id) REFERENCES users(id), FOREIGN KEY (to_user_id) REFERENCES users(id));",
        Down: "DROP TABLE transactions;",
    },
    {
        ID:   "003_create_balances",
        Up:   "CREATE TABLE balances (id SERIAL PRIMARY KEY, user_id INT NOT NULL, balance DECIMAL(10, 2) NOT NULL, created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, FOREIGN KEY (user_id) REFERENCES users(id));",
        Down: "DROP TABLE balances;",
    },
    {
        ID:   "004_add_roles_and_permissions",
        Up:   "CREATE TABLE roles (id SERIAL PRIMARY KEY, name VARCHAR(50) UNIQUE NOT NULL); CREATE TABLE user_roles (user_id INT NOT NULL, role_id INT NOT NULL, FOREIGN KEY (user_id) REFERENCES users(id), FOREIGN KEY (role_id) REFERENCES roles(id));",
        Down: "DROP TABLE user_roles; DROP TABLE roles;",
    },
    {
        ID:   "005_add_idempotency_keys",
        Up:   "ALTER TABLE transactions ADD COLUMN idempotency_key VARCHAR(255) UNIQUE;",
        Down: "ALTER TABLE transactions DROP COLUMN idempotency_key;",
    },
}

func RunMigrations(db *sql.DB) {
    for _, migration := range migrations {
        if err := executeMigration(db, migration); err != nil {
            log.Fatalf("Failed to execute migration %s: %v", migration.ID, err)
        }
    }
}

func executeMigration(db *sql.DB, migration Migration) error {
    _, err := db.Exec(migration.Up)
    if err != nil {
        return err
    }
    log.Printf("Migration %s executed successfully", migration.ID)
    return nil
}

func RollbackMigration(db *sql.DB, migrationID string) error {
    for _, migration := range migrations {
        if migration.ID == migrationID {
            _, err := db.Exec(migration.Down)
            if err != nil {
                return err
            }
            log.Printf("Migration %s rolled back successfully", migration.ID)
            return nil
        }
    }
    return os.ErrNotExist
}