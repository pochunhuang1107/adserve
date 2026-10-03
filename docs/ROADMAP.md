# AdServe Learning Roadmap

## Guiding Constraint

Target learning time: **~6 hours per week**.

The roadmap is intentionally adaptive. Phases may expand, shrink, overlap, or be reordered based on actual learning progress and relevant professional experience.

Do not advance merely because a date has arrived. Advance when the current phase has produced the intended engineering capability.

---

# Phase 0 — Project Setup and Baseline

## Goals

- establish project conventions
- create the repository
- establish canonical documentation
- define the learning workflow
- establish a baseline of current skills

## Deliverables

- project charter
- roadmap
- current-state document
- initial architecture document
- ADR convention
- `adserve` repository
- first technical chat: `01 — Go Foundations`

## Exit Criteria

- project structure exists
- repository is initialized
- next milestone is clear

---

# Phase 1 — Go Foundations

## Goal

Become productive enough in Go to build and reason about production-style backend services.

## Topics

- packages and modules
- structs and interfaces
- error handling
- pointers and value semantics
- slices and maps
- goroutines
- channels
- synchronization
- `context.Context`
- cancellation and timeouts
- HTTP servers
- structured logging
- testing
- graceful shutdown
- basic profiling

## Project Work

Build small components that will later become part of the Campaign Management Service.

Avoid building unrelated toy projects unless a focused exercise helps isolate a language concept.

## Senior-Engineering Questions

- Where should concurrency be used?
- Where should it not be used?
- How should dependencies be represented?
- What belongs in an interface?
- How should failures propagate through the service?

## Exit Criteria

Able to independently implement and test a small Go HTTP service and explain key concurrency and cancellation behavior.

---

# Phase 2 — Campaign Management Service

## Goal

Design and build the system used by advertisers to define campaigns, ad groups, creatives, budgets, and targeting rules.

## Primary Technologies

- Go
- PostgreSQL
- HTTP/REST initially
- Docker

## Topics

- API design
- validation
- relational data modeling
- indexes
- migrations
- transactions
- connection pools
- pagination
- optimistic/pessimistic concurrency where appropriate
- authentication concepts
- integration testing

## Initial Domain Model

- advertiser
- campaign
- ad group
- creative
- targeting configuration

## Required Design Work

Before major implementation:

- functional requirements
- non-functional requirements
- API contract
- data model
- consistency requirements
- indexing strategy
- failure behavior
- rollout strategy

## Exit Criteria

Campaign management supports a coherent end-to-end workflow with tests, migrations, observability basics, and a documented architecture.

---

# Phase 3 — Online Ad Serving

## Goal

Build a latency-sensitive service that selects an eligible ad for a request.

## Primary Technologies

- Go
- PostgreSQL
- Redis
- load-testing tooling

## Request Path

1. receive ad request
2. identify candidate campaigns
3. apply targeting
4. check eligibility
5. apply frequency/budget constraints
6. rank eligible ads
7. return a winner
8. emit serving events

## Engineering Focus

- latency
- caching
- cache invalidation
- hot keys
- fallback behavior
- concurrency
- graceful degradation
- request timeouts
- performance measurement

## Initial SLO Exercise

Define a realistic learning SLO such as:

- availability target
- p95 latency target
- p99 latency target

The exact numbers must be justified rather than copied from a real company.

## Required Design Questions

- What is allowed on the synchronous request path?
- Which data should be precomputed?
- What happens when Redis is unavailable?
- What consistency guarantees does ad serving actually need?
- What breaks at 10x load?

## Exit Criteria

Service can be load tested, has an explicit latency budget, and has documented cache/failure tradeoffs.

---

# Phase 4 — Kafka and Event Architecture

## Goal

Introduce asynchronous event processing because the online serving path should not synchronously perform every downstream task.

## Event Types

Initially:

- ad request
- impression
- click
- conversion

## Topics

- topics
- partitions
- partition keys
- consumer groups
- ordering
- offset management
- retries
- at-least-once processing
- idempotency
- deduplication
- replay
- poison messages
- dead-letter handling
- schema evolution
- backpressure
- hot partitions

## Failure Exercises

Deliberately test:

- duplicate delivery
- consumer crash
- slow consumer
- malformed events
- reprocessing
- broker interruption
- partition imbalance

## Required ADRs

Likely candidates:

- event partitioning strategy
- processing semantics
- retry/DLQ policy

## Exit Criteria

Can explain the system's delivery semantics and recovery behavior without relying on Kafka vocabulary alone.

---

# Phase 5 — Measurement and Reporting

## Goal

Build near-real-time campaign measurement.

## Metrics

Examples:

- impressions
- clicks
- CTR
- conversions
- CVR
- spend
- CPM
- CPC
- CPA

## Architecture Topics

- event aggregation
- late events
- duplicate events
- windowing concepts
- freshness versus correctness
- OLTP versus OLAP
- materialized/precomputed views
- reporting APIs

## Technology Rule

Start with the smallest datastore architecture that works.

