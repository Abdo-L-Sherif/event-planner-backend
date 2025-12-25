# EventPlanner Backend (Go)

This is the REST API for the EventPlanner application, built with Go (Golang), Gin, and GORM.

## 🛠️ Tech Stack
* **Language:** Go (Golang)
* **Framework:** Gin Gonic
* **Database:** MySQL
* **ORM:** GORM

## 🚀 Getting Started

### Prerequisites
* Go (v1.20+)
* MySQL Server

### Installation

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/Abdo-L-Sherif/event-planner-backend/
    cd event-planner-backend
    ```

2.  **Setup the Database:**
<<<<<<< HEAD
    * For **development**: The app defaults to SQLite with a local `users.db` file.
    * For **production/OpenShift**: Use MySQL with environment variables.
    * Create a `.env` file in the root directory with your credentials:
        ```env
        # Server Configuration
        PORT=8080

        # Security
        JWT_SECRET=your_super_secure_jwt_secret_here
        CORS_ORIGINS=http://localhost:3000,http://localhost:8080

        # Database Configuration
        DB_TYPE=mysql  # or 'sqlite' for development
=======
    * Create a MySQL database named `eventplanner`.
    * Create a `.env` file in the root directory with your credentials:
        ```env
        PORT=8080
        JWT_SECRET=supersecretkey
>>>>>>> 15fa3d9ca933de6a2f9567693bff307b392c5d8c
        DB_USER=root
        DB_PASSWORD=your_password
        DB_NAME=eventplanner
        DB_HOST=127.0.0.1
        DB_PORT=3306
<<<<<<< HEAD
        DB_PATH=users.db  # Only used for SQLite
=======
>>>>>>> 15fa3d9ca933de6a2f9567693bff307b392c5d8c
        ```

3.  **Run the Server:**
    ```bash
    go mod tidy
    go run main.go
    ```
<<<<<<< HEAD
    The server will start on the configured port (default: `http://localhost:8080`).

## 🚀 OpenShift Deployment

This application is configured for deployment on Red Hat OpenShift:

### Environment Variables for OpenShift

Set these environment variables in your OpenShift deployment:

```bash
# Required
JWT_SECRET=your_super_secure_jwt_secret_here
DB_TYPE=mysql
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=your_db_name
DB_HOST=your_mysql_service_host
DB_PORT=3306

# Optional
PORT=8080  # OpenShift will assign a random port if not set
CORS_ORIGINS=https://your-frontend-domain.com
```

### Health Checks

The application provides health check endpoints for OpenShift probes:
- `GET /health` - Basic health check
- `GET /ready` - Readiness check (includes database connectivity)

### Container Features

- Graceful shutdown handling for proper container lifecycle
- Non-root user execution
- Multi-stage Docker build for optimized image size
- Database connection retry logic
=======
    The server will start on `http://localhost:8080`.
>>>>>>> 15fa3d9ca933de6a2f9567693bff307b392c5d8c

## 📡 API Endpoints

### Auth
* `POST /register` - Create account
* `POST /login` - Login & get Token

### Events
* `POST /events/` - Create Event
* `GET /events/organized` - View my events
* `GET /events/invited` - View invitations
* `GET /events/search?q=keyword&date=yyyy-mm-dd&role=organizer` - Search
* `GET /events/:id` - Event Details
* `POST /events/:id/invite` - Invite User (Body: `{ "invitee_id": 1 }`)
* `POST /events/:id/rsvp` - RSVP (Body: `{ "status": "Going" }`)
* `DELETE /events/:id` - Delete Event