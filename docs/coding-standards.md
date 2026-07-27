# Mwenza Coding Standards

Status: Frozen

## General

- Code must compile.
- All new code must include tests where practical.
- Run `gofmt` on every Go file.
- Keep functions small and focused.
- Prefer readability over cleverness.

---

## Naming

- Package names are lowercase.
- Avoid abbreviations unless they are universally understood (e.g. ID, URL).
- Use meaningful names that reflect the business domain.

Good:

Product
Inventory
Quotation

Avoid:

Data
Manager
Helper
Util

---

## Package Design

A package should have one responsibility.

Do not create packages named:

- helpers
- utils
- misc
- common

---

## Error Handling

Errors are returned.

Errors are never ignored.

Wrap errors with context where appropriate.

---

## Testing

Every new business capability should have unit tests.

Tests should describe business behaviour, not implementation details.

---

## Dependencies

Dependencies point inward.

Infrastructure depends on the domain.

The domain never depends on infrastructure.

---

## Comments

Comment why.

Avoid commenting what the code obviously does.

---

## Commits

Every commit should represent one logical change.

