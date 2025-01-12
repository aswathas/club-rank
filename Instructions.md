# Club Ranking System: Instructions for AI Agents

## Overview

This document provides instructions for AI agents working on the **Web3-based Club Ranking System**, modeled after ICC player rankings. The system is designed to prioritize the quality of projects over quantity. Only the club head can modify rankings, ensuring a centralized ranking authority.

---

## Goals

- Build a Web3-enabled ranking system.
- Evaluate club members based on performance, focusing on project quality.
- Ensure rankings are immutable by regular users.

---

## Key Functionalities

1. **Authentication**:

   - JWT-based user authentication.
   - Roles: Club Head (admin privileges), Club Members (view-only access).

2. **Club Management**:

   - Register clubs.
   - Manage domains within clubs.

3. **Student Management**:

   - Add and update student profiles.
   - Assign students to specific domains.

4. **Ranking System**:

   - Club head assigns scores to members.
   - Scores are converted into rankings.

5. **Web3 Integration**:

   - Use blockchain for recording rankings.
   - Immutable ranking data stored on-chain.
   - At the end of the semester, average scores are calculated and used to generate tokens.

---

## Project Structure

### File System

```
club-ranking-system/
├── cmd/               # Entry points
├── internal/          # Core modules
│   ├── auth/          # Authentication logic
│   ├── clubs/         # Club-related features
│   ├── students/      # Student management
│   ├── ranking/       # Ranking logic
│   └── utils/         # Utility functions
├── migrations/        # Database migrations
├── config/            # Configuration files
├── docs/              # Documentation
└── go.mod             # Dependency file
```

---

## APIs

### Authentication

- `POST /login`: User login (JWT token).
- `POST /register`: Register club heads.

### Club Management

- `POST /clubs`: Create a new club.
- `GET /clubs`: List all clubs.

### Student Management

- `POST /students`: Add students.
- `GET /students`: List students.

### Ranking

- `POST /rankings`: Add or update rankings.
- `GET /rankings`: Fetch rankings.

---

## Web3 Integration

1. **Blockchain Platform**: Use Ethereum or Polygon.
2. **Smart Contract**:
   - Store rankings and metadata.
   - Ensure write access for club heads only.
   - Generate tokens based on averaged scores at the semester’s end.

---

## Frontend and Backend

### Frontend
- **Framework**: Svelte for a lightweight and responsive UI.
- **Features**:
  - Authentication pages for login and signup.
  - Club management dashboards.
  - Real-time rankings display.

### Backend
- **Language**: Go
- **Framework**: Gin or Echo
- **Database**: PostgreSQL
- **Functionality**:
  - APIs to manage clubs, students, rankings, and authentication.
  - Secure token generation at semester-end.

---

## Development Tools

- **Language**: Go
- **Framework**: Gin or Echo
- **Database**: PostgreSQL
- **Blockchain**: Solidity for smart contracts
- **Others**:
  - GORM for ORM
  - Viper for configuration

---

## Next Steps

1. Set up the Go project and initialize dependencies.
2. Build basic APIs for club and student management.
3. Design and deploy the smart contract.
4. Integrate Web3 for immutable ranking storage.
5. Generate tokens after averaging scores at semester’s end.

6. **Test and Deploy (Beginner-Friendly)**:

   - **Testing**:
     - Use Postman or curl to test each API endpoint (e.g., login, student management).
     - Validate responses and error handling.
     - Test the smart contract using tools like Remix or Hardhat.
     - Simulate token generation by creating sample scores and calculating averages.
   
   - **Deployment**:
     - Deploy the backend on a free or low-cost platform like Heroku or Fly.io.
     - Deploy the frontend using services like Vercel or Netlify.
     - For the blockchain part, deploy the smart contract to a testnet (e.g., Polygon Mumbai).

   - **Teaching Points**:
     - Explain how to use tools like Postman to beginners.
     - Show step-by-step deployment with screenshots or commands.
     - Clarify how to interact with the blockchain via a Web3 wallet like MetaMask.

7. Document the entire process for others to follow easily.

# Creating Backend Directory Structure

## Step 1: Create Backend Directory
```bash
mkdir -p backend/cmd
mkdir -p backend/internal/auth
mkdir -p backend/internal/clubs
mkdir -p backend/internal/students
mkdir -p backend/internal/ranking
mkdir -p backend/internal/utils
mkdir -p backend/migrations
mkdir -p backend/config
mkdir -p backend/docs
touch backend/go.mod
```
