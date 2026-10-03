# AdServe Architecture

## Status

**Initial architecture placeholder — Phase 0**

This document describes the current agreed architecture, not every architecture considered historically.

Historical decisions and rejected alternatives belong in ADRs.

## Architectural Principle

Start with the smallest architecture that supports the current learning objective.

Do not begin with a distributed microservice system.

AdServe should earn complexity through requirements.

---

# Initial System Boundary

The first substantive application will be a **Campaign Management Service**.

Its responsibility will be to manage advertiser configuration such as:

- advertisers
- campaigns
- ad groups
- creatives
- budgets
- schedules
- targeting configuration

## Initial Proposed Shape

```text
Client
  |
  v
Campaign Management API
  |
  v
PostgreSQL
```

Technology:

- Go
- HTTP API
- PostgreSQL
- Docker for local packaging

This is intentionally simple.

---

# Expected Evolution

The system may evolve toward the following conceptual architecture.

This is a direction, not a commitment to build every component.

```text
                     Advertiser / Operator
                              |
                              v
                    Campaign Management API
                              |
                         PostgreSQL
                              |
                    campaign/config data
                              |
                              v

Ad Request ---> Ad Serving Service <---- Redis / Serving Data
                     |
             candidate retrieval
                     |
                 targeting
                     |
             budget eligibility
                     |
                 ranking
                     |
                 winner
                     |
          +----------+-----------+
          |                      |
          v                      v
      Impression               Click
          |                      |
          +----------+-----------+
                     |
                    Kafka
                     |
          +----------+-----------+----------------+
          |                      |                |
          v                      v                v
     Measurement            Attribution       Budget/Pacing
          |                      |                |
          v                      v                v
  Reporting/Analytics        identity/state    spend state
```

Later infrastructure may include:

```text
Docker
   |
Kubernetes
   |
deployment / rollout / rollback

Observability:
metrics + logs + traces + dashboards + alerts
```

---

# Architectural Areas

## 1. Campaign Management

Purpose:
- write path for campaign configuration
- advertiser-facing management APIs

Primary concerns:
- correctness
- relational modeling
- transactions
- validation
- maintainability

Initial datastore:
- PostgreSQL

## 2. Ad Serving

Purpose:
- choose an eligible ad under a bounded latency budget

Primary concerns:
- latency
- availability
- caching
- graceful degradation
- consistency tradeoffs

Potential technologies:
- Go
- Redis
- precomputed serving structures

Redis is not yet approved as a project requirement. It should be introduced when the serving design demonstrates a need for it.

## 3. Event Collection

Purpose:
- move non-critical downstream processing off the synchronous request path

Potential event types:
- ad request
- impression
- click
- conversion

Potential technology:
- Kafka

Kafka is not required until asynchronous processing solves a concrete architectural problem.

## 4. Measurement

Purpose:
- calculate campaign metrics and expose reporting data

Primary concerns:
- duplicate events
- late events
- freshness
- correctness
- aggregation strategy
- OLTP versus OLAP

## 5. Attribution and Identity

Purpose:
- associate impressions/clicks with conversions under explicit rules

Primary concerns:
- identity linkage
- late data
- deterministic attribution
- recomputation
- retention/privacy concepts

## 6. Budget and Pacing

Purpose:
- prevent uncontrolled spend and manage delivery over time

Primary concerns:
- distributed counters
- stale state
- reconciliation
- overspend bounds
- consistency versus latency

## 7. Platform and Reliability

Purpose:
- make the system deployable and operable

Future concerns:
- Docker
- Kubernetes
- rollout
- rollback
- observability
- SLOs
- incident response
- capacity planning

---

# Architecture Decision Rules

Before adding a major dependency, answer:

1. What concrete problem exists?
2. Why is the current design insufficient?
3. What alternatives exist?
4. What complexity will the new dependency introduce?
5. What failure modes does it add?
6. How will it be operated?
7. How will it be rolled back or removed?

Important decisions should produce an ADR.

---

# Current Decisions

## Approved

- Go will be the primary language for new backend services in AdServe.
- PostgreSQL will be the first persistent datastore.
- The initial service should remain simple rather than begin as a large microservice architecture.
- Architecture will evolve in response to demonstrated requirements.
- Production characteristics such as reliability and observability are first-class project concerns.

## Not Yet Decided

- exact repository service boundaries
- REST versus gRPC for internal APIs
- Redis topology and cache strategy
- Kafka topic model and partition keys
- analytical datastore
- Kubernetes environment
- observability tooling
- CI/CD tooling
- identity representation
- pacing algorithm
