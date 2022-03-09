package api

import (
	"log"

	"github.com/404th/hex_arch/internal/ports"
)

type Adapter struct {
	arith ports.ArithmeticPorts
	db    ports.DBPort
}

func NewAdapter(db ports.DBPort, arith ports.ArithmeticPorts) *Adapter {
	return &Adapter{arith, db}
}

func (apia Adapter) GetAddition(a, b int32) (int32, error) {
	result, err := apia.arith.Addition(a, b)
	if err != nil {
		log.Fatalf("Error while getting addition in api.go : %v", err)
		return 0, err
	}

	if err := apia.db.AddToHistory(result, "addition"); err != nil {
		log.Fatalf("Error while adding to history in ADDITION: %v", err)
	}

	return result, nil
}

func (apia Adapter) GetSubtraction(a, b int32) (int32, error) {
	result, err := apia.arith.Subtraction(a, b)
	if err != nil {
		log.Fatalf("Error while getting subtraction in api.go : %v", err)
		return 0, err
	}

	if err := apia.db.AddToHistory(result, "subtraction"); err != nil {
		log.Fatalf("Error while adding to history in SUBTRACTION: %v", err)
	}

	return result, nil
}

func (apia Adapter) GetMultiplication(a, b int32) (int32, error) {
	result, err := apia.arith.Multiplication(a, b)
	if err != nil {
		log.Fatalf("Error while getting multiplication in api.go : %v", err)
		return 0, err
	}

	if err := apia.db.AddToHistory(result, "multiplication"); err != nil {
		log.Fatalf("Error while adding to history in MULTIPLICATION: %v", err)
	}

	return result, nil
}

func (apia Adapter) GetDivision(a, b int32) (int32, error) {
	result, err := apia.arith.Division(a, b)
	if err != nil {
		log.Fatalf("Error while getting division in api.go : %v", err)
		return 0, err
	}

	if err := apia.db.AddToHistory(result, "division"); err != nil {
		log.Fatalf("Error while adding to history in DIVISION: %v", err)
	}

	return result, nil
}
