package common

type Trade struct {
	TakerOrdID int32
	MakerOrdID int32
	Price      uint64
	Quantity   uint64
}
