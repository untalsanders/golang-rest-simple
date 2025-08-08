# Golang REST Simple

A simple REST API built with Go and `gorilla/mux` to manage a list of tasks.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

You need to have Go installed on your machine. You can download it from [here](https://golang.org/dl/).

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/untalsanders/golang-rest-simple.git
   ```
2. Go to the project directory
    ```sh
    cd golang-rest-simple
    ```
3. Download dependencies
    ```sh
    go mod download
    ```

### Running the application

To run the application, execute the following command:

```sh
go run main.go
```

The server will start on port `3500`.

## Usage

You can use a tool like `curl` or Postman to interact with the API. The `requests.http` file is also available with sample requests.

### API Endpoints

*   **`GET /`**: Welcome message.
*   **`GET /tasks`**: Get all tasks.
*   **`POST /tasks`**: Create a new task.
    *   **Body (raw JSON):**
        ```json
        {
            "Name": "New Task",
            "Content": "Some new content"
        }
        ```
*   **`GET /tasks/{id}`**: Get a single task by its ID.
*   **`DELETE /tasks/{id}`**: Delete a task by its ID.
*   **`PUT /tasks/{id}`**: Update a task by its ID.
    *   **Body (raw JSON):**
        ```json
        {
            "Name": "Updated Task Name",
            "Content": "Updated content"
        }
        ```

## Dependencies

This project uses the following external libraries:

*   [github.com/gorilla/mux](https://github.com/gorilla/mux): A powerful URL router and dispatcher for Go.
