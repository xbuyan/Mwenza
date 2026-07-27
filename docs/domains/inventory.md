# Inventory Domain Specification

Status: Frozen
Version: 1.0

---

# Purpose

The Inventory domain is responsible for tracking the quantity and availability
of products owned by a business.

It answers questions such as:

- What products do we have?
- How much stock is available?
- What stock is reserved?
- What stock is damaged?
- When should we reorder?

Inventory is the source of truth for stock levels.

---

# Ubiquitous Language

## Product

An item that the business buys, stores, or sells.

Examples:

- Cement
- Hammer
- Paint
- Nails

A Product has an identity but does not store stock quantities.

---

## SKU

Stock Keeping Unit.

A human-readable unique identifier assigned by the business.

Examples:

CEM-001

HAM-025

PAINT-RED-20L

---

## Inventory

Represents the stock state of one Product.

Inventory owns:

- Available Quantity
- Reserved Quantity
- Damaged Quantity
- Reorder Level

---

## Stock Movement

Any action that changes inventory.

Examples:

- Receive Stock
- Sell Stock
- Reserve Stock
- Release Reservation
- Mark Damaged
- Stock Adjustment

Inventory never changes without a Stock Movement.

---

# Business Rules

1. Available quantity must never be negative.

2. Reserved quantity must never exceed available quantity.

3. Damaged quantity cannot be sold.

4. Every stock change must have a reason.

5. Every stock change must be auditable.

6. Every Product has exactly one Inventory record.

---

# Commands

The Inventory domain accepts the following commands:

- ReceiveStock
- ReserveStock
- ReleaseReservation
- SellStock
- AdjustStock
- MarkDamaged
- SetReorderLevel

---

# Domain Events

The Inventory domain publishes the following events:

- StockReceived
- StockReserved
- ReservationReleased
- StockSold
- StockAdjusted
- StockMarkedDamaged
- ReorderLevelChanged

---

# Ownership

Inventory owns inventory rules.

No other domain may change inventory directly.

All changes must go through Inventory.

