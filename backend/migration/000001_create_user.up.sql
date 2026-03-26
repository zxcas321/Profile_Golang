CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email VARCHAR(150) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    role VARCHAR(20) NOT NULL DEFAULT 'user',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

INSERT INTO users (id, email, password_hash, role, created_at, updated_at)
VALUES (
    uuid_generate_v4(),
    'superuser@admin.com',
    '$2a$10$zX/YywlKjP/I.G2JOt4fbe29ljEK45u/DOsmO/Ui9cvhbDSj0lR8.',
    'superuser',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
);