CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create table for Group
CREATE TABLE IF NOT EXISTS groups (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL
);

-- Create table for Category
CREATE TABLE IF NOT EXISTS categories (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    group_id UUID NOT NULL,
    FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE
);

-- Create table for Spendings
CREATE TABLE IF NOT EXISTS spendings (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    amount DECIMAL(65, 2) NOT NULL,
    date TIMESTAMP NOT NULL,
    comment TEXT NOT NULL,
    category_id UUID NOT NULL,
    FOREIGN KEY (category_id) REFERENCES categories(id)
);
