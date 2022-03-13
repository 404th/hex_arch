package api

import (
	"github.com/404th/hex_arch/internal/ports"
)

type Adapter struct {
	arith ports.ArithPorts
	db    ports.DBPort
}

func NewAdapter(db ports.DBPort, arith ports.ArithPorts) *Adapter {
	return &Adapter{arith, db}
}

func (apia *Adapter) GetAddition(a, b int32) (int32, error) {
	result, err := apia.arith.Addition(a, b)
	if err != nil {
		return 0, err
	}

	if err = apia.db.AddToHistory(result, "addition"); err != nil {
		return 0, err
	}

	return result, nil
}

func (apia *Adapter) GetSubtraction(a, b int32) (int32, error) {
	result, err := apia.arith.Subtraction(a, b)
	if err != nil {
		return 0, err
	}

	if err = apia.db.AddToHistory(result, "subtraction"); err != nil {
		return 0, err
	}

	return result, nil
}

func (apia *Adapter) GetMultiplication(a, b int32) (int32, error) {
	result, err := apia.arith.Multiplication(a, b)
	if err != nil {
		return 0, err
	}

	if err = apia.db.AddToHistory(result, "multiplication"); err != nil {
		return 0, err
	}

	return result, nil
}

func (apia *Adapter) GetDivision(a, b int32) (int32, error) {
	result, err := apia.arith.Division(a, b)
	if err != nil {
		return 0, err
	}

	if err = apia.db.AddToHistory(result, "division"); err != nil {
		return 0, err
	}

	return result, nil
}
