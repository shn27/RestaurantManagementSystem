# RestaurantManagementSystem

A pet project. Basic REST API practise with some CRUD operation.

# How to run
docker compose -f compose.yaml up

# Current work progress
1. parse TransactionDate of purchase history (todo)

# To-DO

1. I will use golang for backend.
2. will use cobra
3. will use go-chi
4. MySql for Database
5. K8s for deployment
6. Redis for cache
7. Will try to add Elastic Search


# Tables
Tables in the database:
go-api    |
go-api    | Table: menus
go-api    | Columns:
go-api    |  - id
go-api    |  - restaurant_id
go-api    |  - dish_name
go-api    |  - price
go-api    |  - created_at
go-api    |  - updated_at
go-api    |  - deleted_at
go-api    |
go-api    | Table: opening_hours
go-api    | Columns:
go-api    |  - id
go-api    |  - restaurant_id
go-api    |  - day
go-api    |  - opening_time
go-api    |  - closing_time
go-api    |  - created_at
go-api    |  - updated_at
go-api    |  - deleted_at
go-api    |
go-api    | Table: purchase_histories
go-api    | Columns:
go-api    |  - id
go-api    |  - user_id
go-api    |  - dish_name
go-api    |  - restaurant_name
go-api    |  - transaction_amount
go-api    |  - time
go-api    |  - created_at
go-api    |  - updated_at
go-api    |  - deleted_at
go-api    |
go-api    | Table: restaurants
go-api    | Columns:
go-api    |  - id
go-api    |  - restaurant_name
go-api    |  - cash_balance
go-api    |  - created_at
go-api    |  - updated_at
go-api    |  - deleted_at
go-api    |
go-api    | Table: users
go-api    | Columns:
go-api    |  - id
go-api    |  - user_name
go-api    |  - cash_balance
go-api    |  - created_at
go-api    |  - updated_at
go-api    |  - deleted_at


# Help
https://dev.to/pradumnasaraf/dockerizing-a-golang-api-with-mysql-and-adding-docker-compose-support-9b1

