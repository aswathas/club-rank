# Comprehensive Frontend Plan for Club Ranking System (Svelte)

## **Overview**

The frontend of the club ranking system will be built using **Svelte**, a modern, lightweight framework designed for simplicity and speed. This document outlines the structure, workflows, and design considerations to ensure a user-friendly and maintainable frontend.

---

## **Project Goals**

1. 🖥️ Provide intuitive interfaces for Super Admin, Club Admins, and Tech Heads.
2. 🔗 Seamlessly integrate with backend APIs for authentication, club management, scoring, and leaderboards.
3. 📱 Ensure responsiveness for desktop and mobile users.
4. 🎨 Follow a modular and reusable component architecture.

---

## **Technology Stack**

- **Framework**: Svelte
- **CSS Framework**: Tailwind CSS (for rapid and consistent styling)
- **State Management**: Svelte's built-in reactive stores
- **API Integration**: Fetch API or Axios for HTTP requests
- **Routing**: SvelteKit (for page routing and SSR support)

---

## **Frontend Directory Structure**

```plaintext
frontend/
├── src/
│   ├── components/       # Reusable Svelte components
│   │   ├── Navbar.svelte
│   │   ├── Sidebar.svelte
│   │   ├── Leaderboard.svelte
│   │   └── ScoreForm.svelte
│   ├── pages/            # Pages corresponding to routes
│   │   ├── Login.svelte
│   │   ├── Dashboard.svelte
│   │   ├── ClubManagement.svelte
│   │   └── Leaderboard.svelte
│   ├── stores/           # State management
│   │   └── auth.js
│   ├── services/         # API integration
│   │   └── api.js
│   ├── styles/           # Global styles
│   │   └── global.css
│   └── App.svelte        # Root Svelte component
├── static/               # Static assets (e.g., images, logos)
├── package.json          # Project dependencies
└── svelte.config.js      # Svelte configuration
```

---

## **Component Breakdown**

### **Core Components**

1. **Navbar.svelte**

   - Displays navigation links based on user roles.
   - Example: Home, Dashboard, Leaderboard.

2. **Sidebar.svelte**

   - Sidebar navigation for Admin and Tech Head dashboards.
   - Links to features like Club Management, Scoring, and Domains.

3. **Leaderboard.svelte**

   - Displays club and domain leaderboards.
   - Supports sorting and filtering by domain or date.

4. **ScoreForm.svelte**

   - Form for Tech Heads to submit scores for students.
   - Includes fields for student selection and score input.

5. **AuthModal.svelte** (Optional)
   - Modal for login and registration.

---

## **Page Design**

### **Pages**

1. **Login.svelte**

   - Features:
     - Login form integrated with the `POST /login` API.
     - Error handling for incorrect credentials.
   - State Management:
     - Stores JWT tokens in a secure store.

2. **Dashboard.svelte**

   - Features:
     - Displays key metrics for the logged-in user.
     - Quick links to Club Management, Domains, and Leaderboards.

3. **ClubManagement.svelte**

   - Features:
     - View, add, and update club details.
     - List all domains under a club.

4. **Leaderboard.svelte**

   - Features:
     - Dynamic leaderboard based on club and domain.
     - Sorting and filtering options.

5. **Error.svelte**
   - Displays error messages for invalid routes or server issues.

---

## **API Integration**

### **API Structure**

- **Authentication**:

  - `POST /login`: User login and JWT token retrieval.
  - `POST /register`: Register new users.

- **Club Management**:

  - `GET /clubs`: Fetch all clubs (Super Admin only).
  - `POST /clubs`: Add a new club.

- **Domain Management**:

  - `GET /domains/{club_id}`: Fetch domains for a specific club.
  - `POST /domains`: Add a new domain.

- **Scoring**:

  - `POST /scores`: Submit scores for students.
  - `GET /scores/{domain_id}/{month}`: Fetch scores by domain and month.

- **Leaderboard**:
  - `GET /leaderboard/{club_id}`: Fetch leaderboard for a specific club.

---

## **Development Workflow**

1. **Set Up the Project**:

   - Install SvelteKit:
     ```bash
     npm create svelte@latest frontend
     cd frontend
     npm install
     ```

2. **Create Basic Pages**:

   - Add `Login.svelte`, `Dashboard.svelte`, and other core pages.

3. **Integrate Tailwind CSS**:

   - Install and configure Tailwind:
     ```bash
     npm install -D tailwindcss postcss autoprefixer
     npx tailwindcss init
     ```
   - Add Tailwind directives to `global.css`:
     ```css
     @tailwind base;
     @tailwind components;
     @tailwind utilities;
     ```

4. **Build Components**:

   - Start with reusable components like `Navbar.svelte` and `Sidebar.svelte`.

5. **Connect APIs**:

   - Use `fetch` or Axios to integrate backend APIs.
   - Example:
     ```javascript
     export async function login(credentials) {
       const response = await fetch("/api/login", {
         method: "POST",
         headers: { "Content-Type": "application/json" },
         body: JSON.stringify(credentials),
       });
       return await response.json();
     }
     ```

6. **Test Each Feature**:

   - Use tools like Postman or browser dev tools to test API integration.

7. **Deploy the Frontend**:
   - Deploy to platforms like Vercel or Netlify.

---

## **Best Practices**

1. **State Management**:

   - Use Svelte stores for global state (e.g., user authentication).

2. **Responsiveness**:

   - Test components on various screen sizes.
   - Use Tailwind’s responsive utilities (e.g., `md:flex`, `lg:grid`).

3. **Error Handling**:

   - Show user-friendly error messages for failed API calls.

4. **Code Quality**:
   - Keep components small and focused.
   - Use meaningful names for files and variables.

---

## **Next Steps**

1. Set up the SvelteKit project and create basic pages.
2. Design reusable components and integrate Tailwind CSS.
3. Connect APIs and test the functionality.
4. Deploy and monitor the frontend.

---
