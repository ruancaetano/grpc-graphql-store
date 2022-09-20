CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id uuid primary key DEFAULT uuid_generate_v4 () ,
    created_at  TIMESTAMP WITH TIME ZONE default now(),
    updated_at  TIMESTAMP WITH TIME ZONE default now(),
    is_active boolean default TRUE,
    name varchar(255),
    email varchar(255) unique,
    password varchar(255)
);

CREATE TABLE products (
    id uuid primary key DEFAULT uuid_generate_v4 () ,
    created_at  TIMESTAMP WITH TIME ZONE default now(),
    updated_at  TIMESTAMP WITH TIME ZONE default now(),
    is_active boolean default TRUE,
    title varchar(255) not null,
    description varchar(255) unique,
    thumb varchar(255),
    availables int
);

CREATE TABLE orders (
    id uuid primary key DEFAULT uuid_generate_v4 () ,
    created_at  TIMESTAMP WITH TIME ZONE default now(),
    updated_at  TIMESTAMP WITH TIME ZONE default now(),
    is_active boolean default TRUE,
    user_id varchar(255) not null,
    product_id varchar(255) unique,
    quantity int
);