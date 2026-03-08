# Courses Management

A web application for instructors to create and manage their courses.

## Tech Stack

- **Backend:** Go, Gin, GORM, JWT, MySQL
- **Frontend:** Vue 3, TypeScript, Vite, Tailwind CSS, TanStack Vue Query

## Project Structure

```
├── backend/
│   ├── cmd/          # Main entry point
│   ├── config/       # Configuration
│   ├── db/           # Database connection
│   ├── dto/          # Data transfer objects
│   ├── handlers/     # Request handlers
│   ├── middleware/   # Custom middleware
│   ├── models/       # Database models
│   ├── routes/       # Route definitions
│   └── utils/        # Utility functions
│
└── frontend/
    └── src/
        ├── components/    # Vue components
        ├── composables/  # Vue composables
        ├── lib/          # API, types, utils
        ├── routes/       # Vue Router config
        └── views/        # Page components
```

## Features

- User registration and login (Instructor)
- JWT-based authentication
- Create, read, update, delete courses
- Dashboard with course overview
- Soft delete for courses

## Getting Started

### Prerequisites

- Go 1.21+
- Node.js 18+
- MySQL database

### Backend Setup

1. Navigate to backend directory
2. Copy `.env.example` to `.env` and configure:
   ```
   DB_HOST=localhost
   DB_PORT=3306
   DB_USER=root
   DB_PASSWORD=
   DB_NAME=courses_db
   JWT_SECRET=your-secret-key
   APP_PORT=8080
   ```
3. Run the server:
   ```bash
   go run cmd/main.go
   ```

### Frontend Setup

1. Navigate to frontend directory
2. Install dependencies:
   ```bash
   npm install
   ```
3. Create `.env` file:
   ```
   VITE_API_URL=http://localhost:8080
   ```
4. Start development server:
   ```bash
   npm run dev
   ```

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | /api/auth/register | Register instructor |
| POST | /api/auth/login | Login |
| GET | /api/courses | Get all courses |
| GET | /api/courses/instructor/:id | Get instructor's courses |
| GET | /api/courses/:id | Get course by ID |
| POST | /api/courses | Create course |
| PUT | /api/courses/:id | Update course |
| DELETE | /api/courses/:id | Delete course |