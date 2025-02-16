DROP TABLE users;
DROP TABLE wishlists;
DROP TABLE users_wishlists;
DROP TYPE wishlist_role;

-- Enable extensions
CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- Create enums
CREATE TYPE wishlist_role_type AS ENUM ('owner', 'contributor');

-- Create tables
CREATE TABLE users (
    user_id BIGINT UNIQUE NULLS NOT DISTINCT GENERATED ALWAYS AS IDENTITY,
    username TEXT NOT NULL,
    password TEXT NOT NULL,
    email TEXT CONSTRAINT email_unique UNIQUE NOT NULL,
	CONSTRAINT pk_users PRIMARY KEY (user_id)
);

CREATE TABLE wishlists (
    wishlist_id BIGINT UNIQUE NULLS NOT DISTINCT GENERATED ALWAYS AS IDENTITY,
    name TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,
	CONSTRAINT pk_wishlists PRIMARY KEY (wishlist_id)
);

CREATE TABLE users_wishlists (
    id BIGINT UNIQUE NULLS NOT DISTINCT GENERATED ALWAYS AS IDENTITY,
    user_id  BIGINT REFERENCES users,
    wishlist_id BIGINT REFERENCES wishlists,
    user_role wishlist_role NOT NULL,
    CONSTRAINT pk_users_wishlists PRIMARY KEY (user_id, wishlist_id, id)
);

-- Insert dummy data
INSERT INTO users (username, password, email) VALUES ('admin', crypt('root', gen_salt('bf')), 'admin@test.com');
INSERT INTO wishlists (name) VALUES ('My very first wishlist');
INSERT INTO users_wishlists (user_id, wishlist_id, user_role) VALUES (1, 4, 'owner');