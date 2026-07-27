# Mwenza Domain Map

Status: Frozen

## Core Domains

### Inventory
Owns products and stock levels.

Responsibilities:
- Products
- Inventory
- Stock movements
- Reorder levels

---

### Sales
Owns selling to customers.

Responsibilities:
- Quotations
- Orders
- Sales
- Receipts

---

### Procurement
Owns purchasing from suppliers.

Responsibilities:
- Purchase Orders
- Goods Received
- Supplier Deliveries

---

### Finance
Owns money-related business records.

Responsibilities:
- Invoices
- Payments
- Outstanding Balances

---

### Customers
Owns customer information.

Responsibilities:
- Customer Profiles
- Contacts
- Customer History

---

### Suppliers
Owns supplier information.

Responsibilities:
- Supplier Profiles
- Supplier Contacts

---

### Identity
Owns authentication and authorization.

Responsibilities:
- Users
- Roles
- Permissions

---

### Notifications
Owns communication.

Responsibilities:
- WhatsApp
- SMS
- Email
- Reminders

---

### Audit
Owns immutable business history.

Responsibilities:
- Audit Events
- Activity History

---

## Rule

Each business capability owns its own data and business rules.

Domains communicate through well-defined interfaces or domain events.

No domain may modify another domain's internal state directly.

