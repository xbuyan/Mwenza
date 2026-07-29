#!/usr/bin/env bash
set -e

echo "======================================"
echo "Platform 001 - Strengthen IDs"
echo "======================================"

########################################
# id.go
########################################

cat > internal/platform/ids/id.go <<'EOGO'
package ids

import (
	"errors"

	"github.com/google/uuid"
)

var ErrInvalidID = errors.New("invalid id")

type ID string

func New() ID {
	return ID(uuid.NewString())
}

func Parse(value string) (ID, error) {
	if value == "" {
		return "", ErrInvalidID
	}

	if _, err := uuid.Parse(value); err != nil {
		return "", ErrInvalidID
	}

	return ID(value), nil
}

func MustParse(value string) ID {
	id, err := Parse(value)
	if err != nil {
		panic(err)
	}

	return id
}

func (id ID) String() string {
	return string(id)
}

func (id ID) IsZero() bool {
	return id == ""
}
EOGO

########################################
# id_test.go
########################################

cat > internal/platform/ids/id_test.go <<'EOGO'
package ids

import "testing"

func TestNewID(t *testing.T) {
	id := New()

	if id.IsZero() {
		t.Fatal("expected non-zero id")
	}
}

func TestParseValidID(t *testing.T) {
	id := New()

	parsed, err := Parse(id.String())
	if err != nil {
		t.Fatal(err)
	}

	if parsed != id {
		t.Fatal("parsed id mismatch")
	}
}

func TestParseInvalidID(t *testing.T) {
	_, err := Parse("abc")

	if err != ErrInvalidID {
		t.Fatalf("expected %v got %v", ErrInvalidID, err)
	}
}

func TestMustParse(t *testing.T) {
	id := New()

	if MustParse(id.String()) != id {
		t.Fatal("unexpected id")
	}
}

func TestIsZero(t *testing.T) {
	var id ID

	if !id.IsZero() {
		t.Fatal("expected zero id")
	}
}
EOGO

gofmt -w internal/platform/ids

go test ./internal/platform/ids

echo
echo "Platform 001 completed successfully."
