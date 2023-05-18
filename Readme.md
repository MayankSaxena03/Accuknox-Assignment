# Accuknox Assignment README

This assignment involves creating an HTTP API for user authentication and note management. The API provides endpoints for user registration, login, listing user notes, creating new notes, and deleting notes. The API is built using Go (Golang) and runs on port 8080 by default.

## API Endpoints

### Endpoint for creating a new user

- **HTTP Method & URL:** [POST] /signup

- **Request:**
  ```
  {
    "name": <string>,
    "email": <string>,
    "password": <string>
  }
  ```

- **Response:**
  - 200 OK (on success)
  - 400 Bad Request (if the request format is invalid)

### Endpoint for user login

- **HTTP Method & URL:** [POST] /login

- **Request:**
  ```
  {
    "email": <string>,
    "password": <string>
  }
  ```

- **Response:**
  - 200 OK
    ```
    {
      "sid": <string>
    }
    ```
    (`sid` is the session ID, which is unique for each user login)
  - 400 Bad Request (if the request format is invalid)
  - 401 Unauthorized (if the username and password do not match)

### Endpoint for listing all the notes created by a user

- **HTTP Method & URL:** [GET] /notes

- **Request:**
  ```
  {
    "sid": <string>
  }
  ```

- **Response:**
  - 200 OK
    ```
    {
      "notes": [
        {
          "id": <uint32>,
          "note": <string>
        },
        {
          "id": <uint32>,
          "note": <string>
        },
        {
          "id": <uint32>,
          "note": <string>
        }
      ]
    }
    ```
  - 400 Bad Request (if the request format is invalid)
  - 401 Unauthorized (if the `sid` is invalid)

### Endpoint for creating a new note

- **HTTP Method & URL:** [POST] /notes

- **Request:**
  ```
  {
    "sid": <string>,
    "note": <string>
  }
  ```

- **Response:**
  - 200 OK
    ```
    {
      "id": <uint32>
    }
    ```
    (`id` is the ID of the newly created note)
  - 400 Bad Request (if the request format is invalid)
  - 401 Unauthorized (if the `sid` is invalid)

### Endpoint for deleting a note

- **HTTP Method & URL:** [DELETE] /notes

- **Request:**
  ```
  {
    "sid": <string>,
    "id": <uint32>
  }
  ```

- **Response:**
  - 200 OK (on success)
  - 400 Bad Request (if the request format or `id` is invalid)
  - 401 Unauthorized (if the `sid` is invalid)

## Usage

1. Make sure you have Go (Golang) installed on your system.

2. Clone the repository and navigate to the project directory.

3. Run the Go application:
   ```
   go run main.go
   ```

4. The API will be available at `http://localhost:8080`.

5. Use an HTTP client (e.g., cURL, Postman) to send requests to the API endpoints mentioned above. You can also use the Postman collection provided in the `Postman` directory.

Please note that the default port for the Go application is 8080, but you can modify it as needed.