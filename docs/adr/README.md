# Architecture Decision Records

This directory contains Architecture Decision Records (ADRs) for important AdServe technical decisions.

## Why ADRs Exist

The goal is not bureaucratic documentation.

ADRs capture decisions where understanding **why** is important for future engineering work.

Examples:

- choosing PostgreSQL for campaign configuration
- selecting a cache strategy
- deciding Kafka partition keys
- choosing processing semantics
- deciding how budget counters work
- adopting a deployment strategy

Small implementation choices do not need ADRs.

## Naming Convention

Use:

`NNNN-short-title.md`

Examples:

- `0001-use-postgresql-for-campaign-management.md`
- `0002-introduce-redis-serving-cache.md`
- `0003-partition-impression-events-by-campaign.md`

Numbers are chronological and never reused.

## ADR Status

Each ADR should use one of:

- Proposed
- Accepted
- Superseded
- Deprecated
- Rejected

If a later ADR replaces an earlier decision, keep the earlier ADR and mark it Superseded.

## Template

```markdown
# ADR-NNNN: Title

## Status

Proposed | Accepted | Superseded | Deprecated | Rejected

## Date

YYYY-MM-DD

## Context

What problem are we solving?

What constraints, requirements, or observed failures make a decision necessary?

## Decision

What are we choosing?

Be precise enough that an engineer can understand what will actually be implemented.

## Alternatives Considered

### Alternative A

Benefits:
- ...

Costs:
- ...

### Alternative B

Benefits:
- ...

Costs:
- ...

## Consequences

### Positive

- ...

### Negative

- ...

### Risks

- ...

## Operational Considerations

- monitoring
- deployment
- migration
- rollback
- failure behavior

## Follow-Up

What work or validation should happen after this decision?
```

## ADR Review Standard

Before accepting an important ADR, be able to answer:

- What requirement drove this decision?
- What alternatives were seriously considered?
- What new failure modes are introduced?
- What happens at 10x load?
- How would the decision be reversed?
- What evidence would cause us to revisit it?
