#!/usr/bin/env bash
set -euo pipefail

echo "======================================"
echo "Mwenza Sprint 002 - Shared Value Objects"
echo "======================================"

mkdir -p \
internal/platform/shared/currency \
internal/platform/shared/money \
internal/platform/shared/quantity

########################################
# Currency
########################################

cat > internal/platform/shared/currency/currency.go <<'EOGO'
package currency

type Currency string

const (
	KES Currency = "KES"
	USD Currency = "USD"
	EUR Currency = "EUR"
)

func (c Currency) String() string {
	return string(c)
}
EOGO

########################################
# Money
########################################

cat > internal/platform/shared/money/money.go <<'EOGO'
package money

import (
	"errors"

	"github.com/mwenza/mwenza/internal/platform/shared/currency"
)

var ErrCurrencyMismatch = errors.New("currency mismatch")

type Money struct {
	amount   int64
	currency currency.Currency
}

func New(amount int64, curr currency.Currency) Money {
	return Money{
		amount: amount,
		currency: curr,
	}
}

func (m Money) Amount() int64 {
	return m.amount
}

func (m Money) Currency() currency.Currency {
	return m.currency
}

func (m Money) Add(other Money) (Money, error) {
	if m.currency != other.currency {
		return Money{}, ErrCurrencyMismatch
	}

	return Money{
		amount: m.amount + other.amount,
		currency: m.currency,
	}, nil
}
EOGO

########################################
# Quantity
########################################

cat > internal/platform/shared/quantity/quantity.go <<'EOGO'
package quantity

import "errors"

var ErrNegativeQuantity = errors.New("quantity cannot be negative")

type Quantity struct {
	value int64
}

func New(v int64) (Quantity, error) {
	if v < 0 {
		return Quantity{}, ErrNegativeQuantity
	}

	return Quantity{value: v}, nil
}

func (q Quantity) Value() int64 {
	return q.value
}
EOGO

########################################
# Tests
########################################

cat > internal/platform/shared/money/money_test.go <<'EOGO'
package money

import (
	"testing"

	"github.com/mwenza/mwenza/internal/platform/shared/currency"
)

func TestAdd(t *testing.T) {
	a := New(100, currency.KES)
	b := New(200, currency.KES)

	c, err := a.Add(b)
	if err != nil {
		t.Fatal(err)
	}

	if c.Amount() != 300 {
		t.Fatalf("expected 300 got %d", c.Amount())
	}
}
EOGO

cat > internal/platform/shared/quantity/quantity_test.go <<'EOGO'
package quantity

import "testing"

func TestNegativeQuantity(t *testing.T) {
	_, err := New(-1)

	if err == nil {
		t.Fatal("expected error")
	}
}
EOGO

gofmt -w internal

echo
echo "Sprint 002 completed successfully."
