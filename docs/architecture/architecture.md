# Mwenza Architecture Blueprint

Version: 1.0
Status: Accepted

---

# Vision

Mwenza is a Business Operating System for African SMEs.

The primary interface is conversational (WhatsApp today, other channels in the future), but the core of the system is a robust business domain—not an AI model.

The AI assists the business.

The AI never owns the business.

---

# Guiding Principles

## 1. Business First

Technology exists to solve business problems.

Every package, service and API must represent a business capability.

---

## 2. Security by Design

Security is a foundation.

It is never added later.

Authentication, authorization, auditing and data protection are architectural concerns.

---

## 3. Domain Driven Design

The domain is the heart of the application.

Infrastructure serves the domain.

The domain never depends on infrastructure.

---

## 4. Hexagonal Architecture

External systems communicate with the application through ports and adapters.

Examples:

- HTTP
- WhatsApp
- CLI
- Scheduled Jobs
- AI Engine

None of these know the business rules.

---

## 5. AI is an Adapter

The AI never modifies business data directly.

Instead:

AI

↓

Command

↓

Domain Validation

↓

Domain Event

↓

Persistence

---

## 6. Every Change is Auditable

Every business action produces an immutable audit trail.

Nothing changes silently.

---

## 7. Explicit Business Language

Every business concept has one meaning.

Examples:

Product

Inventory

Quotation

Invoice

Payment

Customer

Supplier

Stock Movement

No ambiguous names.

---

## 8. No Junk Drawers

The project shall not contain packages such as:

helpers

utils

misc

common

Instead, behavior belongs to the business capability that owns it.

---

# High-Level Layers

Presentation

↓

Application

↓

Domain

↓

Infrastructure

Dependencies point inward.

---

# Core Domains

Inventory

Sales

Procurement

Finance

Customers

Suppliers

Identity

Notifications

Audit

Reporting

AI

---

# Development Rules

Every sprint must:

Compile.

Pass tests.

Be documented.

Be reviewable.

Improve the product.

---

# Long-Term Goal

Build the most trusted Business Operating System for African SMEs.

