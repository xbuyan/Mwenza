#!/usr/bin/env bash
set -e

echo "======================================"
echo "Feature 031 - Sales Aggregate Skeleton"
echo "======================================"

mkdir -p internal/domain/sale

########################################
# sale.go
########################################

cat > internal/domain/sale/sale.go <<'EOGO'
package sale

type Sale struct {
	id string
}
EOGO

########################################
# sale_test.go
########################################

cat > internal/domain/sale/sale_test.go <<'EOGO'
package sale

import "testing"

func TestSaleCanBeCreated(t *testing.T) {
	s := Sale{}

	if s.id != "" {
		t.Fatal("expected empty id")
	}
}
EOGO

gofmt -w internal/domain/sale

go test ./internal/domain/sale

echo
echo "Feature 031 completed successfully."
