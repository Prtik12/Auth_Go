# 🔐 Auth_Golang

Authentication system built with **Golang** (backend) and **React + TypeScript + Vite** (frontend). Features secure user registration, login, JWT-based authentication, and clean UI components.

---

## 🧩 Tech Stack

**Frontend:**
- React + TypeScript
- Vite
- CSS Modules

**Backend:**
- Go (Golang)
- JWT for authentication
- bcrypt for password hashing
- Prisma ORM
- PostgreSQL

---

## ✨ Features

- 🔐 User authentication (Sign Up / Sign In / Sign Out)
- 🔑 Secure JWT token generation and validation
- 🔒 Password hashing with bcrypt
- 🗃️ Prisma-based database access
- 🧼 Clean and modular file structure (both frontend & backend)
- 🌐 CORS enabled for client-server communication

---

## 👾 Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/Prtik12/Auth_Golang.git
cd Auth_Golang
```

---

## 📦 Backend Setup (`/server`)

### 1. Install Go dependencies

```bash
cd server
go mod tidy
```

### 2. Set up `.env` file

Create a `.env` file in the `server/` directory:

```
DATABASE_URL=your_postgres_db_url
```

> ✅ Make sure `.env` is listed in `.gitignore`

### 3. Run database migrations

```bash
npx prisma migrate dev --name init
```

### 4. Start the server

```bash
go run main.go
```

---

## 💻 Frontend Setup (`/client`)

### 1. Install dependencies

```bash
cd client
npm install
```

### 2. Start the frontend dev server

```bash
npm run dev
```

The app will run at `http://localhost:5173` (or whichever port Vite chooses).

---

## 🗂️ Project Structure

```
prtik12-auth_golang/
├── client/
│   ├── src/
│   │   ├── components/
│   │   │   ├── AuthPage.tsx
│   │   │   ├── SignIn.tsx
│   │   │   ├── SignOut.tsx
│   │   │   └── SignUp.tsx
│   │   └── ...
│   └── ...
└── server/
    ├── controllers/
    │   └── auth.go
    ├── models/
    │   └── user.go
    ├── utils/
    │   ├── bycrpt.go
    │   └── jwt.go
    ├── prisma/
    │   ├── schema.prisma
    │   └── migrations/
    └── main.go
```

---

## 🙌 Acknowledgements

Built as a learning project to explore full-stack authentication using Go.

