# Product Domain Specification

Status: Frozen
Version: 1.0

---

# Purpose

The Product domain defines what the business buys and sells.

A Product describes an item.

It does not track stock levels.

Stock is owned by the Inventory domain.

---

# Ubiquitous Language

## Product

A uniquely identifiable item offered by the business.

Examples:

- Bamburi Cement 50kg
- Roofing Nails 4 Inch
- Red Oxide Paint 20L

---

## SKU

The business identifier for a product.

Requirements:

- Unique
- Human-readable
- Never reused

Example:

CEM-001

---

## Unit of Measure

Defines how a product is counted.

Examples:

- Piece
- Bag
- Box
- Kilogram
- Litre
- Metre

---

## Product Status

A product may be:

- Active
- Inactive
- Discontinued

Inactive products cannot be sold.

Discontinued products remain in history but cannot be created again.

---

# Product Attributes

Every Product has:

- Product ID
- SKU
- Name
- Description (optional)
- Unit of Measure
- Status

---

# Business Rules

1. SKU must be unique.

2. Product name cannot be empty.

3. Unit of Measure is mandatory.

4. Deleted products are never physically removed.

They become inactive to preserve business history.

---

# Commands

- CreateProduct
- UpdateProduct
- ActivateProduct
- DeactivateProduct
- DiscontinueProduct

---

# Domain Events

- ProductCreated
- ProductUpdated
- ProductActivated
- ProductDeactivated
- ProductDiscontinued

---

# Ownership

The Product domain owns product information.

It never owns inventory quantities.

