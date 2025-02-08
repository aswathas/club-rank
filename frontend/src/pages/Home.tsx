import { Button } from '@/components/ui/button' // shadcn UI component

export default function Home() {
  return (
    <div className="container mx-auto p-4">
      <h1 className="text-3xl font-bold mb-4">Welcome to Club Rank</h1>
      <Button>Get Started</Button>
    </div>
  )
}
