import { Link } from "react-router-dom";
import { Button } from "@/components/ui/button";

export default function Navbar() {
  return (
    <nav className="bg-white shadow-sm">
      <div className="container mx-auto px-4">
        <div className="flex h-16 items-center justify-between">
          <div className="flex space-x-8">
            <Link to="/clubs" className="text-gray-700 hover:text-gray-900">
              Clubs
            </Link>
            <Link to="/domains" className="text-gray-700 hover:text-gray-900">
              Domains
            </Link>
            <Link
              to="/tech-heads"
              className="text-gray-700 hover:text-gray-900"
            >
              Tech Heads
            </Link>
            <Link to="/students" className="text-gray-700 hover:text-gray-900">
              Students
            </Link>
          </div>
          <Button
            variant="outline"
            onClick={() => {
              /* handle logout */
            }}
          >
            Logout
          </Button>
        </div>
      </div>
    </nav>
  );
}
