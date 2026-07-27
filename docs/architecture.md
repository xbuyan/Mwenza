# Mwenza Architecture v1.0

Status: Frozen

## Purpose

Mwenza is a Business Operating System for Kenyan SMEs.

This document defines the architectural principles of the project.
These principles are considered stable and should only change when a
business requirement, security requirement, or technical limitation
makes them insufficient.

---

## Principles

### 1. Business First

Business requirements drive the software.
Technology serves the business—not the other way around.

### 2. Security by Design

Security is a foundational concern.
Authentication, authorization, auditing, and data protection are built
into the architecture from the beginning.

### 3. Domain-Driven Design

Business logic lives in the domain.

Infrastructure depends on the domain.

The domain never depends on infrastructure.

### 4. Hexagonal Architecture

External interfaces (HTTP, WhatsApp, CLI, AI, etc.) communicate with
the application through ports and adapters.

Business rules remain independent of delivery mechanisms.

### 5. AI is an Adapter

Artificial Intelligence may recommend actions.

It never bypasses business rules.

All business changes occur through domain commands.

### 6. Every Business Action is Auditable

Every change to business data must produce an audit trail.

Nothing changes silently.

### 7. Cohesion Over Convenience

Packages exist to represent a business capability or a technical
capability.

Avoid generic packages such as:

- helpers
- utils
- misc
- common

### 8. Progressive Development

The system is built one feature at a time.

Each completed feature:

- compiles
- passes tests
- is committed
- is not redesigned without a documented reason

