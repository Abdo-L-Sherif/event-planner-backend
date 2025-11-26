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
    * Create a MySQL database named `eventplanner`.
    * Create a `.env` file in the root directory with your credentials:
        ```env
        PORT=8080
        JWT_SECRET=supersecretkey
        DB_USER=root
        DB_PASSWORD=your_password
        DB_NAME=eventplanner
        DB_HOST=127.0.0.1
        DB_PORT=3306
        ```

3.  **Run the Server:**
    ```bash
    go mod tidy
    go run main.go
    ```
    The server will start on `http://localhost:8080`.

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