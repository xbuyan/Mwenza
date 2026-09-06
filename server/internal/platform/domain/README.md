# Domain Package

This package contains the fundamental building blocks used by every
bounded context in Mwenza.

Rules:

- Business logic lives in the domain.
- Infrastructure depends on the domain.
- The domain never depends on infrastructure.
- Aggregates emit domain events.
- Aggregates enforce invariants.
