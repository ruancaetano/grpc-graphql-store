CREATE DATABASE grpc-graphql-ranking


CREATE TABLE users (
    id serial primary key,
    name varchar(255),
    email varchar(255)
)