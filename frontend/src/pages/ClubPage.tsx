import { Button } from "@/components/ui/button";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import {
  Card,
  CardHeader,
  CardTitle,
  CardDescription,
  CardContent,
} from "@/components/ui/card";

export default function ClubPage() {
  return (
    <div className="container mx-auto p-6">
      <Card>
        <CardHeader>
          <CardTitle className="text-3xl font-bold">Clubs</CardTitle>
          <CardDescription>Manage your clubs here</CardDescription>
        </CardHeader>
        <CardContent>
          <div className="flex justify-end mb-4">
            <Button>Add New Club</Button>
          </div>
          <Table>
            <TableHeader>
              <TableRow>
                <TableHead>Name</TableHead>
                <TableHead>Domain</TableHead>
                <TableHead>Members</TableHead>
                <TableHead>Actions</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {/* Example row */}
              <TableRow>
                <TableCell>Tech Club</TableCell>
                <TableCell>Technology</TableCell>
                <TableCell>25</TableCell>
                <TableCell className="space-x-2">
                  <Button variant="outline" size="sm">
                    Edit
                  </Button>
                  <Button variant="destructive" size="sm">
                    Delete
                  </Button>
                </TableCell>
              </TableRow>
            </TableBody>
          </Table>
        </CardContent>
      </Card>
    </div>
  );
}
