import { SignIn } from "@clerk/clerk-react";
import {
  Card,
  CardContent,
  CardHeader,
  CardTitle,
  CardDescription,
} from "@/components/ui/card";

export default function LoginPage() {
  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100">
      <Card className="w-[400px]">
        <CardHeader>
          <CardTitle className="text-2xl font-bold">Sign In</CardTitle>
          <CardDescription>
            Enter your credentials to access the app
          </CardDescription>
        </CardHeader>
        <CardContent className="pt-6">
          <SignIn
            appearance={{
              elements: {
                rootBox: "mx-auto",
                card: "shadow-none",
                footer: "hidden",
              },
            }}
          />
        </CardContent>
      </Card>
    </div>
  );
}
