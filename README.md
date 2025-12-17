# Binary Keeda Api Server

A **production-ready, modular Go backend** using **MongoDB Atlas**, clean architecture, environment-based configuration, and optional Docker support.

This setup is optimized for **local development without Docker** and is **production-aligned** with Atlas.

---

## 🚀 Tech Stack

* Go (net/http + chi)
* MongoDB Atlas
* Official MongoDB Go Driver
* Modular Clean Architecture
* Optional Docker / Docker Compose

---

## 📁 Project Structure

```
go-backend/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── db/
│   │   └── mongo.go
│   ├── handlers/
│   │   ├── health.go
│   │   └── user.go
│   ├── repository/
│   │   └── user_repo.go
│   ├── routes/
│   │   └── routes.go
│   └── models/
│       └── user.go
├── run.sh
├── Dockerfile
├── docker-compose.yml
├── go.mod
└── .env
```

---

## 🧩 Architecture Overview

```
HTTP → Routes → Handlers → Repository → MongoDB Atlas
```

* **Handlers**: HTTP request / response logic
* **Repository**: Database abstraction
* **Models**: Domain entities
* **Config**: Environment configuration
* **DB**: MongoDB connection handling

---

##  MongoDB Atlas Setup

1. Create a cluster in **MongoDB Atlas**
2. Create a **Database User**
3. Whitelist your IP:

   * Atlas → Network Access → Add IP Address
4. Copy the **connection URI**

Example:

```
mongodb+srv://USER:PASSWORD@cluster0.xxxxx.mongodb.net/app_db?retryWrites=true&w=majority
```

---

## ⚙️ Environment Variables

Create a `.env` file in project root:

```env
APP_PORT=8080
MONGO_URI=mongodb+srv://USER:PASSWORD@cluster0.xxxxx.mongodb.net/app_db?retryWrites=true&w=majority
```

⚠️ **Do not commit real credentials**

Recommended:

```
.env        # template
.env.local  # real secrets (gitignored)
```

---

## ▶️ Run Locally (No Docker)

### 1. Install dependencies

```bash
go mod tidy
```

### 2. Make script executable

```bash
chmod +x run.sh
```

### 3. Start server

```bash
./run.sh
```

Expected output:

```
Connected to MongoDB Atlas
Server running on :8080
```

---

## 🧪 API Endpoints

### Health Check

```http
GET /health
```

Response:

```
OK
```

---

### Create User

```http
POST /users
Content-Type: application/json

{
  "email": "test@example.com"
}
```

---

### Get Users

```http
GET /users
```

---

## 🐳 Docker (Optional)

Run API + MongoDB locally using Docker:

```bash
docker-compose up --build
```

---

## 🔒 Production Best Practices

* Use **MongoDB Atlas IP restrictions**
* Store secrets using **env vars / secret manager**
* Add **indexes** for frequently queried fields
* Use **context timeouts** for DB operations
* Enable **structured logging** (zap / zerolog)
* Add **graceful shutdown**

---

## 🔜 Roadmap

* JWT authentication
* Role-based access control (RBAC)
* Request validation
* MongoDB indexes & transactions
* Rate limiting
* Kubernetes deployment

---

## 🧑‍💻 Author Notes

This setup is designed to scale from **local development → production** with minimal changes.

If you need:

* Auth
* Multi-service setup
* Monorepo
* Kubernetes

Feel free to extend this architecture.

---

Happy coding 🚀
