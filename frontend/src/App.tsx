// src/App.tsx
import {
  ClerkProvider,
  SignedIn,
  SignedOut,
  SignIn,
  SignUp,
} from "@clerk/clerk-react";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import ClubPage from "./pages/ClubPage";
import Navbar from "./components/Navbar";

const clerkFrontendApi = process.env.REACT_APP_CLERK_FRONTEND_API as string;

const App = () => {
  return (
    <ClerkProvider publishableKey={clerkFrontendApi}>
      <BrowserRouter>
        {/* For signed in users */}
        <SignedIn>
          <Navbar />
          <Routes>
            <Route path="/clubs" element={<ClubPage />} />
            {/* Add other authenticated routes here */}
          </Routes>
        </SignedIn>

        {/* For signed out users */}
        <SignedOut>
          <Routes>
            <Route path="/" element={<SignIn />} />
            <Route path="/signup" element={<SignUp />} />
            {/* You can add a separate login route if you need */}
          </Routes>
        </SignedOut>
      </BrowserRouter>
    </ClerkProvider>
  );
};

export default App;
