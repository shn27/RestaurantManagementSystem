# RestaurantManagementSystem

A pet project. Basic REST API practise with some CRUD operation.

# How to run
docker compose -f compose.yaml up

# To-DO

1. Use go routine for better perfomance
2. Seed the database first

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
## Database schema
We have modified the dataset if necessary and doable. Based on modified dataset here are the database schema.

![Database Schema](./data/Resturant%20Management%20System.svg)

# Help
https://dev.to/pradumnasaraf/dockerizing-a-golang-api-with-mysql-and-adding-docker-compose-support-9b1
https://tutorialedge.net/golang/go-redis-tutorial/
https://www.freecodecamp.org/news/go-elasticsearch/
