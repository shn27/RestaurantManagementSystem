# RestaurantManagementSystem

A pet project. Basic REST API practise with some CRUD operation.

# How to run
docker compose -f compose.yaml up

# Current work progress
1. parse TransactionDate of purchase history (todo)
2. write README

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

  Table: menus
  Columns:
   - id
   - restaurant_id
   - dish_name
   - price
   - created_at
   - updated_at
   - deleted_at
 
  Table: opening_hours
  Columns:
   - id
   - restaurant_id
   - day
   - opening_time
   - closing_time
   - created_at
   - updated_at
   - deleted_at
 
  Table: purchase_histories
  Columns:
   - id
   - user_id
   - dish_name
   - restaurant_name
   - transaction_amount
   - time
   - created_at
   - updated_at
   - deleted_at
 
  Table: restaurants
  Columns:
   - id
   - restaurant_name
   - cash_balance
   - created_at
   - updated_at
   - deleted_at
 
  Table: users
  Columns:
   - id
   - user_name
   - cash_balance
   - created_at
   - updated_at
   - deleted_at


# Help
https://dev.to/pradumnasaraf/dockerizing-a-golang-api-with-mysql-and-adding-docker-compose-support-9b1
https://tutorialedge.net/golang/go-redis-tutorial/
https://www.freecodecamp.org/news/go-elasticsearch/
