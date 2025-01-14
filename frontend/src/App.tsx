// src/App.tsx
import React from "react";
import { SignIn, SignedIn, SignedOut, UserButton } from "@clerk/clerk-react";

const App = () => {
  return (
    <div>
      <h1>Welcome to My App</h1>

      {/* Show sign-in form if the user is not signed in */}
      <SignedOut>
        <SignIn />
      </SignedOut>

      {/* Show user button if the user is signed in */}
      <SignedIn>
        <UserButton />
        <p>You are signed in!</p>
      </SignedIn>
    </div>
  );
};

export default App;
