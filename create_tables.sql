-- Create table for Group
CREATE TABLE groups (
    id CHAR(36) PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL
);

-- Create table for Category
CREATE TABLE categories (
    id CHAR(36) PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL,
    group_id CHAR(36) NOT NULL,
    FOREIGN KEY (group_id) REFERENCES groups(id)
);

-- Create table for Spendings
CREATE TABLE spendings (
    id CHAR(36) PRIMARY KEY NOT NULL,
    amount DECIMAL(65, 30) NOT NULL,
    date TIMESTAMP NOT NULL,
    comment TEXT NOT NULL,
    category_id CHAR(36) NOT NULL,
    FOREIGN KEY (category_id) REFERENCES categories(id)
);
