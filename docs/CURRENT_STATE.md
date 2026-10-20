# AdServe Current State

Last updated: 2026-10-20

## Current Phase

**Phase 1 — Go Foundations (In Progress)**

Phase 0 — Project Setup and Baseline is complete.

## Current Objective

Build a solid foundation in Go and backend engineering by implementing and testing a small HTTP service.

The current focus is **Go context, cancellation, timeouts, and graceful shutdown**.

The goal is to understand the fundamentals deeply before introducing PostgreSQL or additional infrastructure.

## Completed

### Project Setup

- Created the AdServe GitHub repository.
- Initialized the Go module: `github.com/pochunhuang1107/adserve`.
- Established canonical project documentation under `docs/`.
- Established the approximately 6-hour-per-week learning budget.
- Adopted a learning approach emphasizing design reasoning, implementation, testing, and code review.

### Go Foundations

Practiced and implemented:

- Variables, types, structs, methods, and pointers.
- Value receivers versus pointer receivers.
- Packages, exported identifiers, and module imports.
- Constructors and domain validation.
- Error handling and sentinel errors.
- Slices and table-driven testing.
- Goroutines and `sync.WaitGroup`.
- `sync.Mutex` and concurrency synchronization.
- Value versus pointer semantics and data ownership.
- Go's standard `net/http` package.
- JSON decoding, URL path parameters, and HTTP status codes.
- HTTP endpoint testing with `httptest`.

These are practiced fundamentals, not yet a declaration of complete proficiency.

### AdServe Implementation

Implemented an in-memory campaign-management prototype with four responsibilities:

**Campaign Domain**
- Campaign construction and validation.
- Name, budget, and active status.
- Business operations including `CanServe()` and `Deactivate()`.

**Campaign Store**
- In-memory storage using `map[CampaignID]Campaign`.
- Sequential campaign ID generation.
- Mutex-protected create, read, and deactivate operations.
- Copy-on-create and copy-on-read behavior to prevent callers from mutating stored state directly.
- Idempotent campaign deactivation.
- `ErrCampaignNotFound` for missing campaigns.

**Campaign Service**
- `CreateCampaign()`
- `DeactivateCampaign()`
- Coordination of application operations and storage.

**HTTP Layer**
- `GET /health` — health check.
- `POST /campaigns` — create a campaign.
- `POST /campaigns/{id}/deactivate` — deactivate a campaign.

The deactivation endpoint uses:
- `204 No Content` for successful or repeated deactivation.
- `404 Not Found` for nonexistent campaigns.
- `400 Bad Request` for malformed campaign IDs.
- `405 Method Not Allowed` for unsupported HTTP methods.

### Testing

Implemented automated tests covering:

- Campaign validation and business behavior.
- Store creation, retrieval, and deactivation.
- Copy isolation between stored and returned campaigns.
- Concurrent campaign creation.
- Service-level operations and error propagation.
- HTTP request handling and response contracts.
- Repeated deactivation and idempotency.

Most recent verification:

```bash
go test -race -count=1 ./...
```

All tests passed on 2026-10-20.

## Current Architecture

AdServe is currently a **single-process Go HTTP application with in-memory storage**, organized into HTTP, Service, Store, and Campaign domain responsibilities.

Conceptual request flow:

```text
HTTP Handler
    ↓
Campaign Service
    ↓
Campaign Store
    ↓
Campaign Domain Operations
```

This is a learning prototype, not a production-ready system.

Current limitations:

- Data is lost when the process restarts.
- No PostgreSQL persistence.
- No distributed coordination across multiple processes.
- No authentication or authorization.
- No production observability or SLOs.
- No request cancellation propagation or graceful shutdown implementation yet.

The HTTP endpoints currently serve as exercises for Go foundations. The formal Campaign Management v1 API and requirements will be designed in a later phase.

## Canonical Documents

- `docs/PROJECT_CHARTER.md`
- `docs/ROADMAP.md`
- `docs/CURRENT_STATE.md`
- `docs/ARCHITECTURE.md`
- `docs/adr/README.md`

Chats contain working discussions. Canonical documents contain the current agreed state.

## Current Technical Focus

**Go context and cancellation**

Upcoming learning topics:

1. `context.Context`, cancellation, and `ctx.Done()`.
2. Context propagation across application layers.
3. Deadlines and timeouts.
4. Cancellation-aware concurrent operations.
5. HTTP server lifecycle and graceful shutdown.
6. Phase 1 integration review and milestone assessment.

Introduce PostgreSQL and additional technologies only when the roadmap and system requirements justify them.

## Immediate Next Actions

1. Commit and push the current working implementation.
2. Update this document.
3. Continue with Go context and cancellation fundamentals.
4. Apply cancellation concepts to AdServe.
5. Implement and test graceful HTTP server shutdown.
6. Review Phase 1 exit criteria before proceeding.

## Known Risks

### Scope Creep

Avoid introducing extra campaign features, infrastructure, or technologies before they serve a clear learning or architectural purpose.

### Tutorial Dependence

Continue requiring independent implementation, reasoning about design choices, and meaningful tests rather than copying complete solutions.

### Incomplete Production Readiness

The current system is intentionally simple. Passing tests and a race-detector run do not establish production reliability or distributed-system correctness.

### Time Sustainability

Maintain approximately six hours per week and prioritize understanding over implementation speed.

### Over-Optimizing for One Company

Continue building transferable backend and distributed-systems engineering skills rather than optimizing for a particular company's current technology stack.

## Open Questions

Defer these decisions to the appropriate phases:

- What should Campaign Management v1 support?
- What are the expected scale and performance requirements?
- What consistency guarantees do campaign updates require?
- When should PostgreSQL replace in-memory storage?
- What will justify Redis or Kafka?
- How should AdServe evolve from its current single-process architecture?
- What observability and reliability requirements should be established?
- Should archiving be reversible, and what campaign lifecycle should be supported?

## Next Review Trigger

Review this document when:

- Phase 1 Go Foundations reaches its milestone review.
- A significant architectural or roadmap decision changes.
- Approximately four weeks have passed since the previous progress review.