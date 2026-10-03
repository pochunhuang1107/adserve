# AdServe Project Charter

## Purpose

AdServe is a long-term learning project for developing from an early-career software engineer into a strong senior backend/distributed-systems engineer.

Reddit Ads engineering is an important target, but the project is optimized for transferable engineering judgment rather than reproducing one company’s current technology stack.

The objective is not to build the largest possible side project.

The objective is to become the engineer capable of designing, building, operating, explaining, and evolving production systems like these.

## Time Budget

Target: approximately **6 hours per week**.

This is a sustainable learning budget, not a quota. Depth, retention, and engineering judgment are more important than schedule speed.

## Flagship System

We are building **AdServe**, a realistic miniature advertising platform.

Over time, AdServe will include:

- campaign management
- ad serving
- targeting and candidate retrieval
- impression/click/conversion events
- measurement and reporting
- attribution and identity
- budget tracking and pacing
- caching
- event-driven architecture
- deployment and orchestration
- observability
- reliability engineering
- production debugging

## Technical Areas

Primary technologies and topics:

- Go
- backend/API design
- PostgreSQL
- Redis
- Kafka
- event-driven architecture
- distributed systems
- Docker
- Kubernetes
- observability
- SLOs and reliability
- system design
- load testing
- incident analysis

Technologies are introduced only when they solve a real system problem.

## Engineering Principles

1. Start simple.
2. Introduce complexity only when the current design is insufficient.
3. Treat AdServe as a production system rather than a tutorial.
4. Make important assumptions explicit.
5. Understand failure modes before scaling the design.
6. Prefer measurable requirements over vague goals.
7. Document important architectural decisions with ADRs.
8. Build observability alongside functionality.
9. Include deployment and rollback thinking in major changes.
10. Refactor existing systems rather than endlessly adding new components.

## Learning Model

The project should progressively change the questions being asked.

### SDE1

> How do I implement this correctly?

Focus:
- correctness
- language fundamentals
- testing
- debugging
- APIs
- data access
- code quality

### SDE2

> How should this service or system be designed?

Focus:
- interfaces
- data modeling
- performance
- consistency
- failure modes
- scalability
- maintainability
- service ownership

### Senior

> What problem should we solve, what tradeoffs should we make, how do multiple systems interact, and how do we deliver and operate this safely?

Focus:
- problem framing
- architecture
- tradeoffs
- cross-system interactions
- operational readiness
- migrations
- rollout/rollback
- ownership
- communication
- simplifying systems
- helping other engineers succeed

## Professional Experience

Relevant production experience from work should be used to reinforce concepts where appropriate.

However:

- AdServe remains a personal learning project.
- Employer code must never be copied into the project.
- Confidential architecture, credentials, internal documents, customer data, private URLs, and proprietary implementation details must not be shared.
- Professional examples should be generalized to the engineering concept being studied.

## Project Organization

### Control-Tower Chat

`00 — Roadmap & Progress`

Use this chat for:

- weekly planning
- roadmap changes
- milestone reviews
- prioritization
- progress assessment
- career-development discussion
- senior-engineer reviews

### Technical Chats

Create a dedicated chat when the primary engineering problem or subsystem changes.

Examples:

- `01 — Go Foundations`
- `02 — Campaign Management Service`
- `03 — Ad Serving System`
- `04 — Kafka & Event Architecture`

Continue using the same technical chat while working on the same subsystem.

### Canonical Documents

Chats contain working discussion.

Canonical documents contain the current agreed state.

The initial canonical documents are:

- `PROJECT_CHARTER.md`
- `ROADMAP.md`
- `CURRENT_STATE.md`
- `ARCHITECTURE.md`
- `adr/README.md`

When a discussion changes architecture, roadmap, status, or project conventions, update the relevant canonical document.

## Review Cadence

### Weekly

Use the control-tower chat to:

- review what was completed
- identify blockers
- choose the next highest-value work
- plan approximately six hours

### Monthly

Run a Senior Engineer Review.

Review:

- technical depth
- architecture judgment
- debugging ability
- ownership
- communication
- production thinking
- ability to identify risk
- ability to simplify

## Success Criteria

Success is not measured primarily by:

- technologies learned
- features shipped
- repository size
- LeetCode count

Success is measured by improved ability to:

- frame ambiguous engineering problems
- design systems with explicit tradeoffs
- diagnose failures
- reason about distributed behavior
- operate production systems
- communicate technical decisions
- safely evolve existing systems
- take ownership of meaningful projects
