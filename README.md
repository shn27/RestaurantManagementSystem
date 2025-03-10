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


--------
### Host & Base Path
- **Host**: `<yourhost>:8080`
- **Base Path**: `/`

### Endpoints
### 1. **Process a Purchase**
**POST** `/purchase`

This endpoint allows a user to purchase a dish from a restaurant. It handles the potiential race condition and transaction atomically by:
- Deducting the user's balance.
- Updating the restaurant's balance.
- Recording the purchase history.

#### Request
- **Consumes**: `application/json`
- **Request Body**:
  ```json
  {
    "dish_id": 1,
    "user_id": 1
  }
  ```

#### Responses
- **200**: Purchase successful.
- **400**: Invalid request.
- **500**: Internal server error.

---

### 2. **Get Open Restaurants**
**GET** `/restaurants/open`

Retrieve a list of restaurants that are open at a specific date and time.

#### Request
- **Consumes**: `application/json`
- **Query Parameters**:
  - `datetime` (string, required): Datetime in the format `YYYY-MM-DD HH:MM:SS`.

#### Responses
- **200**: A list of restaurants open at the specified datetime.
  ```json
  [
    {
      "restaurant_name": "Restaurant A"
    }
  ]
  ```
- **400**: Missing or invalid `datetime` parameter.
- **500**: Internal server error.

---

### 3. **Get Top Restaurants**
**GET** `/restaurants/top`

Retrieve the top restaurants based on the number of dishes within a specified price range. Users can specify:
- Minimum and maximum price of dishes.
- Whether to get restaurants with more or fewer dishes than a certain count.
- The maximum number of restaurants to return.

#### Request
- **Consumes**: `application/json`
- **Query Parameters**:
  - `minPrice` (number, required): Minimum price of dishes.
  - `maxPrice` (number, required): Maximum price of dishes.
  - `minDishes` (integer, required): Minimum or maximum number of dishes, depending on `moreOrLess`.
  - `limit` (integer, required): Maximum number of restaurants to return.
  - `moreOrLess` (string, required): Condition for dish count (`"more"` or `"less"`).

#### Responses
- **200**: A list of top restaurants based on the number of dishes.
  ```json
  [
    {
      "id": 1,
      "name": "Restaurant A",
      "dish_count": 10
    }
  ]
  ```
- **400**: Invalid or missing query parameters.
- **500**: Internal server error.

---

### 4. **Search for Restaurants and Dishes**
**GET** `/search`

Search for restaurants and dishes by name. The search term is matched against both restaurant and dish names.

#### Request
- **Consumes**: `application/json`
- **Query Parameters**:
  - `search` (string, required): Search term.

#### Responses
- **200**: A list of matching restaurants and dishes.
  ```json
  [
    {
      "name": "Pizza",
      "type": "dish"
    },
    {
      "name": "Italian Restaurant",
      "type": "restaurant"
    }
  ]
  ```
- **400**: Missing or invalid `search` parameter.
- **500**: Internal server error.


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

