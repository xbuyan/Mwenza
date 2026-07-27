#!/usr/bin/env bash
set -e

echo "======================================"
echo "Feature 032 - Sale Constructor"
echo "======================================"

########################################
# status.go
########################################

cat > internal/domain/sale/status.go <<'EOGO'
package sale

type Status string

const (
	StatusDraft     Status = "draft"
	StatusConfirmed Status = "confirmed"
	StatusCancelled Status = "cancelled"
	StatusCompleted Status = "completed"
)
EOGO

########################################
# sale.go
########################################

cat > internal/domain/sale/sale.go <<'EOGO'
package sale

import "errors"

var ErrEmptySaleID = errors.New("sale id cannot be empty")

type Sale struct {
	id     string
	status Status
}

func New(id string) (*Sale, error) {
	if id == "" {
		return nil, ErrEmptySaleID
	}

	return &Sale{
		id:     id,
		status: StatusDraft,
	}, nil
}
EOGO

########################################
# sale_test.go
########################################

cat > internal/domain/sale/sale_test.go <<'EOGO'
package sale

import "testing"

func TestCreateSale(t *testing.T) {
	s, err := New("sale-001")
	if err != nil {
		t.Fatal(err)
	}

	if s.id != "sale-001" {
		t.Fatalf("unexpected sale id")
	}

	if s.status != StatusDraft {
		t.Fatalf("expected draft status")
	}
}

func TestCreateSaleWithoutID(t *testing.T) {
	_, err := New("")

	if err != ErrEmptySaleID {
		t.Fatalf("expected %v got %v", ErrEmptySaleID, err)
	}
}
EOGO

gofmt -w internal/domain/sale

go test ./internal/domain/sale

echo
echo "Feature 032 completed successfully."
