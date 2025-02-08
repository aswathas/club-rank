
const students = [
  { name: 'Alice', rank: 1 },
  { name: 'Bob', rank: 2 },
  { name: 'Charlie', rank: 3 },
]

export default function StudentRanking() {
  return (
    <div className="p-4">
      <h1 className="text-2xl font-bold mb-4">Club Student Rankings</h1>
      <ul>
        {students.map((student) => (
          <li key={student.rank} className="mb-2">
            #{student.rank} - {student.name}
          </li>
        ))}
      </ul>
    </div>
  )
}
