Comprehensive Backend Plan for Club Ranking System

Overview

The backend of the club ranking system is designed to handle user management, club data, domain and scoring systems, and leaderboards. It follows a modular and scalable architecture to ensure flexibility and ease of maintenance. This document details the design, structure, and workflows of the backend to help any newcomer quickly understand and contribute.

Project Goals

🎯 Manage multiple clubs with isolated data.

🔑 Provide role-based access (Super Admin, Club Admin, Tech Head).

🏆 Enable scoring and ranking of club members.

🌐 Offer APIs for seamless frontend integration.

Technology Stack

🖥️ Programming Language: Go (Golang)

⚙️ Framework: Echo (for HTTP server and routing)

🗄️ Database: PostgreSQL (Relational database for structured data)

🔒 Authentication: JWT (JSON Web Token) for secure access control

☁️ Hosting: Railway or Heroku (for deployment)

Architecture

Layered Architecture

📥 Controller Layer:

Handles HTTP requests and sends responses.

Interacts with the service layer for business logic.

🧠 Service Layer:

Implements business logic (e.g., validating user input, calculating scores).

Acts as an intermediary between the controller and repository layers.

💾 Repository Layer:

Handles database queries (CRUD operations).

Isolates direct database access from other layers.

🛡️ Middleware:

Ensures security and logging (e.g., authentication, request validation).

Backend Directory Structure

backend/
├── cmd/ # Entry point of the backend application
│ └── main.go # Main server file
├── config/ # Configuration (DB, environment variables)
│ └── config.go
├── controllers/ # Handles HTTP requests
│ ├── user_controller.go
│ ├── club_controller.go
│ └── score_controller.go
├── models/ # Data structures and database models
│ ├── user.go
│ ├── club.go
│ └── score.go
├── repositories/ # Database queries
│ ├── user_repository.go
│ ├── club_repository.go
│ └── score_repository.go
├── routes/ # API route definitions
│ └── routes.go
├── services/ # Business logic
│ ├── user_service.go
│ ├── club_service.go
│ └── score_service.go
├── middleware/ # Authentication and other middleware
│ └── auth.go
└── go.mod # Go module file

Database Design

Tables

Users

Fields:

🆔 id: Unique identifier

👤 name: Name of the user

✉️ email: Email address (unique)

🔑 password: Encrypted password

🏷️ role: Role of the user (super_admin, club_admin, tech_head)

🏛️ club_id: Associated club ID (nullable for Super Admins)

📅 created_at: Timestamp of creation

Clubs

Fields:

🆔 id: Unique identifier

📛 name: Name of the club

🖼️ logo: Club logo URL

🌐 subdomain: Unique subdomain for the club

📅 created_at: Timestamp of creation

Domains

Fields:

🆔 id: Unique identifier

🏷️ name: Domain name (e.g., Coding, Design)

🏛️ club_id: Associated club ID

📅 created_at: Timestamp of creation

Students

Fields:

🆔 id: Unique identifier

👤 name: Name of the student

🔢 roll_number: Unique roll number

🏷️ domain_id: Associated domain ID

📅 created_at: Timestamp of creation

Scores

Fields:

🆔 id: Unique identifier

👥 student_id: Associated student ID

🏷️ domain_id: Associated domain ID

🏆 score: Score given

📅 month: Month of the score

📅 created_at: Timestamp of creation

API Design

Authentication

🔐 POST /login

Authenticate user and return a JWT token.

📝 POST /register

Register a new user (Super Admin or Club Admin).

Club Management

🏛️ POST /clubs

Create a new club.

📋 GET /clubs

List all clubs (Super Admin only).

🔍 GET /clubs/{id}

Get details of a specific club.

Domain Management

➕ POST /domains

Add a new domain to a club.

📋 GET /domains/{club_id}

List all domains in a club.

Student Management

➕ POST /students

Add a student to a domain.

📋 GET /students/{domain_id}

List all students in a domain.

Scoring

📝 POST /scores

Assign a score to a student.

📊 GET /scores/{domain_id}/{month}

Fetch scores for a domain in a specific month.

Leaderboard

🏆 GET /leaderboard/{club_id}

Fetch the overall leaderboard for a club.

Development Workflow

Set Up the Environment:

⚙️ Install Go, PostgreSQL, and necessary libraries.

🛠️ Initialize the Go module with go mod init.

Build Features Incrementally:

🔍 Start with a health check endpoint (GET /health).

🔐 Implement authentication (/login, /register).

🏛️ Add club and domain management APIs.

🏆 Build scoring and leaderboard features.

Test Each Feature:

🧪 Use Postman or curl to verify API functionality.

✅ Write unit tests for critical components.

Deploy and Monitor:

☁️ Deploy the backend to Railway or Heroku.

📊 Monitor logs and performance using platform tools.

Best Practices

Error Handling:

🚨 Always handle errors and return meaningful HTTP status codes.

Security:

🔒 Use JWT for authentication.

🔑 Encrypt passwords using a library like bcrypt.

Scalability:

📈 Use a multi-tenant database design with club_id as a tenant identifier.

Code Quality:

🧹 Follow Go naming conventions.

🧩 Keep functions small and focused.