Introduce an analytical datastore or processing framework only after identifying why PostgreSQL or the current design is insufficient.

## Exit Criteria

Measurement produces reproducible campaign metrics with documented correctness/freshness tradeoffs.

---

# Phase 6 — Attribution and Identity

## Goal

Understand one of the hardest cross-cutting Ads domains: connecting impressions/clicks to conversions under explicit attribution rules.

## Topics

- click-through attribution
- view-through attribution
- attribution windows
- late conversion events
- deduplication
- identity linkage
- anonymous versus logged-in identities
- privacy-aware modeling concepts
- rule evolution
- reprocessing historical events

## Identity Examples

- anonymous user ID
- account ID
- device ID
- advertiser-side identifier

## Required Design Questions

- When may identities be linked?
- How are merges represented?
- What happens when identity information arrives late?
- Can attribution be recomputed?
- What data is authoritative?

## Exit Criteria

Attribution behavior is deterministic, testable, and documented for key edge cases.

---

# Phase 7 — Budgeting and Pacing

## Goal

Prevent campaigns from spending too quickly or exceeding intended budget behavior.

## Topics

- distributed counters
- atomic operations
- eventual consistency
- online versus reconciled spend
- pacing algorithms
- overspend bounds
- stale data
- reconciliation

## Initial Exercise

Compare:

- naive daily-budget decrementing
- distributed real-time budget counters
- periodic pacing targets with reconciliation

## Exit Criteria

Can explain correctness versus latency tradeoffs in a distributed budget system.

---

# Phase 8 — Deployment and Kubernetes

## Goal

Deploy AdServe as a set of independently managed services.

## Topics

- Docker images
- configuration
- secrets
- Kubernetes Deployments
- Services
- readiness/liveness probes
- resource requests/limits
- autoscaling
- rolling deployments
- rollback
- service discovery
- graceful termination

## Reliability Exercises

- kill serving pods under load
- simulate unhealthy dependencies
- deploy a broken release
- practice rollback

## Exit Criteria

Can explain how the system is deployed, upgraded, rolled back, and recovered.

---

# Phase 9 — Observability and Reliability

## Goal

Operate AdServe rather than merely run it.

## Service Metrics

Examples:

- QPS
- error rate
- p50/p95/p99 latency
- cache hit ratio
- database latency
- queue/consumer lag

## Business Metrics

Examples:

- impression rate
- click rate
- spend
- no-ad rate
- conversion rate

## Topics

- logs
- metrics
- traces
- dashboards
- alerts
- SLOs
- SLIs
- error budgets
- incident response
- root-cause analysis
- capacity planning

## Required Exercise

Create at least one realistic incident, write an incident timeline, identify root cause, and produce corrective actions.

## Exit Criteria

Can detect and diagnose important failure classes using system telemetry.

---

# Phase 10 — Architecture Evolution / Senior Mode

## Goal

Stop primarily adding features and begin evolving the system deliberately.

## Exercises

Compare architecture generations.

Examples:

### V1
PostgreSQL on the request path.

### V2
Redis cache added to reduce latency.

### V3
Precomputed serving indexes introduced.

For every evolution, document:

- problem being solved
- evidence that the problem exists
- alternatives
- migration plan
- rollback plan
- new failure modes
- operational cost
- complexity added

## Senior-Level Focus

- prioritization
- reducing complexity
- cross-service contracts
- migrations
- backwards compatibility
- operational ownership
- capacity planning
- technical communication
- mentoring-style explanations

## Exit Criteria

Can defend an architecture, identify where it will fail, and propose a safe evolution plan.

---

# Continuous Tracks

These continue throughout the roadmap.

## System Design

Approximately one focused design exercise each week or every other week, preferably derived from AdServe.

Examples:

- ad serving
- impression tracking
- click tracking
- frequency capping
- attribution
- identity graph
- budget pacing
- real-time reporting
- experimentation infrastructure

## Coding Interviews

Maintain without dominating the program.

Focus:

- arrays/hash maps
- trees
- graphs
- heaps
- intervals
- binary search
- BFS/DFS
- basic dynamic programming

Prefer consistency over high volume.

## Low-Level Design

Practice:

- API boundaries
- domain models
- interfaces
- extensibility
- testability
- state machines
- concurrency

## Professional Growth

Look for opportunities in real production work to practice:

- design docs
- ownership
- reliability
- incident response
- migrations
- architectural reasoning
- cross-team communication

Do not share confidential employer information.

---

# Review Cadence

## Weekly

In `00 — Roadmap & Progress`:

1. What was completed?
2. What was learned?
3. What was harder than expected?
4. What is blocked?
5. What are the next ~6 hours?

## Every Four Weeks

Run a Senior Engineer Review covering:

- correctness
- system design
- distributed-systems reasoning
- reliability
- debugging
- ownership
- prioritization
- communication
- simplification

## At Milestone Boundaries

Update:

- `CURRENT_STATE.md`
- `ARCHITECTURE.md` if architecture changed
- relevant ADRs
- roadmap if priorities changed
