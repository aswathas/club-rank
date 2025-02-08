import { Card, CardHeader, CardTitle, CardContent } from "@/components/ui/card";

export default function RankingPage() {
  return (
    <div className="container mx-auto p-4">
      <Card>
        <CardHeader>
          <CardTitle>Student Rankings</CardTitle>
        </CardHeader>
        <CardContent>
          <ul className="list-decimal pl-5">
            <li>Alice</li>
            <li>Bob</li>
            <li>Charlie</li>
          </ul>
        </CardContent>
      </Card>
    </div>
  );
}
