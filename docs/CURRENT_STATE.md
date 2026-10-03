# AdServe Current State

Last updated: 2026-10-03

## Current Phase

**Phase 0 — Project Setup and Baseline**

## Current Objective

Finish project initialization and begin **Phase 1 — Go Foundations**.

## Completed

- Created ChatGPT Project: `AdTech Backend → Senior SWE`
- Set project memory to Project-only
- Renamed control-tower chat to `00 — Roadmap & Progress`
- Moved the roadmap conversation into the project
- Defined approximately 6 hours/week as the sustainable learning budget
- Chosen **AdServe** as the flagship project
- Established the goal of developing senior backend/distributed-systems skills
- Established Reddit Ads engineering as an important target, but not the sole optimization target
- Established a canonical-document model:
  - chats contain working discussion
  - canonical documents contain current agreed state

## Canonical Documents

- `PROJECT_CHARTER.md`
- `ROADMAP.md`
- `CURRENT_STATE.md`
- `ARCHITECTURE.md`
- `adr/README.md`

## Current Architecture

No production architecture exists yet.

The intended first real service is the **Campaign Management Service**.

Initial technology sequence:

1. Go
2. PostgreSQL
3. HTTP API
4. Docker

Redis, Kafka, Kubernetes, and additional infrastructure will be introduced later when justified by system requirements.

See `ARCHITECTURE.md` for the evolving architecture.

## Current Technical Focus

**Go foundations**

The next technical chat should be:

`01 — Go Foundations`

## Immediate Next Actions

1. Create the `adserve` GitHub repository.
2. Add these canonical documents under `docs/`.
3. Create `docs/adr/`.
4. Add a top-level `README.md`.
5. Create the technical chat `01 — Go Foundations`.
6. Establish a baseline Go environment and begin Phase 1.
7. Before building the Campaign Management Service, define its v1 requirements and non-functional requirements.

## Known Risks

### Scope Creep

AdTech provides many opportunities to add technologies and features. The project must avoid becoming a technology-collection exercise.

Mitigation:
- require each major technology to solve an identified problem
- review roadmap priorities regularly

### Tutorial Dependence

There is a risk of following implementation instructions without developing design judgment.

Mitigation:
- reason about requirements and tradeoffs before major design answers
- require explanations for important decisions
- use design reviews

### Over-Optimizing for Reddit's Current Stack

The target company’s tools may change.

Mitigation:
- prioritize backend/distributed-systems fundamentals
- treat company-specific technologies as secondary unless they solve a project need

### Time Sustainability

Six hours/week is limited.

Mitigation:
- prefer depth over breadth
- keep milestones small
- avoid unnecessary infrastructure
- allow roadmap timing to adapt

## Open Questions

These should be answered during upcoming phases rather than decided prematurely.

- What exact API style should Campaign Management v1 use?
- What initial scale assumptions should AdServe use?
- What initial latency and availability goals should the online serving path have?
- At what point does Redis become justified?
- At what point does Kafka become justified?
- Should the repository begin as a modular monolith or separate service directories?
- What observability stack should be introduced when operational work begins?

## Next Review Trigger

Review this document after:

- the Go foundations milestone is complete, or
- a major architecture/roadmap decision changes, or
- one month of work has elapsed.
