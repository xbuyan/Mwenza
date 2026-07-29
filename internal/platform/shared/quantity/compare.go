package quantity

func (q Quantity) Compare(other Quantity) int {
	switch {
	case q.value < other.value:
		return -1
	case q.value > other.value:
		return 1
	default:
		return 0
	}
}

func (q Quantity) Equal(other Quantity) bool {
	return q.Compare(other) == 0
}

func (q Quantity) GreaterThan(other Quantity) bool {
	return q.Compare(other) > 0
}

func (q Quantity) LessThan(other Quantity) bool {
	return q.Compare(other) < 0
}

func (q Quantity) GreaterOrEqual(other Quantity) bool {
	return q.Compare(other) >= 0
}

func (q Quantity) LessOrEqual(other Quantity) bool {
	return q.Compare(other) <= 0
}
