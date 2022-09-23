# gRPC + GraphQl + Go Store

This project seeks to simulate the basic functionalities of an online store, with user creation, authentication and authorization, issuing a purchase order and managing products through the admin user.

The main objective was to explore the go language, applying [gRPC](https://grpc.io/) as a communication tool between services, [GraphQL](https://graphql.org/) to implement a [BFF](https://samnewman.io/patterns/architectural/bff/), serving as a gateway to the application.

## Overview

The diagram below exemplifies how services communicate when the client makes a request through bff. In order to explain, let's imagine issuing a purchase order sent by the client, the following flow will occur:

1 - The GraphQL server will communicate with the order service through the gRPC client.

2 - The authentication token will be validated through the gRPC interceptor

3 - After validating the token, the order service consults the users and products service (gRPC) in order to carry out some validations, for example, check if the quantity of products in the order is available in stock

4 - Validations made, the client receives its response

<p align="center">
<img src="./docs/store.png" rel="Diagram" />
</p>


<sub>Some decisions such as fully decoupling modules, multiple products per order, secret management and some others were de-prioritized to keep the focus on the mentioned technologies.</sub>

## How to play with it?

In your terminal:

```
git clone git@github.com:ruancaetano/grpc-graphql-store.git

docker compose up
```

In your browser, access the GraphQL playground through the url `http://localhost:8004/`

Two users must be available
`ruan@caetano.com/test123` and
`admin@admin.com/teste123`