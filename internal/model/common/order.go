package common

type Order struct {
	Name     string
	ReqOrdID int32
	Price    uint64
	Quantity uint64
	Type     int
	Side     int
}
