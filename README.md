# 📸 GoShare — A Simple Social Media Platform Built with Go

GoShare is a lightweight social media web application built with **Golang**, **Chi router**, and **MySQL**.  
It provides a simple API for users to share posts, follow others, and explore shared content — designed for learning backend development and API design in Go.

---

## 🚀 Features

- 👤 User registration and authentication (JWT-based)
- 📝 Create, read, update, and delete (CRUD) posts
- ❤️ Like and follow functionality
- 🔒 Secure routes using middleware
- 🗄️ PostgreSQL as the main data store
- ⚡ Built using [Chi](https://github.com/go-chi/chi) — a lightweight and idiomatic Go router

---

## 🧱 Tech Stack

| Layer          | Technology |
| -------------- | ---------- |
| Language       | Go 1.22+   |
| Web Framework  | Chi Router |
| Database       | PostgreSQL |
| Drive          | pq         |
| Authentication | JWT        |
| Environment    | Docker     |

---


## 🧰 Development Setup

### ✅ 1. Environment Variables with direnv

This project uses **direnv** to automatically load environment variables.

1. Install direnv: https://direnv.net/
2. Create a `.envrc` file at the project root:
   ```sh
   export DB_USER=postgres
   export DB_PASSWORD=yourpassword
   export DB_NAME=social
   export DB_PORT=5432
   export DB_HOST=localhost
   ```
3. Allow direnv:
   ```sh
   direnv allow .
   ```

direnv will now auto-load these variables whenever you enter the project directory.

---

### ✅ 2. Run Using Docker Compose

Make sure Docker is installed.  
To start all services:

```sh
docker compose up --build -d
```

To stop containers:

```sh
docker compose down
```

To view running containers:

```sh
docker ps
```

---

