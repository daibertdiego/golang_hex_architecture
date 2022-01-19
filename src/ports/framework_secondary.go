package ports

import arith "github.com/daibertdiego/golang_hex_architecture/src/adapters/core/arithmetic"

type DbPort interface {
	CloseDbConnection()
	AddToHistory(answer int32, operation string) (arith.Adapter, error)
}
