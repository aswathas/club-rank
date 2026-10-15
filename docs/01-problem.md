# 01 — The problem

*Grounded in the repo's own commit history (`git log`), January 2025.*

## What was being solved

`Instructions.md` opens with the premise: rank student clubs the way the ICC ranks
cricket players. Not "who did the most projects" but "who did the best work". The
stated intent is to make ranking **centralised and non-negotiable** — only a club head
changes a ranking, and members can look but not touch.

Three problems follow from that:

1. **Quality has no unit.** Activity is easy to count (commits, PRs, events). Quality
   is not. The design answers it with a weighted-criteria scoring system, so a club
   that scores well on the criteria its admin actually cares about can outrank a club
   that is simply busier.
2. **Rankings drift when nobody is looking.** If every member can edit their own score,
   the leaderboard is worthless. Hence the role split in the `users` table —
   `super_admin`, `club_admin`, `tech_head` — and the original JWT plan.
3. **Scores go stale in one direction.** Recent effort and long-term effort are not
   worth the same. The implemented algorithm weights the last 3 months at `0.7` and
   everything older at `0.3`, so a club cannot coast on a good semester forever.

## How the project actually progressed

Six weeks, two contributors (`aswathas`, `DRAGONDINESH`), with a third
(`DINESHLINGAM`) merging branches:

| Date | Milestone |
| --- | --- |
| 2025-01-12 | Instruction module and boilerplate |
| 2025-01-14 | Initial commit; frontend directory and Clerk auth set up |
| 2025-01-25 | API routing; migrations folder; club schema |
| 2025-02-08 | Swagger documentation; frontend basic setup |
| 2025-02-09 | Error debugging; Swagger for users, domains and rankings; ranking module initialised |

Note what the timeline shows: **the ranking logic was the last thing built, a month
after the CRUD scaffolding.** The data model came first — `users`, `clubs`, `domains`,
`students`, `scores` all exist — and the scoring algorithm landed on top of it at the
end.

## What was never reached

The original brief also called for Web3: rankings recorded immutably on-chain, and
tokens minted from end-of-semester average scores. Nothing in this repo implements
that — no Solidity, no chain client, no token code. The brief also planned student and
score endpoints; the tables and models exist, the routes do not.

So the shipped project solves problem 1 (the scoring model) and sets up the shape for
problem 2 (roles exist in the schema) but does not enforce it — `middleware.Auth()`
is currently a no-op. Problem 3 (score recency) is solved in
`services/ranking_service.go`.

See [03-decision.md](03-decision.md) for what was chosen along the way and where the
code diverged from the plan.