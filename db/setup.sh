#!/bin/bash
set -euo pipefail

function query {
    psql -d "$POSTGRES_DB" -U "$POSTGRES_USER" -w -c "$1"
}

query 'DROP TABLE IF EXISTS users CASCADE;'
query 'CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    display_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);'

query 'DROP TABLE IF EXISTS todos CASCADE;'
query 'CREATE TABLE IF NOT EXISTS todos ( 
    id SERIAL PRIMARY KEY,
    ref UUID DEFAULT gen_random_uuid(),
    user_id INTEGER NOT NULL,
    title VARCHAR(255) NOT NULL,
    done BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);'
