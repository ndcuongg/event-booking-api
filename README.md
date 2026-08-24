# Event Booking API

A REST API for managing events that allows users to sign up, log in, create/manage events, and register to attend them. Built with **Go (Gin)**, backed by a **SQLite** database, and authenticated with **JWT**.

## Tech Stack

| Component        | Technology                               |
| ---------------- | ---------------------------------------- |
| Language         | Go 1.25+                                 |
| Web framework    | [Gin](https://github.com/gin-gonic/gin)   |
| Database         | SQLite (`github.com/mattn/go-sqlite3`) |
| Authentication   | JWT (`github.com/golang-jwt/jwt/v5`)   |
| Password hashing | bcrypt (`golang.org/x/crypto`)         |

## Prerequisites

- **Go** >= 1.25 — [download here](https://go.dev/dl/)
- **Git**
- On Linux, install `gcc` (CGO is required to build the `go-sqlite3` driver):
  ```bash
  # Debian/Ubuntu
  sudo apt install build-essential
  ```

## Installation & Getting Started

### 1. Clone the project

```bash
git clone https://github.com/ndcuongg/event-booking-api.git
cd event-booking-api
```

### 2. Install dependencies

```bash
go mod download
# or
go mod tidy
```

### 3. Run the server

```bash
go run main.go
```

The server will start at: **http://localhost:8080**

> The SQLite database (`api.db`) and its tables (`users`, `events`, `registrations`) are **automatically created** on the first server startup (see `db/db.go`).

### Build for production (optional)

```bash
go build -o bin/server main.go
./bin/server
```

## API Endpoints

### Public

| Method | Endpoint        | Description                    |
| ------ | --------------- | ------------------------------ |
| POST   | `/signup`     | Create a new account           |
| POST   | `/login`      | Log in and receive a JWT token |
| GET    | `/events`     | List all events                |
| GET    | `/events/:id` | Get details of a single event  |

### Requires authentication (Header `Authorization: <token>`)

| Method | Endpoint                 | Description               |
| ------ | ------------------------ | ------------------------- |
| POST   | `/events`              | Create a new event        |
| PUT    | `/events/:id`          | Update an event           |
| DELETE | `/events/:id`          | Delete an event           |
| POST   | `/events/:id/register` | Register for an event     |
| DELETE | `/events/:id/cancel`   | Cancel event registration |

### Examples

```bash
# Sign up
curl -X POST http://localhost:8080/signup \
  -H "Content-Type: application/json" \
  -d '{"email": "test@example.com", "password": "123456"}'

# Log in
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"email": "test@example.com", "password": "123456"}'

# Create an event (use the token from the login step)
curl -X POST http://localhost:8080/events \
  -H "Content-Type: application/json" \
  -H "Authorization: <your-token>" \
  -d '{
    "name": "Go Conference",
    "description": "An event about Go",
    "location": "Hanoi",
    "date_time": "2026-09-01T09:00:00Z"
  }'
```

The `api-test/` directory contains ready-made API test files (for the VS Code REST Client).

## Project Structure

```
event-booking-api/
├── main.go           # Entry point: init DB + start server
├── db/
│   └── db.go         # SQLite connection & auto table creation
├── models/           # Data models + SQL queries
│   ├── events.go
│   └── users.go
├── routes/           # Route handlers (controllers)
│   ├── routes.go     # Route registration
│   ├── events.go
│   ├── users.go
│   └── register.go
├── middlewares/
│   └── auth.go       # JWT authentication middleware
├── utils/
│   ├── jwt.go        # Generate/verify JWT tokens
│   └── hash.go       # Password hashing (bcrypt)
└── api.db            # SQLite database file
```
