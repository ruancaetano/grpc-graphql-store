CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE users_roles_enum AS ENUM ('admin', 'user');

-- Users
CREATE TABLE users (
    id uuid primary key DEFAULT uuid_generate_v4 () ,
    created_at  TIMESTAMP WITH TIME ZONE default now(),
    updated_at  TIMESTAMP WITH TIME ZONE default now(),
    is_active boolean default TRUE,
    name varchar(255),
    email varchar(255) unique,
    password varchar(255),
    role users_roles_enum default 'user'::users_roles_enum
);

INSERT INTO users (
    id,
    name,
    email,
    password,
    role
) VALUES (
    'eceab8f4-ef4a-4775-a2c0-15d38afa6fd4'::uuid,
    'Ruan Caetano',
    'ruan@caetano.com',
    '$2a$08$EjNMu7WjMQ05ej9mk7PbpublmbAkngADG0tApg9XEnTJggEieE5ju',
    'user'::users_roles_enum
),(
    'cc50b3d7-576a-40c7-9624-a224b38bdc63'::uuid,
    'Ruan Admin',
    'admin@admin.com',
    '$2a$08$EjNMu7WjMQ05ej9mk7PbpublmbAkngADG0tApg9XEnTJggEieE5ju',
    'admin'::users_roles_enum
);


-- Products
CREATE TABLE products (
    id uuid primary key DEFAULT uuid_generate_v4 () ,
    created_at  TIMESTAMP WITH TIME ZONE default now(),
    updated_at  TIMESTAMP WITH TIME ZONE default now(),
    is_active boolean default TRUE,
    title varchar(255) not null,
    description varchar(255),
    thumb varchar(255),
    availables int,
    price float
);

INSERT INTO products (
    id,
    title,
    description,
    thumb,
    availables,
    price
) VALUES (
    '481b107c-e3ac-47cf-b1e2-da6a88c0bc05'::uuid,
    'Camiseta bonita',
    'Camiseta muito bonita',
    'thumb.com/camisetabonita',
    10,
    100
);


-- Orders 
CREATE TABLE orders (
    id uuid primary key DEFAULT uuid_generate_v4 () ,
    created_at  TIMESTAMP WITH TIME ZONE default now(),
    updated_at  TIMESTAMP WITH TIME ZONE default now(),
    is_active boolean default TRUE,
    user_id uuid not null,
    product_id uuid not null,
    quantity int
);

INSERT INTO orders (
    id,
    user_id,
    product_id,
    quantity
) VALUES (
    '58319585-1eb3-49ed-bcce-7b6c79a6ee83'::uuid,
    'eceab8f4-ef4a-4775-a2c0-15d38afa6fd4'::uuid,
    '481b107c-e3ac-47cf-b1e2-da6a88c0bc05'::uuid,
    1
);
