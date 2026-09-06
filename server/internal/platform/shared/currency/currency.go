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
