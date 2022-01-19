package ports

import (
	arith "github.com/daibertdiego/golang_hex_architecture/src/adapters/core/arithmetic"
)

type APIPort interface {
	GetAddition(a, b int32) (arith.Adapter, error)
	GetSubtraction(a, b int32) (arith.Adapter, error)
	GetMultiplication(a, b int32) (arith.Adapter, error)
	GetDivision(a, b int32) (arith.Adapter, error)
}
