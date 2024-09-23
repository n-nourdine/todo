# Todo App

A simple, efficient todo list application built with Go, using the Echo framework for the backend and SQLite3 for data storage.

## Author

Nourdine Nasser
- GitHub: [@n-nourdine](https://github.com/n-nourdine)
- Email: nassernourdine6@gmail.com

## Features

- Create, read, update, and delete todo items
- Mark todos as complete or incomplete
- Filter todos by status
- Persistent storage using SQLite3
- RESTful API design

## Prerequisites

Before you begin, ensure you have the following installed:
- Go (version 1.16 or later)
- SQLite3

## Installation

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/todo-app.git
   cd todo-app
   ```

2. Install dependencies:
   ```
   go mod tidy
   ```

3. Set up the SQLite database:
   ```
   sqlite3 todo.db < schema.sql
   ```

## Configuration

Create a `.env` file in the root directory with the following content:

```
DB_PATH=./todo.db
PORT=8080
```

Adjust these values as needed.

## Running the Application

To start the server, run:

```
go run main.go
```

The server will start on `http://localhost:8080` (or the port you specified in the `.env` file).

## API Endpoints

- `GET /todos`: Retrieve all todos
- `GET /todos/:id`: Retrieve a specific todo
- `POST /todos`: Create a new todo
- `PUT /todos/:id`: Update an existing todo
- `DELETE /todos/:id`: Delete a todo

## Project Structure

```
todo-app
├── main.go
├── handlers
│   └── todo_handlers.go
├── models
│   └── todo.go
├── database
│   └── db.go
├── public
│   ├── 404.html
│   ├── css
│   │   └── style.css
│   └── index.html
├── todo.db
├── .env
├── .gitignore
└── README.md
```

## Testing

Run the tests with:

```
go test ./...
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
