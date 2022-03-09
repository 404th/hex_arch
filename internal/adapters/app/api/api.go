package api

import (
	"log"

	"github.com/404th/hex_arch/internal/ports"
)

type Adapter struct {
	arith ports.ArithmeticPorts
}

func (apia Adapter) GetAddition(a, b int32) (int32, error) {
	result, err := apia.arith.Addition(a, b)
	if err != nil {
		log.Fatalf("Error while getting addition in api.go : %v", err)
		return 0, err
	}

	return result, nil
}

func (apia Adapter) GetSubtraction(a, b int32) (int32, error) {
	result, err := apia.arith.Subtraction(a, b)
	if err != nil {
		log.Fatalf("Error while getting subtraction in api.go : %v", err)
		return 0, err
	}

	return result, nil
}

func (apia Adapter) GetMultiplication(a, b int32) (int32, error) {
	result, err := apia.arith.Multiplication(a, b)
	if err != nil {
		log.Fatalf("Error while getting multiplication in api.go : %v", err)
		return 0, err
	}

	return result, nil
}

func (apia Adapter) GetDivision(a, b int32) (int32, error) {
	result, err := apia.arith.Division(a, b)
	if err != nil {
		log.Fatalf("Error while getting division in api.go : %v", err)
		return 0, err
	}

	return result, nil
}
