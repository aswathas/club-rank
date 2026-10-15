# 03 — Decisions

*Grounded in the repo's own commit history (`git log`) and the code as it stands.*

Each decision below is one the repo actually made. Where the plan and the code
disagree, both are recorded, because the disagreement is itself the useful finding.

## The plan was written first, and the code followed loosely

`Instructions.md` (committed 2025-01-12, `6df92b0`, "instruction module set up") and
`Backend Plan (llm).md` (`44263d2`, "boiler plate set up") predate every line of Go.
They are a spec, not a record.

The spec says **Echo**, **Viper**, **JWT**, **Svelte**. The code uses **Gin**, a
hardcoded DSN with no Viper, a **no-op auth middleware**, and **React**. Swaggo was
substituted for hand-written docs — commit `37c2151`, "Swagger Documentation".

**Decision as built:** follow the spec's *shape* (layered Go backend, PostgreSQL,
GORM, migrations folder, PostgreSQL) and pick the most common tool for each slot.
The frontend framework diverged entirely, which is the biggest single departure.

## Layered architecture, kept

`routes.go` wires repository → service → controller explicitly. This is the standard
`Backend Plan (llm).md` layout, and it survived — controllers hold no queries, services
hold no HTTP. Worth preserving: it is why adding a resource means adding three files
rather than editing one.

## Raw SQL migrations, not GORM AutoMigrate

`backend/migrations/001_create_tables.sql` (commit `cb1ab90`) is hand-written SQL with
named constraints and check constraints:

```sql
CONSTRAINT check_reg_no_format CHECK (reg_no ~ '^RA[0-9]{13}$')
```

AutoMigrate cannot express a regex check. So the schema carries real constraints, and
it means `InitDB()` connecting is not enough to run the app — you must apply the file.
`config.go` opens the connection and does nothing else. This is a real setup step that
the README has to state.

## Domain as the scoring unit, not student

The data model puts students inside domains, and `calculateDomainScore` averages over a
domain's students before the criterion weight is applied. The alternative — weighting
every student globally — would let a large domain swamp a small one. Commit `2f06f9b`
("club schema update") and `ab1b576` ("Ranking Domain Intialised") reflect that the
domain layer was deliberately added, not incidental.

## Recency weighting instead of a rolling average

`RECENT_WEIGHT = 0.7`, `HISTORICAL_WEIGHT = 0.3`, `RECENT_MONTHS = 3`. A plain moving
average would let a club coast; a strict recent-only window would discard the record.
The split weights both.

## Inactive students are excluded, not zeroed

`MIN_SCORES = 3`: a student with fewer than 3 scores returns `(0.0, false)` and is
left out of the domain average entirely. Treating them as a zero would let non-participation
drag a club down, which is not what the ranking is measuring — it measures what active
members did.

## Rank number is derived from row order

`GetAllRankings` assigns `Rank: i + 1` while iterating stored rankings, with no `ORDER BY`
guarantee visible in the service. This is the weakest decision in the codebase: the rank
number is not a real ranking, it is a slice index. It needs an explicit sort by score
descending before it means anything.

## Decisions deferred, and what they cost

| Deferred | Evidence | Cost |
| --- | --- | --- |
| Authentication | `middleware.Auth()` is `ctx.Next()` and nothing else | Every `/api` route is open. The role column in `users` means nothing yet. |
| Config from env | DSN is a literal in `config.go` | Cannot deploy anywhere but the original author's LAN (`172.27.132.255`). |
| Web3 rankings | No Solidity or chain client in the tree | The headline goal of `Instructions.md` is unimplemented. |
| Score repo wiring | `NewRankingService` never sets `scoreRepo`; the field stays nil | `calculateStudentScore` returns `false` on every call, so **every domain scores 0**. The headline feature returns zeros. |

That last row is the one to fix first. It is a single field assignment in
`services/ranking_service.go`, and until it is fixed nothing about the ranking output
is meaningful.