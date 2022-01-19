package api

import (
	arith "github.com/daibertdiego/golang_hex_architecture/src/adapters/core/arithmetic"
	"github.com/daibertdiego/golang_hex_architecture/src/ports"
)

type Adapter struct {
	arith ports.ArithmeticPort
	db    ports.DbPort
}

func NewAdapter(arith ports.ArithmeticPort, db ports.DbPort) *Adapter {
	return &Adapter{arith: arith, db: db}
}

func (api *Adapter) GetAddition(a, b int32) (arith.Adapter, error) {
	answer, err := api.arith.Addition(a, b)
	if err != nil {
		return arith.Adapter{}, err
	}

	result, err := api.db.AddToHistory(answer, "addition")
	if err != nil {
		return arith.Adapter{}, err
	}
	return result, nil
}

func (api *Adapter) GetSubtraction(a, b int32) (arith.Adapter, error) {
	answer, err := api.arith.Subtraction(a, b)
	if err != nil {
		return arith.Adapter{}, err
	}

	result, err := api.db.AddToHistory(answer, "subtraction")
	if err != nil {
		return arith.Adapter{}, err
	}
	return result, nil
}

func (api *Adapter) GetMultiplication(a, b int32) (arith.Adapter, error) {
	answer, err := api.arith.Multiplication(a, b)
	if err != nil {
		return arith.Adapter{}, err
	}

	result, err := api.db.AddToHistory(answer, "multiplication")
	if err != nil {
		return arith.Adapter{}, err
	}
	return result, nil
}

func (api *Adapter) GetDivision(a, b int32) (arith.Adapter, error) {
	answer, err := api.arith.Division(a, b)
	if err != nil {
		return arith.Adapter{}, err
	}

	result, err := api.db.AddToHistory(answer, "division")
	if err != nil {
		return arith.Adapter{}, err
	}
	return result, nil
}
